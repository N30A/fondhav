CREATE TABLE IF NOT EXISTS funds (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    isin CHAR(12) UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    benchmark_index TEXT NULL,
    as_of_date DATE NOT NULL 
);

CREATE INDEX funds_benchmark_idx ON funds(benchmark_index); 

CREATE TABLE IF NOT EXISTS instruments (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    isin CHAR(12) NULL,
    name TEXT NOT NULL
);

CREATE UNIQUE INDEX instruments_unique_idx ON instruments(isin, name) NULLS NOT DISTINCT;

CREATE TABLE IF NOT EXISTS fund_holdings (
    fund INT REFERENCES funds(id) ON DELETE CASCADE,
    instrument INT REFERENCES instruments(id) ON DELETE CASCADE,
    weight_percent NUMERIC(9,3) NOT NULL,
    PRIMARY KEY (fund, instrument)
);

CREATE INDEX fund_holdings_fund_idx ON fund_holdings(fund);
