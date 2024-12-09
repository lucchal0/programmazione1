package rettangolo

type rettangolo struct {
	base, altezza float64
}

func NewRettangolo(b, h float64) *rettangolo {
	return &rettangolo{b, h}
}

func Area(r rettangolo) float64 {
	return r.base * r.altezza
}
