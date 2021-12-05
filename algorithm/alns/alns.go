package alns

import (
	"TSP/TSPSolution"
	"TSP/ioinfo"
	"TSP/util"
	"fmt"
	"math/rand"
	"sort"
)

type ALNS struct {
	destroySize, repairSize int
	destroyWeight, repairWeight []float32
	destroyCnt, repairCnt []int
	destroyScore, repairScore []float32

	theta1, theta2, theat3 float32 // 分数 优于最优，优于当前解，接受不优于当前解的分数
	alpha float32

	Solution *TSPSolution.TSPSolution
}

func (a *ALNS) Init() {
	rand.Seed(0)
	a.destroySize, a.repairSize = 2, 2
	a.destroyWeight, a.repairWeight = []float32{rand.Float32(), rand.Float32()}, []float32{rand.Float32(), rand.Float32()}
	a.destroyCnt, a.repairCnt = []int{0, 0}, []int{0, 0}
	a.destroyScore, a.repairScore = []float32{rand.Float32(), rand.Float32()}, []float32{rand.Float32(), rand.Float32()}
	a.theta1, a.theta2, a.theat3 = 20.0, 12.0, 8.0
	a.alpha = 0.9
	rand.Seed(1)
}

func (a *ALNS) Solve() (TSPSolution.TSPSolution, float64) {
/*
	初始解、当前score、最优score
	选择破坏算子
	选择修复算子
	更新分数、score、迭代次数
	更新算子的概率
*/

	score := a.Solution.GetLength()
	fmt.Println(score)
	curSolution, bestSolution := util.CopySolution(*a.Solution), util.CopySolution(*a.Solution)
	curScore, bestScore := score, score

	iterations := 1000

	iter := 0
	for iterations >= 0 {
		iterations -= 1

		// destroy
		destroyInd := util.RouletteSelect(a.destroyWeight)
		a.destroyCnt[destroyInd] += 1
		tmpSol, destroyList := a.Destroy(destroyInd, curSolution)

		// repair
		repairInd := util.RouletteSelect(a.repairWeight)
		a.repairCnt[repairInd] += 1
		newSol, newScore := a.Repair(repairInd, tmpSol, destroyList)

		if newScore < curScore {
			curSolution = util.CopySolution(newSol)
			curScore = util.GetResult(curSolution.GetPath())
			if newScore < bestScore {
				bestScore = newScore
				bestSolution = util.CopySolution(newSol)

				a.destroyScore[destroyInd] += a.theta1
				a.repairScore[repairInd] += a.theta1
				fmt.Println(bestSolution.GetPath(), bestScore, destroyInd, repairInd)
			} else {
				a.destroyScore[destroyInd] += a.theta2
				a.repairScore[repairInd] += a.theta2
			}
		} else {
			if rand.Float32() < 0.05 {
				curSolution = util.CopySolution(newSol)
				a.destroyScore[destroyInd] += a.theat3
				a.repairScore[repairInd] += a.theat3
			}
		}

		iter += 1

		//	更新参数
		if iter/20 == 0 {
			destroyWeightSum := float32(0.0)
			for i := range a.destroyWeight {
				if a.destroyCnt[i] == 0 {
					a.destroyWeight[i] *= a.alpha
				} else {
					a.destroyWeight[i] = a.destroyWeight[i] * a.alpha + (1-a.alpha) * a.destroyScore[i] / float32(a.destroyCnt[i])
				}
				destroyWeightSum += a.destroyWeight[i]
			}
			for i := range a.destroyWeight {
				a.destroyWeight[i] /= destroyWeightSum
			}

			repairWeightSum := float32(0.0)
			for i := range a.repairWeight {
				if a.repairCnt[i] == 0 {
					a.repairWeight[i] *= a.alpha
				} else {
					a.repairWeight[i] = a.repairWeight[i] * a.alpha + (1-a.alpha)*a.repairWeight[i] / float32(a.repairWeight[i])
				}
				repairWeightSum += a.repairWeight[i]
			}
			for i := range a.repairWeight {
				a.repairWeight[i] = a.repairWeight[i] / repairWeightSum
			}
		}
	}
	return bestSolution, bestScore
}

