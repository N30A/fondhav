package repository

import (
	"reflect"
	"testing"
	"time"

	"github.com/N30A/fondhav/internal/models"
)

func TestInstrumentToModel(t *testing.T) {
	isin := "SE0015658109"
	instrument := instrument{
		ISIN: &isin,
		Name: "Epiroc A",
	}

	got := instrument.toModel()

	want := models.Instrument{
		ISIN: "SE0015658109",
		Name: "Epiroc A",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestInstrumentToModelNilISIN(t *testing.T) {
	instrument := instrument{
		ISIN: nil,
		Name: "Epiroc A",
	}

	got := instrument.toModel()

	want := models.Instrument{
		ISIN: "",
		Name: "Epiroc A",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestHoldingToModel(t *testing.T) {
	isin := "SE0015658109"
	holding := holding{
		Instrument: instrument{
			ISIN: &isin,
			Name: "Epiroc A",
		},
		WeightPercent: 5.0,
	}

	got := holding.toModel()

	want := models.Holding{
		Instrument: models.Instrument{
			ISIN: "SE0015658109",
			Name: "Epiroc A",
		},
		WeightPercent: 5.0,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestFundToModel(t *testing.T) {
	isin := "SE0015658109"
	holdings := []holding{
		{
			Instrument: instrument{
				ISIN: &isin,
				Name: "Epiroc A",
			},
			WeightPercent: 5.0,
		},
	}

	benchmark := "SIX Portfolio Return Index"
	fund := fund{
		ISIN:           "SE0014991535",
		Name:           "PLUS Allabolag Sverige Index",
		BenchmarkIndex: &benchmark,
		AsOfDate:       time.Date(2026, 04, 26, 12, 0, 0, 0, time.UTC),
	}

	got := fund.toModel(holdings)

	want := models.Fund{
		ISIN:           "SE0014991535",
		Name:           "PLUS Allabolag Sverige Index",
		BenchmarkIndex: "SIX Portfolio Return Index",
		AsOfDate:       time.Date(2026, 04, 26, 12, 0, 0, 0, time.UTC),
		Holdings: []models.Holding{
			{
				Instrument: models.Instrument{
					ISIN: "SE0015658109",
					Name: "Epiroc A",
				},
				WeightPercent: 5.0,
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestFundToModelNoHoldings(t *testing.T) {
	holdings := []holding{}

	benchmark := "SIX Portfolio Return Index"
	fund := fund{
		ISIN:           "SE0014991535",
		Name:           "PLUS Allabolag Sverige Index",
		BenchmarkIndex: &benchmark,
		AsOfDate:       time.Date(2026, 04, 26, 12, 0, 0, 0, time.UTC),
	}

	got := fund.toModel(holdings)

	want := models.Fund{
		ISIN:           "SE0014991535",
		Name:           "PLUS Allabolag Sverige Index",
		BenchmarkIndex: "SIX Portfolio Return Index",
		AsOfDate:       time.Date(2026, 04, 26, 12, 0, 0, 0, time.UTC),
		Holdings:       []models.Holding{},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}
