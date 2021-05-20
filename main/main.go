package main

import (
	algo "TSP/algorithm"
	"TSP/algorithm/antColonyAlgorithm/algorithm"
	"TSP/algorithm/genericAlgorithm"
	"TSP/test"
	"TSP/util"
	"fmt"
)

var DataFileName = "3.eil76.tsp.txt"
//var DataFileName = "5.att10.tsp.txt"

func testSa() {

	//fmt.Printf("模拟退火启动:\n文件名称:%s\n初始化温度:%.1f\n温度下界限:%f\n温度下降率:%f\n内循环次数:%d\n最短路径:", DataFileName, Temp, TMin, Delta, SearchTime)
	//开始模拟退火
	fmt.Println("模拟退火启动：")
	r := algo.Sa(DataFileName)
	fmt.Println(r)

	//best := test.GetBest(DataFileName)
	//fmt.Printf("\n退火解为: %7.3f ,最优解为: %7.3f ,误差率为: %4.2f", r, best, (r-best)/best*100)
	//fmt.Println("%")
}

func testLa() {
	fmt.Println("\r\n延迟接受启动：")
	r := algo.La(DataFileName)
	fmt.Println(r)
	//best := test.GetBest(DataFileName)
	//fmt.Printf("\n延迟接受解为: %7.3f ,最优解为: %7.3f ,误差率为: %4.2f", r, best, (r-best)/best*100)
	//fmt.Println("%")
}

func testGa() {
	fmt.Println("\r\n遗传算法启动：")
	data := util.GetSampleData(DataFileName, false)
	pGa := new(genericAlgorithm.GenericAlgorithm)
	pGa.Init(40, 42, data)
	pGa.Evolution()
	pGa.Evaluate()
}

func testACO() {
	fmt.Println("\r\n蚁群算法启动：")
	data := util.GetSampleData(DataFileName, false)
	data = data[:len(data) - 1]

	pAntNums := 20
	pIterNum := 1000
	l := float64(pAntNums) / 800.0

	tPheromoneMatrix := make([][]float64, len(data))
	for i := range tPheromoneMatrix {
		tPheromoneMatrix[i] = make([]float64, len(data))
	}
	for i := 0; i < len(tPheromoneMatrix); i++ {
		for j := 0; j < len(tPheromoneMatrix); j++ {
			tPheromoneMatrix[i][j] = l
		}
	}

	pAco := new(algorithm.AntColony)
	pAco.Init(pAntNums, tPheromoneMatrix, data)
	pAco.SetIterNum(pIterNum)
	pAco.Solve()
}

func main() {
	testSa()
	testLa()
	testGa()
	testACO()

	fmt.Println(test.GetBest(DataFileName))
	//test.Permutation(DataFileName)
}
