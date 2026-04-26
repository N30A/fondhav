package models

import (
	"time"
)

type Fund struct {
	ISIN           string
	Name           string
	BenchmarkIndex string
	AsOfDate       time.Time
	Holdings       []Holding
	HoldingsCount  int
}

type Holding struct {
	Instrument    Instrument
	WeightPercent float64
}

type Instrument struct {
	ISIN string
	Name string
}
