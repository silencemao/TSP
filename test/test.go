package test

import (
	"TSP/ioinfo"
	"TSP/util"
	"fmt"
	"strconv"
	"strings"
)

var cityNum []int

func GetBest(pDataFileName string) float64 {
	InitTestData(pDataFileName)
	data := util.GetSampleData(pDataFileName, false)
	dataDest := make([]ioinfo.Data, 0)
	for _, i := range cityNum {
		for _, d := range data {
			if d.GetCityNum() == i {
				dataDest = append(dataDest, d)
			}
		}
	}
	fmt.Println("\r\n best result ")
	fmt.Println(dataDest)
	return util.GetResult(dataDest)
}

func InitTestData(pDataFileName string) {
	cityNum = cityNum[:0]
	data := make([]int, 0)
	str := strings.Split(pDataFileName, ".")
	file, _ := util.GetFileLines("./data/" + str[0] + "." + str[1] + ".opt.tour.txt")
	file = file[5 : len(file)-1]
	for _, s := range file {
		i, _ := strconv.Atoi(s)
		data = append(data, i)
	}
	cityNum = append(data[:len(data)-1], data[0])
}

func HeapPermutation(a []ioinfo.Data, size int, b []ioinfo.Data, best []ioinfo.Data) {
	if size == 1 {
		c := make([]ioinfo.Data, len(a))
		copy(c, a)
		c = append([]ioinfo.Data{b[0]}, c...)
		c = append(c, b[len(b)-1])
		if util.GetResult(c) < util.GetResult(best) {
			copy(best, c)
			fmt.Println(best)
		}
	}

	for i := 0; i < size; i++ {
		HeapPermutation(a, size-1, b, best)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func nextPermutation(data []ioinfo.Data) bool {
	var i int
	var flag bool = false
	for i = len(data) - 2; i >= 0; i-- {
		if data[i].GetCityNum() < data[i+1].GetCityNum() {
			flag = true
			break
		}
	}

	if flag {
		var j int
		for j = len(data) - 1; j > i; j-- {
			if data[j].GetCityNum() > data[i].GetCityNum() {
				break
			}
		}
		data[i], data[j] = data[j], data[i]
	}
	i++
	for j := len(data) - 1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return flag
}

func Permutation(pDataFileName string) {
	data := util.GetSampleData(pDataFileName, false)

	pBestLength := util.GetResult(data)
	pBestRes := make([]ioinfo.Data, len(data))

	pFirst := data[0]
	data = data[1:len(data) - 1]
	pRes := make([]ioinfo.Data, len(data))

	for nextPermutation(data) {
		pRes = pRes[:0]
		pRes = append([]ioinfo.Data{pFirst}, data...)
		pRes = append(pRes, pFirst)
		if pBestLength >= util.GetResult(pRes) {
			pBestLength = util.GetResult(pRes)
			copy(pBestRes, pRes)
			fmt.Println(pBestRes, util.GetResult(pRes))
		}
	}
}
