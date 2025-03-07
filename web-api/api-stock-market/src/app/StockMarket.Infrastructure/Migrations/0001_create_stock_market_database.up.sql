CREATE TABLE IF NOT EXISTS stock_market (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    ticker VARCHAR(10) NOT NULL,
    target_from STRING NOT NULL,
    target_to STRING NOT NULL,
    company STRING NOT NULL,
    action STRING NOT NULL,
    brokerage STRING NOT NULL,
    rating_from STRING NOT NULL,
    rating_to STRING NOT NULL,
    time TIMESTAMPTZ NOT NULL,

    INDEX (ticker)
    );

-- Para buscas por texto completo
-- CREATE INVERTED INDEX ON stock_market (company, brokerage);