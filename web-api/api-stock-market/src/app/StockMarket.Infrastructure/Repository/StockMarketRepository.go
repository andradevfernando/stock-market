package Repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"api-stock-market/src/app/StockMarket.Domain/Models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockRepository struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewStockRepository(ctx context.Context, pool *pgxpool.Pool) *StockRepository {
	return &StockRepository{
		pool: pool,
		ctx:  ctx,
	}
}

func (r *StockRepository) Create(modelList []*Models.StockMarketModel) error {
	const insertStockQuery = `
        INSERT INTO stock_market (
            ticker, target_from, target_to, company, 
            action, brokerage, rating_from, rating_to, "time"
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9
        )`

	for _, model := range modelList {
		_, err := r.pool.Exec(r.ctx, insertStockQuery,
			model.Ticker,
			model.TargetFrom,
			model.TargetTo,
			model.Company,
			model.Action,
			model.Brokerage,
			model.RatingFrom,
			model.RatingTo,
			model.Time,
		)
		if err != nil {
			return fmt.Errorf("error inserting stock %s: %w", model.Ticker, err)
		}
	}
	return nil
}
func (r *StockRepository) GetList(
	page int, limit *int,
	startDate, endDate *time.Time,
	companyName string,
) ([]*Models.StockMarketModel, error) {

	var args []interface{}
	queryBuilder := strings.Builder{}

	// Query base
	queryBuilder.WriteString(`
        SELECT 
            ticker, target_from, target_to, company, 
            action, brokerage, rating_from, rating_to, time
        FROM stock_market
    `)

	// Filtros
	conditions := []string{}

	// Filtro por data
	if startDate != nil && endDate != nil {
		conditions = append(conditions, "time BETWEEN $1 AND $2")
		args = append(args, *startDate, *endDate)
	}

	// Filtro por nome da empresa
	if companyName != "" {
		conditions = append(conditions, fmt.Sprintf("company ILIKE $%d", len(args)+1))
		args = append(args, "%"+companyName+"%")
	}

	// Combinar condições
	if len(conditions) > 0 {
		queryBuilder.WriteString("WHERE " + strings.Join(conditions, " AND "))
	}

	// Ordenação
	queryBuilder.WriteString("\n ORDER BY time DESC")

	if limit != nil {
		// Posição dos parâmetros de LIMIT e OFFSET
		limitOffsetPos := len(args) + 1

		queryBuilder.WriteString(fmt.Sprintf("\n LIMIT $%d OFFSET $%d", limitOffsetPos, limitOffsetPos+1))

		// Calcula o offset e adiciona os valores dos parâmetros
		offset := *limit * (page - 1)
		args = append(args, *limit, offset)
	}

	// Monta a query final
	finalQuery := queryBuilder.String()

	rows, err := r.pool.Query(r.ctx, finalQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar consulta paginada: %w", err)
	}
	defer rows.Close()

	var stocks []*Models.StockMarketModel
	for rows.Next() {
		var stock Models.StockMarketModel
		if err := rows.Scan(
			&stock.Ticker,
			&stock.TargetFrom,
			&stock.TargetTo,
			&stock.Company,
			&stock.Action,
			&stock.Brokerage,
			&stock.RatingFrom,
			&stock.RatingTo,
			&stock.Time,
		); err != nil {
			return nil, fmt.Errorf("erro ao escanear a linha: %w", err)
		}
		stocks = append(stocks, &stock)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das linhas: %w", err)
	}
	return stocks, nil
}
