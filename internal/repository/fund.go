package repository

import (
	"context"
	"log"

	"github.com/N30A/fondhav/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FundRepository struct {
	pool *pgxpool.Pool
}

func NewFundRepository(pool *pgxpool.Pool) *FundRepository {
	return &FundRepository{pool}
}

func (r *FundRepository) GetFundByISIN(ctx context.Context, ISIN string) (models.Fund, error) {
	query := `
		SELECT
			f.isin,
			f.name,
			f.benchmark_index,
			f.as_of_date,
			COUNT(fh.fund) AS holdings_count
		FROM funds f
		LEFT JOIN fund_holdings fh ON fh.fund = f.id
		WHERE f.isin = $1
		GROUP BY
			f.isin,
			f.name,
			f.benchmark_index,
			f.as_of_date;
	`

	var fund fund
	err := r.pool.QueryRow(ctx, query, ISIN).Scan(&fund.ISIN, &fund.Name, &fund.BenchmarkIndex, &fund.AsOfDate, &fund.HoldingsCount)
	if err != nil {
		log.Printf("unable to find a fund with ISIN: %s", ISIN)
		return models.Fund{}, err
	}

	return fund.toModel(), nil
}
