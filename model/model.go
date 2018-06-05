package model

import (
	"math"
	"math/rand"

	"gonum.org/v1/gonum/mat"

	"github.com/ddkdl/svm/kernel"
)

// Model is a good model
type Model struct {
	w             mat.Vector
	b             float64
	alpha         []float64
	kernel        kernel.Kernel
	tolerance     float64
	X             *mat.Dense
	Y             mat.Vector
	C             float64
	documentCount int
	errors        []float64
}

// NewModel is a good contructor
func NewModel(kernel kernel.Kernel, C, tolerance float64) *Model {
	model := new(Model)

	model.w = nil
	model.b = 0
	model.alpha = nil

	model.kernel = kernel
	model.tolerance = tolerance

	model.X = mat.NewDense(0, 0, nil)
	model.Y = nil
	model.C = C
	model.documentCount = 0
	model.errors = nil

	return model
}

func (model *Model) loadTrainingSet(documentTermMatrix *mat.Dense, trainingLabels mat.Vector) {
	model.X = documentTermMatrix
	model.documentCount, _ = model.X.Dims()
	model.Y = trainingLabels
	model.alpha = make([]float64, model.documentCount)
	model.errors = make([]float64, model.documentCount)
}

// Train goes chuu chuu
func (model *Model) Train(maxPasses int) {
	passes := 0

	for passes < maxPasses {
		var numAlphasChanged = 0

		for i := 0; i < model.documentCount; i++ {
			model.errorFor(i)

			if model.checkBoundariesFor(i) {
				j := rand.Intn(model.documentCount)

				for i == j {
					j = rand.Intn(model.documentCount)
				}

				model.errorFor(j)

				oldAlphaI := model.alpha[i]
				oldAlphaJ := model.alpha[j]

				L, H := model.computeBounds(i, j)

				if L == H {
					continue
				}

				nParam := model.computeNParam(i, j)

				if nParam >= 0 {
					continue
				}

				model.alpha[j] = model.clipAlpha(i, j, L, H, nParam)

				// if the alpha difference is less than 1e-5
				if math.Abs(model.alpha[j]-oldAlphaJ) < 1e-5 {
					continue
				}

				// compute new alpha i
				model.alpha[i] = model.updateAlpha(i, j, oldAlphaJ)

				// compute b1
				b1 := model.computeB1(i, j, oldAlphaI, oldAlphaJ)

				// compute b2
				b2 := model.computeB2(i, j, oldAlphaI, oldAlphaJ)

				// compute b
				model.b = model.computeB(i, j, b1, b2)

				// increment number of alphas changed
				numAlphasChanged++
			}
		}

		if numAlphasChanged == 0 {
			passes++
		} else {
			passes = 0
		}
	}

}

// This is for checking KKT conditions
// TODO: Rename this functions and whatnot
func (model *Model) checkBoundariesFor(i int) bool {
	conditionOne := ((model.Y.At(i, 0) * model.errors[i]) < -(model.tolerance)) && (model.alpha[i] < model.C)
	conditionTwo := ((model.Y.At(i, 0) * model.errors[i]) > (model.tolerance)) && (model.alpha[i] > 0)

	return conditionOne || conditionTwo
}

// Classify does great things
func (model *Model) Classify(document mat.Vector) float64 {
	sum := 0.0

	for j := 0; j < model.documentCount; j++ {
		sum += model.alpha[j] * model.Y.At(j, 0) * model.kernel.Evaluate(model.X.RowView(j), document)
	}

	sum += model.b

	if sum >= 0 {
		return 1
	}

	return -1

}

func (model *Model) errorAt(i int) float64 {
	return model.errors[i]
}

func (model *Model) errorFor(i int) {
	model.errors[i] = model.Classify(model.X.RowView(i)) - model.Y.At(i, 0)
}

func (model *Model) computeBounds(i, j int) (float64, float64) {
	var L float64
	var H float64

	if model.Y.At(i, 0) != model.Y.At(j, 0) {
		diff := model.alpha[j] - model.alpha[i]
		L = math.Max(0, diff)
		H = math.Min(model.C, model.C+diff)
	} else {
		sum := model.alpha[j] + model.alpha[i]
		L = math.Max(0, sum-model.C)
		H = math.Min(model.C, sum)
	}

	return L, H
}

func (model *Model) computeNParam(i, j int) float64 {
	kij := model.kernel.Evaluate(model.X.RowView(i), model.X.RowView(j))
	kii := model.kernel.Evaluate(model.X.RowView(i), model.X.RowView(i))
	kjj := model.kernel.Evaluate(model.X.RowView(j), model.X.RowView(j))

	return 2*kij - kii - kjj
}

func (model *Model) clipAlpha(i, j int, L, H, nParam float64) float64 {
	step := model.Y.At(j, 0) * (model.errors[i] - model.errors[j]) / nParam

	newAlpha := model.alpha[j] - step

	if newAlpha > H {
		return H
	} else if newAlpha < L {
		return L
	} else {
		return newAlpha
	}
}

func (model *Model) updateAlpha(i, j int, oldAlphaJ float64) float64 {
	return model.alpha[i] + model.Y.At(i, 0)*model.Y.At(j, 0)*(oldAlphaJ-model.alpha[j])
}

func (model *Model) computeB1(i, j int, oldAlphaI, oldAlphaJ float64) float64 {
	partOne := model.Y.At(i, 0) * (model.alpha[i] - oldAlphaI) * model.kernel.Evaluate(model.X.RowView(i), model.X.RowView(i))
	partTwo := model.Y.At(j, 0) * (model.alpha[j] - oldAlphaJ) * model.kernel.Evaluate(model.X.RowView(i), model.X.RowView(j))

	return model.b - model.errors[i] - partOne - partTwo
}

func (model *Model) computeB2(i, j int, oldAlphaI, oldAlphaJ float64) float64 {
	partOne := model.Y.At(i, 0) * (model.alpha[i] - oldAlphaI) * model.kernel.Evaluate(model.X.RowView(i), model.X.RowView(j))
	partTwo := model.Y.At(j, 0) * (model.alpha[j] - oldAlphaJ) * model.kernel.Evaluate(model.X.RowView(j), model.X.RowView(j))

	return model.b - model.errors[j] - partOne - partTwo
}

func (model *Model) computeB(i, j int, b1, b2 float64) float64 {
	if model.alpha[i] > 0 && model.alpha[i] < model.C {
		return b1
	} else if model.alpha[j] > 0 && model.alpha[j] < model.C {
		return b2
	} else {
		return (b1 + b2) / 2
	}
}
