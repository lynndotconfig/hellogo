package main

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Car
}

type Car interface {
	Drive() error
	Stop() error
}


// 使用示例
// assembly := car.NewBuilder().Paint(car.RedColor)
//
// familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
// familyCar.Drive()
//
// sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
// sportsCar.Drive()