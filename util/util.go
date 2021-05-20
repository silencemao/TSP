package util

import (
	"TSP/ioinfo"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func GetSampleData(pDataFileName string, pInit bool) []ioinfo.Data {
	data := make([]ioinfo.Data, 0)
	var d ioinfo.Data
	file, _ := GetFileLines("./data/" + pDataFileName)
	file = file[6 : len(file)-1]

	for _, s := range file {
		sp := strings.Split(s, " ")
		sp0, _ := strconv.Atoi(sp[0])
		sp1, _ := strconv.Atoi(sp[1])
		sp2, _ := strconv.Atoi(sp[2])
		d = new(ioinfo.BaseData)
		d.SetCityNum(sp0)
		d.SetPosX(float64(sp1))
		d.SetPosY(float64(sp2))
		data = append(data, d)
	}
	data = append(data, data[0])
	processMatrix(data)

	if pInit {

		pInRoute := make(map[ioinfo.Data]bool, 0)
		pInRoute[data[0]] = true

		var dataDst []ioinfo.Data
		dataDst = append(dataDst, data[0])
		for i := 1; i < len(data) - 1; i++ {
			pNearestInd := -1
			pNearestDis := math.MaxFloat64
			for j := 1; j < len(data) - 1; j++ {
				if pInRoute[data[j]] == false {
					if GetDistance(dataDst[len(dataDst) - 1], data[j]) < pNearestDis {
						pNearestInd = j
						pNearestDis = GetDistance(dataDst[len(dataDst) - 1], data[j])
					}
				}
			}
			pNext := data[pNearestInd]
			dataDst = append(dataDst, pNext)
			pInRoute[pNext] = true
		}
		dataDst = append(dataDst, data[len(data) - 1])
		return dataDst
	}
	return data
}

func processMatrix(data []ioinfo.Data) {
	distanceMatrix := make([][]float64, len(data) - 1)
	for i := 0; i < len(distanceMatrix); i++ {
		distanceMatrix[i] = make([]float64, len(data) - 1)
	}

	for i := 0; i < len(data) - 1; i++ {
		for j := 0; j < len(data) - 1; j++ {
			distanceMatrix[i][j] = math.Hypot(data[i].GetPosX() - data[j].GetPosX(), data[i].GetPosY() - data[j].GetPosY())
		}
	}
	for _, tData := range data {
		tData.SetDistanceMatrix(distanceMatrix)
	}
}

//得到这组排列总里程数
func GetResult(data []ioinfo.Data) float64 {
	pDistanceMatrix := data[0].GetDistanceMatrix()
	sum := 0.0
	for i := 0; i < len(data)-1; i++ {
		sum += pDistanceMatrix[data[i].GetCityNum() - 1][data[i+1].GetCityNum() - 1]
	}

	return sum
}

func GetDistance(d1 ioinfo.Data, d2 ioinfo.Data) float64 {
	return d1.DistanceTo(d2)
}

func GetFileLines(filePath string) ([]string, error) {
	var result []string
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file: %v error: %v", filePath, err)
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	return result, nil
}
