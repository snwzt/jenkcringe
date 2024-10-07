package factory

import "fmt"

type Airplane interface {
	SetSpeed(float64)
	GetSpeed() float64
}

type airbusA380 struct {
	speed float64
}

type f35 struct {
	speed float64
}

func (a *airbusA380) SetSpeed(s float64) {
	a.speed = s
}

func (a *airbusA380) GetSpeed() float64 {
	return a.speed
}

func (f *f35) SetSpeed(s float64) {
	f.speed = s
}

func (f *f35) GetSpeed() float64 {
	return f.speed
}

func NewPlane(plane string) (Airplane, error) {
	switch plane {
	case "AirbusA380":
		return &airbusA380{}, nil
	case "F35":
		return &f35{}, nil
	default:
		return nil, fmt.Errorf("plane does not exist: %s", plane)
	}
}