func (a *ALNS) Destroy(destroyInd int, sol TSPSolution.TSPSolution) (TSPSolution.TSPSolution, []ioinfo.Data) {
	curSolution := util.CopySolution(sol)
	var destroyList []ioinfo.Data

	mvNum := rand.Intn(20)
	if destroyInd == 0 { // 随机移除
		size := len(curSolution.GetPath())

		var inds []int

		for i := 0; i < mvNum; i++ {  // 记录被删除的元素
			inds = append(inds, rand.Intn(size))

			destroyList = append(destroyList, curSolution.GetPath()[inds[i]])
		}

		sort.Ints(inds)
		tmpPaths := curSolution.GetPath()
		for i := len(inds)-1; i >= 0; i-- {

		}
		curSolution.SetPath(tmpPaths)
	}

	if destroyInd == 1 {
		tmpPaths := curSolution.GetPath()
		distances := make([]int, len(tmpPaths)-1)
		set := make(map[int]float64)

		set[0] = tmpPaths[0].DistanceTo(tmpPaths[1]) + tmpPaths[len(tmpPaths)-2].DistanceTo(tmpPaths[0])
		set[len(distances)-1]  = tmpPaths[len(tmpPaths)-2].DistanceTo(tmpPaths[len(tmpPaths)-1]) + tmpPaths[len(tmpPaths)-3].DistanceTo(tmpPaths[len(tmpPaths)-2])
		distances[len(distances)-1] = len(distances)-1
		for i := 1; i < len(tmpPaths)-2; i++ {
			distances[i] =  i
			set[i] = tmpPaths[i-1].DistanceTo(tmpPaths[i]) + tmpPaths[i].DistanceTo(tmpPaths[i+1])
		}
		sort.Slice(distances, func(i, j int) bool {
			return set[distances[i]] > set[distances[j]]
		})

		inds := []int{}
		for i := 0; i < mvNum; i++ {
			inds = append(inds, distances[i])
			destroyList = append(destroyList, tmpPaths[distances[i]])
		}
		sort.Ints(inds)
		for i := len(inds)-1; i >= 0; i-- {
			tmpPaths = append(tmpPaths[:inds[i]], tmpPaths[inds[i]+1:]...)
		}
		curSolution.SetPath(tmpPaths)
	}
	return curSolution, destroyList

}

func (a *ALNS) Repair(repairInd int, sol TSPSolution.TSPSolution, destroyList []ioinfo.Data) (TSPSolution.TSPSolution, float64) {
	if repairInd == 0 { // 贪心插入
		for _, city := range destroyList {
			tmpPaths := sol.GetPath()

			bestInd := -1
			bestDis := float64(1<<31-1)
			for i := 1; i < len(tmpPaths)-1; i++ {
				disDiff := tmpPaths[i-1].DistanceTo(city) + city.DistanceTo(tmpPaths[i+1]) - tmpPaths[i].DistanceTo(tmpPaths[i+1])
				if disDiff < bestDis {
					bestDis = disDiff
					bestInd = i
				}
			}
			tmpPaths = append(append(tmpPaths[:bestInd], append([]ioinfo.Data{city}, tmpPaths[bestInd:]...)...))
			sol.SetPath(tmpPaths)
		}
	}

	if repairInd == 1 { // 扰动插入
		maxDis := 0.0
		tmpPaths := sol.GetPath()
		for _, c1 := range tmpPaths {
			for _, c2 := range tmpPaths {
				if c1.DistanceTo(c2) > maxDis {
					maxDis = c1.DistanceTo(c2)
				}
			}
		}
		for _, city := range destroyList {
			tmpPaths := sol.GetPath()

			bestInd := -1
			bestDis := float64(1<<31-1)
			for i := 1; i < len(tmpPaths)-1; i++ {
				disDiff := tmpPaths[i-1].DistanceTo(city) + city.DistanceTo(tmpPaths[i+1]) - tmpPaths[i].DistanceTo(tmpPaths[i+1]) + maxDis * rand.Float64()
				if disDiff < bestDis {
					bestDis = disDiff
					bestInd = i
				}
			}
			tmpPaths = append(append(tmpPaths[:bestInd], append([]ioinfo.Data{city}, tmpPaths[bestInd:]...)...))
			sol.SetPath(tmpPaths)
		}
	}
	distance := util.GetResult(sol.GetPath())
	sol.SetLength(distance)
	return sol, distance
}
