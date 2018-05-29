package model

type Trainer struct {
	kernel kernel.Kernel
	C int
	w []float64
	b float64
	m int
	n int
	alphas []float64
	tolerance float64
	X []float64
	y []int
}

func NewTrainer() Trainer {
	// create a new trainer
}

func (trn Trainer) train() {
	// initialize alphas and b with 0
	// initialize passes to 0
	// while passes is less than the max number of passes
		// number of alphas changed is set to zero
		// for i in range of number of documents
			// calculate the error (classification value - training label)
			// if label*error less than - tolerance and multiplier less than C
			//		or label*error greater than tolerance and multiplier greater than 0
				// select an index randomly as long as it is different from the current i
				// calculate the error for this index
				// save the old value of the multipliers
				// compute L and H
				// if L and H are equal 
					// go on to the next iteration
				// compute fancy n
				// if fancy n greater than or equal to 0
					// go on to the next iteration
				// compute new alpha using those one equations
}

func (trn Trainer) classify(i int) int {
	sum := 0

	for j:= 0; j < m; j++ {
		sum += alphas[j] * y[j] * kernel.Evaluate(X[j], X[i])
	}

	sum += b

	if sum >= 0 {
		return 1
	}
	else {
		return -1
	}
}

func (trn Trainer) error(index int) float64 {
	return classify(index) - y[index]
}

func (trn Trainer) computeBounds(i, j int) float64, float64 {
	var L float64
	var H float64

	if y[i] != y[j] {
		diff := alpha[j] - alpha[i]
		L = math.Max(0, diff)
		H = math.Min(C, C + diff)
	}
	else {
		sum := alpha[j] + alpha[i]
		L = math.Max(0, sum - C)
		H = math.Min(C, sum)
	}

	return L, H
}

func (trn Trainer) computeNParam(i, j int) float64 {
	kij := kernel.Evaluate(X[i], X[j])
	kii := kernel.Evaluate(X[i], X[i])
	kjj := kernel.Evaluate(X[j], X[j])

	return 2 * kij - kii - kj
}

func (trn Trainer) clipAlpha(i, j int, L, H, n_param float64) float64 {
	newAlpha := alphas[j] - y[j] * (error(i) - error(j)) / n_param

	if newAlpha > H {
		return H
	}
	else if newAlpha < L {
		return L
	}
	else {
		return newAlpha
	}
}