package algorithm

import (
	"TSP/ioinfo"
	"TSP/util"
	"fmt"
	"math/rand"
	"time"
)

type LateAcceptance struct {
	n       int
	tScore []float64
}

func (l *LateAcceptance) Init(n int, pStartScore float64) {
	l.n = n + 1

	for i := 0; i < l.n; i++ {
		l.tScore = append(l.tScore, pStartScore)
	}
}

func (l *LateAcceptance) Accept(pScore float64) bool {
	var pAccept bool

	if pScore <= l.tScore[0] {
		pAccept = true
	} else if pScore <= l.tScore[len(l.tScore) - 1] {
		pAccept = true
	} else {
		pAccept = false
	}
	if pAccept {
		l.tScore = append(l.tScore, pScore)
	} else {
		l.tScore = append(l.tScore, l.tScore[len(l.tScore) - 1])
	}
	if len(l.tScore) >= l.n {
		l.tScore = l.tScore[1:]
	}

	return pAccept
}


func changCity1(dataSrc []ioinfo.Data, R1 rand.Rand) []ioinfo.Data {
	pos1 := R1.Intn(len(dataSrc)-3) + 1
	pos2 := R1.Intn(len(dataSrc)-pos1-1) + pos1
	dataDest := make([]ioinfo.Data, len(dataSrc))
	copy(dataDest, dataSrc)
	r := R1.Intn(2)

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

func La(pFileName string) float64 {
	data := util.GetSampleData(pFileName, false)
	r := util.GetResult(data)

	pLa := new(LateAcceptance)
	pLa.Init(215, r)

	var R1 = rand.New(rand.NewSource(0))

	s1 := time.Now().UnixNano()
	pIterNum := 200000
	for pIterNum > 0 {
		temp := changCity1(data, *R1)
		rn := util.GetResult(temp)
		if pLa.Accept(rn) {
			data = temp
		}
		pIterNum--
	}
	s2 := time.Now().UnixNano()
	fmt.Println(data, s2 -s1)
	return util.GetResult(data)
}
