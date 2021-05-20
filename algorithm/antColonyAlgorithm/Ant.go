package antColonyAlgorithm

import (
	"TSP/TSPSolution"
	"TSP/ioinfo"
	"TSP/util"
	"math"
	"math/rand"
)

type Ant struct {
	id int

	tVisitNodes []ioinfo.Data

	tUnVisitNodes []ioinfo.Data

	tSelectNodes []ioinfo.Data

	tSolution TSPSolution.TSPSolution

	tPheromone Pheromone

	tRand rand.Rand

	tTotalLength float64
}

func (a *Ant) SetId(pId int) {
	a.id = pId
}

func (a *Ant) GetId() int {
	return a.id
}

func (a *Ant) SetPheromone(pPheromone Pheromone) {
	a.tPheromone = pPheromone
}

func (a *Ant) GetPheromone() Pheromone {
	return a.tPheromone
}

func (a *Ant) InitNodes(pNodes []ioinfo.Data) {
	if len(pNodes) == 0 {
		panic("Ant:: city num is zero")
	}

	a.tVisitNodes, a.tUnVisitNodes, a.tSelectNodes = a.tVisitNodes[:0], a.tUnVisitNodes[:0], a.tSelectNodes[:0]
	a.tVisitNodes = append(a.tVisitNodes, pNodes[0])

	a.tUnVisitNodes = append(a.tUnVisitNodes, pNodes[1:]...)

	a.tSelectNodes = append(a.tSelectNodes, pNodes...)
}

func (a *Ant) SetRand(pRand rand.Rand) {
	a.tRand = pRand
}

func (a *Ant) GetRand() rand.Rand {
	return a.tRand
}

func (a *Ant) GetPath() []ioinfo.Data {
	return a.tVisitNodes
}

func (a *Ant) UpdateNodes(pInd int) {
	a.tVisitNodes = append(a.tVisitNodes, a.tUnVisitNodes[pInd])
	a.tUnVisitNodes = append(a.tUnVisitNodes[:pInd], a.tUnVisitNodes[pInd+1:]...)
}

func (a *Ant) NextNode() int {
	tTotalProbables := 0.0
	tCurNode := a.tVisitNodes[len(a.tVisitNodes) - 1]
	var tProbables []float64
	for _, tNext := range a.tUnVisitNodes {
		t := a.tPheromone.GetPheromoneMatrix()[tCurNode.GetCityNum()-1][tNext.GetCityNum()-1]
		n := 1.0 / util.GetDistance(tCurNode, tNext)
		p := math.Pow(t, a.tPheromone.GetAlpha()) * math.Pow(n, a.tPheromone.GetBeta())

		tProbables = append(tProbables, p)
		tTotalProbables += p
	}

	for i := 0; i < len(tProbables); i++ {
		tProbables[i] = tProbables[i] / tTotalProbables
	}

	pSelectP := a.tRand.Float64()

	pDistributionP := 0.0
	for i := 0; i < len(tProbables); i++ {
		pDistributionP += tProbables[i]

		if pSelectP < pDistributionP {
			return i
		}
	}

	panic("error ")
}

func (a *Ant) UpdatePheromone() {
	a.tVisitNodes = append(a.tVisitNodes, a.tVisitNodes[0])
	a.tTotalLength = util.GetResult(a.tVisitNodes)
	tPheromoneAdd := 1.0 / a.tTotalLength

	for i := 0; i < len(a.tVisitNodes) - 1; i++ {
		a.GetPheromone().Volatilize(a.tVisitNodes[i].GetCityNum() - 1, a.tVisitNodes[i+1].GetCityNum() - 1)
		a.GetPheromone().Add(a.tVisitNodes[i].GetCityNum() - 1, a.tVisitNodes[i+1].GetCityNum() - 1, tPheromoneAdd)
	}
}

func (a *Ant) Solve() bool {
	pNextInd := a.NextNode()
	a.UpdateNodes(pNextInd)

	if len(a.tUnVisitNodes) == 0 {
		return true
	} else {
		return false
	}
}



