package genericAlgorithm

import (
	"TSP/TSPSolution"
	"TSP/ioinfo"
	"TSP/util"
	"fmt"
	"math/rand"
	"time"
)

type GenericAlgorithm struct {
	tGroupNum int
	tSonNum   int
	tRand     *rand.Rand

	tGroupSolutions   []TSPSolution.TSPSolution
	tSonSolutions     []TSPSolution.TSPSolution

	tCopulation   float64  // 杂交概率
	tInheritance  float64  // 变异概率
}

func (g *GenericAlgorithm) Init(pGroupNum, pSonNum int, pInitSolution []ioinfo.Data) {
	g.tGroupNum = pGroupNum
	g.tSonNum = pSonNum

	g.tRand = rand.New(rand.NewSource(0))

	g.tCopulation = 0.8
	g.tInheritance = 0.2

	for i := 0; i < g.tGroupNum; i++ {
		tSolution := new(TSPSolution.TSPSolution)
		pos1 := g.tRand.Intn(len(pInitSolution) - 3) + 1
		pos2 := g.tRand.Intn(len(pInitSolution) - pos1 - 1) + pos1
		for pos1==pos2 {
			pos2 = g.tRand.Intn(len(pInitSolution) - pos1 - 1) + pos1
		}
		dataDst := make([]ioinfo.Data, len(pInitSolution))
		copy(dataDst, pInitSolution)
		for pos1 < pos2 {
			dataDst[pos1], dataDst[pos2] = dataDst[pos2], dataDst[pos1]
			pos1++
			pos2--
		}

		tSolution.SetPath(dataDst)
		tSolution.SetLength(util.GetResult(tSolution.GetPath()))
		g.tGroupSolutions = append(g.tGroupSolutions, *tSolution)
	}
	g.CalculateProbability(g.tGroupSolutions)
}

func (g *GenericAlgorithm) Evolution() {
	pIter := 5000

	s1 := time.Now().UnixNano()
	for pIter > 0 {

		g.tSonSolutions = g.tSonSolutions[:0]

		// 交叉
		pInd1 := g.EvolutionSelect()
		pInd2 := g.EvolutionSelect()
		for pInd1 == pInd2 {
			pInd2 = g.EvolutionSelect()
		}

		pFatherSolution, pMotherSolution := g.tGroupSolutions[pInd1], g.tGroupSolutions[pInd2]

		pCrossIter := g.tGroupNum - g.tGroupNum / 2
		for pCrossIter > 0 {
			if g.tRand.Float64() > g.tCopulation {

			} else {
				g.EvolutionCross(pFatherSolution, pMotherSolution)
				pCrossIter--
			}
		}

		// 变异
		for i := 0; i < len(g.tSonSolutions); i++ {
			if g.tRand.Float64() < g.tInheritance {
				g.EvolutionVariation(i)
			}
		}

		for i := 0; i < len(g.tSonSolutions); i++ {
			CheckPath(g.tSonSolutions[i])
		}

		g.SonSolutionParam()

		g.CalculateProbability(g.tSonSolutions)

		g.EvolutionGroup()

		//g.Evaluate()

		pIter--
	}
	s2 := time.Now().UnixNano()
	fmt.Println(s2 - s1)
}

func (g *GenericAlgorithm) SonSolutionParam() {
	for i := 0; i < len(g.tSonSolutions); i++ {
		g.tSonSolutions[i].SetLength(util.GetResult(g.tSonSolutions[i].GetPath()))
	}
}

func (g *GenericAlgorithm) EvolutionGroup() {
	for i := 0; i < len(g.tSonSolutions) - 1; i++ {
		for j := 0; j < len(g.tSonSolutions) - 1 - i; j++ {
			if g.tSonSolutions[j].GetLength() > g.tSonSolutions[j+1].GetLength() {
				g.tSonSolutions[j], g.tSonSolutions[j+1] = g.tSonSolutions[j+1], g.tSonSolutions[j]
			}
		}
	}

	for i := 0; i < len(g.tSonSolutions); i++ {
		for j := 0; j < g.tGroupNum; j++ {
			if g.tGroupSolutions[j].GetLength() > g.tSonSolutions[i].GetLength() {
				g. tGroupSolutions[j] = g.tSonSolutions[i]
				break
			}
		}
	}
}
// 计算概率很关键，决定了下一次迭代选择的父亲与母亲
func (g *GenericAlgorithm) CalculateProbability(pSolutions []TSPSolution.TSPSolution) {
	pTotalP := 0.0
	pTotalLength := 0.0
	for i := 0; i < len(pSolutions); i++ {
		pTotalLength += pSolutions[i].GetLength()
	}

	for i := 0; i < len(pSolutions); i++ {
		pSolutions[i].SetProbability((1.0 / pSolutions[i].GetLength()) * pTotalLength)
		pTotalP += pSolutions[i].GetProbability()
	}

	for i := 0; i < len(pSolutions); i++ {
		pSolutions[i].SetProbability(pSolutions[i].GetProbability() / pTotalP)
	}
}

func (g *GenericAlgorithm) EvolutionSelect() int {
	pSelectPro := g.tRand.Float64()

	pDistribution := 0.0

	for i := 0; i < g.tGroupNum; i++ {
		pDistribution += g.tGroupSolutions[i].GetProbability()
		if pSelectPro < pDistribution {
			return i
		}
	}
	return g.tGroupNum - 1
}

