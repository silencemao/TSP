package TSPSolution

import (
	"TSP/ioinfo"
)
//var DataFileName = "1.att48.tsp.txt"

//type Data struct {
//	CityNum int
//	PosX    float64
//	PosY    float64
//}
//
//func (d *Data) GetCityNum() int {
//	return d.CityNum
//}
//
//func (d Data) String() string {
//	return fmt.Sprintf("%2d ->", d.CityNum)
//}
//
//var data []Data
//
//func ReadData() []Data {
//	return data
//}
//
//func InitSampleData() []Data {
//	data = make([]Data, 0)
//	var d Data
//	file, _ := GetFileLines("./data/" + DataFileName)
//	file = file[6 : len(file)-1]
//	for _, s := range file {
//		sp := strings.Split(s, " ")
//		sp0, _ := strconv.Atoi(sp[0])
//		sp1, _ := strconv.Atoi(sp[1])
//		sp2, _ := strconv.Atoi(sp[2])
//		d = Data{sp0, float64(sp1), float64(sp2)}
//		data = append(data, d)
//	}
//	data = append(data, data[0])
//	return data
//}
//
//func GetFileLines(filePath string) ([]string, error) {
//	var result []string
//	b, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		fmt.Printf("read file: %v error: %v", filePath, err)
//		return result, err
//	}
//	s := string(b)
//	for _, lineStr := range strings.Split(s, "\n") {
//		lineStr = strings.TrimSpace(lineStr)
//		if lineStr == "" {
//			continue
//		}
//		result = append(result, lineStr)
//	}
//	return result, nil
//}
//
//func GetResult(data []Data) float64 {
//	sum := 0.0
//	for i := 0; i < len(data)-1; i++ {
//		sum += GetDistance(data[i], data[i+1])
//	}
//	return sum
//}
//
//func GetDistance(d1, d2 Data) float64 {
//	return math.Hypot(d1.PosX-d2.PosX, d1.PosY-d2.PosY)
//}



type TSPSolution struct {
	tLength float64
	tPath   []ioinfo.Data
	tProbability float64
}

func (t *TSPSolution) SetLength(pLength float64) {
	t.tLength = pLength
}

func (t *TSPSolution) GetLength() float64 {
	return t.tLength
}

func (t *TSPSolution) SetPath(pPath []ioinfo.Data) {
	t.tPath = pPath
}

func (t *TSPSolution) GetPath() []ioinfo.Data {
	return t.tPath
}

func (t *TSPSolution) SetProbability(pProbability float64) {
	t.tProbability = pProbability
}

func (t *TSPSolution) GetProbability() float64 {
	return t.tProbability
}


