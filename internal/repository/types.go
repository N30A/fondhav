package repository

import (
	"time"

	"github.com/N30A/fondhav/internal/models"
)

type fund struct {
	ISIN           string
	Name           string
	BenchmarkIndex *string
	AsOfDate       time.Time
	HoldingsCount  int
}

func (f fund) toModel(dbholdings []holding) models.Fund {
	benchmark := ""
	if f.BenchmarkIndex != nil {
		benchmark = *f.BenchmarkIndex
	}

	holdings := make([]models.Holding, len(dbholdings))
	for i, holding := range dbholdings {
		holdings[i] = holding.toModel()
	}

	return models.Fund{
		ISIN:           f.ISIN,
		Name:           f.Name,
		BenchmarkIndex: benchmark,
		AsOfDate:       f.AsOfDate,
		Holdings:       holdings,
		HoldingsCount:  f.HoldingsCount,
	}
}

type holding struct {
	Instrument    instrument
	WeightPercent float64
}

func (h holding) toModel() models.Holding {
	return models.Holding{
		Instrument:    h.Instrument.toModel(),
		WeightPercent: h.WeightPercent,
	}
}

type instrument struct {
	ISIN *string
	Name string
}

func (i instrument) toModel() models.Instrument {
	isin := ""
	if i.ISIN != nil {
		isin = *i.ISIN
	}
	return models.Instrument{
		ISIN: isin,
		Name: i.Name,
	}
}

type PaginatedList[T any] struct {
	Total int
	Items []T
}

type Pagination struct {
	Limit  int
	Offset int
}

type SortDirection string

const (
	SortAsc  SortDirection = "ASC"
	SortDesc SortDirection = "DESC"
)

type Sorting struct {
	Column    string
	Direction SortDirection
}
