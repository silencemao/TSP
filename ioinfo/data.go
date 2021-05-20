package ioinfo

import "fmt"

type Data interface {

	String() string

	SetDistanceMatrix([][]float64)
	GetDistanceMatrix() [][]float64

	DistanceTo(Data) float64

	SetCityNum(int)
	GetCityNum() int

	SetPosX(float64)
	GetPosX() float64

	SetPosY(float64)
	GetPosY() float64
}

type BaseData struct {
	CityNum         int
	PosX            float64
	PosY            float64
	PDistanceMatrix [][]float64
}

func (d BaseData) String() string {
	return fmt.Sprintf("%2d ->", d.CityNum)
}

func (d *BaseData) SetCityNum(pNum int) {
	d.CityNum = pNum
}

func (d *BaseData) GetCityNum() int {
	return d.CityNum
}

func (d *BaseData) SetPosX(pX float64) {
	d.PosX = pX
}

func (d *BaseData) GetPosX() float64 {
	return d.PosX
}

func (d *BaseData) SetPosY(pY float64) {
	d.PosY = pY
}

func (d *BaseData) GetPosY() float64 {
	return d.PosY
}

func (d *BaseData) SetDistanceMatrix(pMatrix [][]float64) {
	d.PDistanceMatrix = pMatrix
}

func (d *BaseData) GetDistanceMatrix() [][]float64 {
	return d.PDistanceMatrix
}

func (d *BaseData) DistanceTo(d1 Data) float64 {
	return d.PDistanceMatrix[d.CityNum - 1][d1.GetCityNum() - 1]
}


