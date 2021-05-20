package algorithm

import (
	"TSP/algorithm/antColonyAlgorithm"
	"TSP/ioinfo"
	"TSP/util"
	"fmt"
	"math"
	"math/rand"
	"time"
)
// https://www.cnblogs.com/asxinyu/p/Path_Optimization_Tsp_Problem_Ant_System_CSharp.html#opennewwindow
// https://blog.csdn.net/georgesale/article/details/80466909
type AntColony struct {
	tAnts []*antColonyAlgorithm.Ant

	tIterNum int

	tNodes []ioinfo.Data
}

func (a *AntColony) Init(pAntNum int, pPheromoneMatrix [][]float64, pNodes []ioinfo.Data) {
	a.tNodes = pNodes

	tPheromone := new(antColonyAlgorithm.BasePheromones)
	tPheromone.SetPheromoneMatrix(pPheromoneMatrix)
	tPheromone.SetAlpha(0.00001)
	tPheromone.SetBeta(13)
	tPheromone.SetRatio(0.5)

	tRand := rand.New(rand.NewSource(0))

	for i := 0; i < pAntNum; i++ {
		tAnt := new(antColonyAlgorithm.Ant)

		tAnt.SetPheromone(tPheromone)
		tAnt.SetId(i)
		tAnt.InitNodes(pNodes)

		tAnt.SetRand(*tRand)
		a.tAnts = append(a.tAnts, tAnt)
	}
}

func (a *AntColony) SetIterNum(pIterNum int) {
	a.tIterNum = pIterNum
}

func (a *AntColony) GetIterNum() int {
	return a.tIterNum
}

func (a *AntColony) GetAnts() []*antColonyAlgorithm.Ant {
	return a.tAnts
}

func (a *AntColony) Solve() {
	s1 := time.Now().UnixNano()
	var bestLength = math.MaxFloat64
	var bestRes []ioinfo.Data
	for a.tIterNum > 0 {

		var pFinish bool
		for true {
			for _, tAnt := range a.GetAnts() {
				pFinish = tAnt.Solve()
			}
			if pFinish {
				break
			}
		}

		for _, tAnt := range a.tAnts {
			tPath := append(tAnt.GetPath(), tAnt.GetPath()[0]) // 在最后加上起点，表示回到起点
			if util.GetResult(tPath) < bestLength {
				bestLength = util.GetResult(tPath)
				bestRes = tPath
			}
		}

		for _, tAnt := range a.GetAnts() {
			// 更新信息素
		   tAnt.UpdatePheromone()
		   // 重新更新Ant的node
		   tAnt.InitNodes(a.tNodes)
		}

		a.tIterNum--
	}
	s2 := time.Now().UnixNano()
	fmt.Println(bestRes, bestLength, s2 - s1)
}

