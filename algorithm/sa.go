package algorithm

import (
	"TSP/ioinfo"
	"TSP/util"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var Temp = 10000.0    //初始化温度
var TMin = 0.0001     //温度下界限
var Delta = 0.98     //温度下降率
var SearchTime = 500 //内循环次数

var R = rand.New(rand.NewSource(0))

func Sa(pFileName string) float64 {
	t := Temp
	data := util.GetSampleData(pFileName, false)

	s1 := time.Now().UnixNano()

	r := util.GetResult(data)
	if SearchTime > len(data) * len(data) {
		SearchTime = len(data) * len(data)
	}
	for t > TMin {
		for i := 0; i < SearchTime; i++ {
			temp := changCity(data)
			rn := util.GetResult(temp)
			if accept(rn-r, t) {
				data = temp
				r = rn
			}
		}
		t = t * Delta
	}
	s2 := time.Now().UnixNano()
	fmt.Println(data, s2 - s1)
	return util.GetResult(data)
}

//是否愿意接受新解
func accept(delta float64, temper float64) bool {
	if delta <= 0 {
		return true
	}
	return R.Float64() <= math.Exp((-delta)/temper)
}

func changCity(dataSrc []ioinfo.Data) []ioinfo.Data {
	pos1 := R.Intn(len(dataSrc)-3) + 1
	pos2 := R.Intn(len(dataSrc) - pos1-1) + pos1
	dataDest := make([]ioinfo.Data, len(dataSrc))
	copy(dataDest, dataSrc)

	r := R.Intn(2)
	if r == 0 {
		for pos1 < pos2 {
			dataDest[pos1], dataDest[pos2] = dataDest[pos2], dataDest[pos1]
			pos1 ++
			pos2 --
		}
	} else {
		dataDest[pos1], dataDest[pos2] = dataDest[pos2], dataDest[pos1]
	}

	return dataDest
}