func (g *GenericAlgorithm) EvolutionCross(pFatherSolution, pMotherSolution TSPSolution.TSPSolution) {
	pCrossI := g.tRand.Intn(len(pFatherSolution.GetPath()) - 3) + 1
	pCrossJ := g.tRand.Intn(len(pFatherSolution.GetPath()) - 1 - pCrossI) + 1

	if pCrossI > pCrossJ {
		pCrossI, pCrossJ = pCrossJ, pCrossI
	}

	pFatherS := copySolution(pFatherSolution)
	pMotherS := copySolution(pMotherSolution)

	for i := pCrossI; i <= pCrossJ; i++ {
		pFatherS.GetPath()[i], pMotherS.GetPath()[i] = pMotherS.GetPath()[i], pFatherS.GetPath()[i]
	}
	_, pInd1 := g.findConflict(pFatherS, pMotherS, pCrossI, pCrossJ)
	_, pInd2 := g.findConflict(pMotherS, pFatherS, pCrossI, pCrossJ)
	for i := 0; i < len(pInd1); i++ {
		pFatherS.GetPath()[pInd1[i]], pMotherS.GetPath()[pInd2[i]] = pMotherS.GetPath()[pInd2[i]], pFatherS.GetPath()[pInd1[i]]
	}

	g.tSonSolutions = append(g.tSonSolutions, pFatherS)
	g.tSonSolutions = append(g.tSonSolutions, pMotherS)
}

func (g *GenericAlgorithm) findConflict(pFatherSolution, pMotherSolution TSPSolution.TSPSolution, pCrossI, pCrossJ int) ([]ioinfo.Data, []int) {
	pFatherSelect := pFatherSolution.GetPath()[pCrossI: pCrossJ+1]
	pMotherSelect := pMotherSolution.GetPath()[pCrossI: pCrossJ+1]
	//fmt.Println(pFatherSelect)
	//fmt.Println(pMotherSelect)
	//fmt.Println()

	var pConflict []ioinfo.Data
	for _, pDataI := range pFatherSelect {
		var isOK bool
		for _, pDataJ := range pMotherSelect {
			if pDataI.GetCityNum() == pDataJ.GetCityNum() {
				isOK = true
			}
		}
		if !isOK {
			pConflict = append(pConflict, pDataI)
		}
	}

	var pRes []ioinfo.Data
	var pInd []int
	for i := 0; i < pCrossI; i++ {
		for _, pData := range pConflict {
			if pFatherSolution.GetPath()[i].GetCityNum() == pData.GetCityNum() {
				pRes = append(pRes, pData)
				pInd = append(pInd, i)
			}
		}
	}

	for i := pCrossJ+1; i < len(pFatherSolution.GetPath()) - 1; i++ {
		for _, pData := range pConflict {
			if pFatherSolution.GetPath()[i].GetCityNum() == pData.GetCityNum() {
				pRes = append(pRes, pData)
				pInd = append(pInd, i)
			}
		}
	}
	return pRes, pInd
}

func (g *GenericAlgorithm) EvolutionVariation(pInd int) {
	pSelectSolution := g.tSonSolutions[pInd]

	pos1 := g.tRand.Intn(len(pSelectSolution.GetPath()) - 3) + 1
	pos2 := g.tRand.Intn(len(pSelectSolution.GetPath()) - pos1 - 1) + pos1

	pVariableSolution := TSPSolution.TSPSolution{}
	dataDest := make([]ioinfo.Data, len(pSelectSolution.GetPath()))
	copy(dataDest, pSelectSolution.GetPath())

	r := g.tRand.Intn(2)
	if r == 0 {
		for pos1 < pos2 {
			dataDest[pos1], dataDest[pos2] = dataDest[pos2], dataDest[pos1]
			pos1 ++
			pos2 --
		}
	} else {
		dataDest[pos1], dataDest[pos2] = dataDest[pos2], dataDest[pos1]
	}

	pVariableSolution.SetPath(dataDest)
	pVariableSolution.SetLength(util.GetResult(dataDest))

	g.tSonSolutions = append(g.tSonSolutions, pVariableSolution)
}

//得到这组排列总里程数
//func (g *GenericAlgorithm) GetResult(data []ioinfo.Data) float64 {
//	sum := 0.0
//	for i := 0; i < len(data) - 1; i++ {
//		sum += g.GetDistance(data[i], data[i+1])
//	}
//	return sum
//}

//两个城市间的距离 Map
//func (g *GenericAlgorithm) GetDistance(d1 ioinfo.Data, d2 ioinfo.Data) float64 {
//	return math.Hypot(d1.GetPosX()-d2.GetPosX(), d1.GetPosY()-d2.GetPosY())
//}

func copySolution(pSolution TSPSolution.TSPSolution) TSPSolution.TSPSolution {
	pNewSolution := TSPSolution.TSPSolution{}
	dataDst := make([]ioinfo.Data, len(pSolution.GetPath()))
	copy(dataDst, pSolution.GetPath())
	pNewSolution.SetPath(dataDst)
	pNewSolution.SetLength(pSolution.GetLength())
	return pNewSolution
}

func (g *GenericAlgorithm) Evaluate() TSPSolution.TSPSolution {
	bestInd := 0
	for i := 0; i < g.tGroupNum; i++ {
		if g.tGroupSolutions[i].GetLength() < g.tGroupSolutions[bestInd].GetLength() {
			bestInd = i
		}
	}
	fmt.Println("best solution : \n", g.tGroupSolutions[bestInd].GetPath())
	fmt.Println("best length : ", g.tGroupSolutions[bestInd].GetLength())
	fmt.Println()
	return g.tGroupSolutions[bestInd]
}

func CheckPath(pSolution TSPSolution.TSPSolution) {
	for i := 1; i < len(pSolution.GetPath()); i++ {
		for j := i+1; j < len(pSolution.GetPath()); j++ {
			if pSolution.GetPath()[i].GetCityNum() == pSolution.GetPath()[j].GetCityNum() {
				panic(" Solution 中存在重复的城市 ")
			}
		}
	}
}

