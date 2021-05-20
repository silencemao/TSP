package algorithm

import (
	"TSP/util"
	"fmt"
)

func dp(pFileName string) float64 {
	data := util.GetSampleData(pFileName, false)
	data = data[1: len(data)-1]
	//r := util.GetResult(data)

	M := 2^len(data)
	N := len(data)

	visited := make([]bool, len(data))
	for i := range visited {
		visited[i] = false
	}

	for i := 0; i < len(data); i++ {

		dis := make([][]float64, M)
		for j := range data {
			dis[j] = make([]float64, len(data))
		}
		for m := 0; m < M; m++ {
			for n := 0; n < N; n++ {
				dis[m][n] = -1
			}
		}

		s := 0
		init := 0
		fmt.Println(calculate(s, init, dis, visited))

	}

	return 0.0
}

func calculate(s, init int, dis [][]float64, visited []bool) float64 {
	if init==len(dis) {
		return 0
	}
	if dis[s][init] != -1 {
		return dis[s][init]
	}

	minLen := 1000.0
	for i := 0; i < len(dis); i++ {
		if !visited[i] {
			visited[i] = true
			temp := dis[s][i] + calculate(i, init+1, dis, visited)
			if temp < minLen {
				minLen = temp
			}
		}
	}
	dis[s][init] = minLen
	return dis[s][init]
}