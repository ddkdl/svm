package model

import (
	"math"
	"math/rand"

	"github.com/ddkdl/svm/kernel"
)

type Trainer struct {
	kernel            kernel.Kernel
	C                 float64
	w                 []float64
	b                 float64
	numberOfDocuments int
	alpha             []float64
	tolerance         float64
	X                 []float64
	y                 []float64
}

func (trn Trainer) train(maxPasses int) {
	// initialize alphas and b with 0
	for i := 0; i < trn.numberOfDocuments; i++ {
		trn.alpha[i] = 0
	}
	// initialize passes to 0
	passes := 0

	// while passes is less than the max number of passes
	for passes < maxPasses {
		// number of alphas changed is set to zero
		var numAlphasChanged = 0

		// for i in range of number of documents
		for i := 0; i < trn.numberOfDocuments; i++ {
			// calculate the error (classification value - training label)
			errorI := trn.error(i)

			// if label*error less than - tolerance and multiplier less than C
			// or label*error greater than tolerance and multiplier greater than 0
			if trn.y[i]*errorI < -trn.tolerance && trn.alpha[i] < trn.C ||
				trn.y[i]*errorI > trn.tolerance && trn.alpha[i] > 0 {

				// select an index randomly as long as it is different from the current i
				j := rand.Intn(trn.numberOfDocuments)

				for i == j {
					j = rand.Intn(trn.numberOfDocuments)
				}

				// calculate the error for this index
				errorJ := trn.error(j)

				// save the old value of the multipliers
				oldAlphaI := trn.alpha[i]
				oldAlphaJ := trn.alpha[j]

				// compute L and H
				L, H := trn.computeBounds(i, j)

				// if L and H are equal
				if L == H {
					// go on to the next iteration
					continue
				}

				// compute fancy n
				nParam := trn.computeNParam(i, j)

				// if fancy n greater than or equal to 0
				if nParam >= 0 {
					// go on to the next iteration
					continue
				}

				// compute new alpha using those one equations
				trn.alpha[j] = trn.clipAlpha(i, j, L, H, nParam)

				// if the alpha difference is less than 1e-5
				if math.Abs(trn.alpha[j]-oldAlphaJ) < 1e-5 {
					// go on to the next iteration
					continue
				}

				// compute new alpha i
				trn.alpha[i] = trn.updateAlpha(i, j, oldAlphaJ)

				// compute b1
				b1 := trn.computeB1(i, j, errorI, oldAlphaI, oldAlphaJ)

				// compute b2
				b2 := trn.computeB2(i, j, errorJ, oldAlphaI, oldAlphaJ)

				// compute b
				trn.b = trn.computeB(i, j, b1, b2)

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

func (trn Trainer) classify(i int) float64 {
	sum := 0.0

	for j := 0; j < trn.numberOfDocuments; j++ {
		sum += trn.alpha[j] * float64(trn.y[j]) * trn.kernel.Evaluate(trn.X[j], trn.X[i])
	}

	sum += trn.b

	if sum >= 0 {
		return 1
	}

	return -1

}

func (trn Trainer) error(index int) float64 {
	return trn.classify(index) - trn.y[index]
}

func (trn Trainer) computeBounds(i, j int) (float64, float64) {
	var L float64
	var H float64

	if trn.y[i] != trn.y[j] {
		diff := trn.alpha[j] - trn.alpha[i]
		L = math.Max(0, diff)
		H = math.Min(trn.C, trn.C+diff)
	} else {
		sum := trn.alpha[j] + trn.alpha[i]
		L = math.Max(0, sum-trn.C)
		H = math.Min(trn.C, sum)
	}

	return L, H
}

func (trn Trainer) computeNParam(i, j int) float64 {
	kij := trn.kernel.Evaluate(trn.X[i], trn.X[j])
	kii := trn.kernel.Evaluate(trn.X[i], trn.X[i])
	kjj := trn.kernel.Evaluate(trn.X[j], trn.X[j])

	return 2*kij - kii - kjj
}

func (trn Trainer) clipAlpha(i, j int, L, H, nParam float64) float64 {
	newAlpha := trn.alpha[j] - trn.y[j]*(trn.error(i)-trn.error(j))/nParam

	if newAlpha > H {
		return H
	} else if newAlpha < L {
		return L
	} else {
		return newAlpha
	}
}

func (trn Trainer) updateAlpha(i, j int, oldAlphaJ float64) float64 {
	return trn.alpha[i] + trn.y[i]*trn.y[j]*(oldAlphaJ-trn.alpha[j])
}

func (trn Trainer) computeB1(i, j int, errorI, oldAlphaI, oldAlphaJ float64) float64 {
	partOne := trn.y[i] * (trn.alpha[i] - oldAlphaI) * trn.kernel.Evaluate(trn.X[i], trn.X[i])
	partTwo := trn.y[j] * (trn.alpha[j] - oldAlphaJ) * trn.kernel.Evaluate(trn.X[i], trn.X[j])

	return trn.b - errorI - partOne - partTwo
}

func (trn Trainer) computeB2(i, j int, errorJ, oldAlphaI, oldAlphaJ float64) float64 {
	partOne := trn.y[i] * (trn.alpha[i] - oldAlphaI) * trn.kernel.Evaluate(trn.X[i], trn.X[j])
	partTwo := trn.y[j] * (trn.alpha[j] - oldAlphaJ) * trn.kernel.Evaluate(trn.X[j], trn.X[j])

	return trn.b - errorJ - partOne - partTwo
}

func (trn Trainer) computeB(i, j int, b1, b2 float64) float64 {
	if trn.alpha[i] > 0 && trn.alpha[i] < trn.C {
		return b1
	} else if trn.alpha[j] > 0 && trn.alpha[j] < trn.C {
		return b2
	} else {
		return (b1 + b2) / 2
	}
}
