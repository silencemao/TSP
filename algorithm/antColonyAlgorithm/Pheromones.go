package antColonyAlgorithm

type Pheromone interface {

	SetAlpha(float64)
	GetAlpha() float64

	SetBeta(float64)
	GetBeta() float64

	SetRatio(float64)
	GetRatio() float64

	SetPheromoneMatrix([][]float64)
	GetPheromoneMatrix() [][]float64

	Volatilize(int, int)
	Add(int, int, float64)
}

type BasePheromones struct {

	tPheromoneMatrix [][]float64
	tRatio float64
	tAlpha float64
	tBeta float64
}

func (p *BasePheromones) SetRatio(pRatio float64) {
	p.tRatio = pRatio
}

func (p *BasePheromones) GetRatio() float64 {
	return p.tRatio
}

func (p *BasePheromones) SetAlpha(pAlpha float64) {
	p.tAlpha = pAlpha
}

func (p *BasePheromones) GetAlpha() float64 {
	return p.tAlpha
}

func (p *BasePheromones) SetBeta(pBeta float64) {
	p.tBeta = pBeta
}

func (p *BasePheromones) GetBeta() float64 {
	return p.tBeta
}

func (p *BasePheromones) SetPheromoneMatrix(pMatrix [][]float64) {
	p.tPheromoneMatrix = pMatrix
}

func (p *BasePheromones) GetPheromoneMatrix() [][]float64 {
	return p.tPheromoneMatrix
}

func (p *BasePheromones) Volatilize(i, j int) {
	p.tPheromoneMatrix[i][j] = p.tPheromoneMatrix[i][j] * (1 - p.tRatio)
}

func (p *BasePheromones) Add(i, j int, pPheromone float64) {
	p.tPheromoneMatrix[i][j] += pPheromone
}
