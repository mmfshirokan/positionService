package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mmfshirokan/positionService/internal/model"
	"github.com/shopspring/decimal"
)

type postgres struct {
	dbpool *pgxpool.Pool
}

type DBInterface interface {
	Add(ctx context.Context, position model.Position) error
	Deleete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, id uuid.UUID) ([]model.Position, error)
	Update(ctx context.Context, position model.Position) error
	GetLaterThen(ctx context.Context, t time.Time) ([]model.Position, error)
	GetOneState(ctx context.Context, operID uuid.UUID) (open bool, err error)
}

func NewPostgresRepository(conn *pgxpool.Pool) DBInterface {
	return &postgres{
		dbpool: conn,
	}
}

// NOTE: Add without close price
func (p *postgres) Add(ctx context.Context, position model.Position) error {
	_, err := p.dbpool.Exec(ctx, "INSERT INTO trading.positions (operation_id, user_id, symbol, created_at, open_price, buy, open) VALUES ($1, $2, $3, $4, $5, $6)",
		position.OperationID,
		position.UserID,
		position.Symbol,
		position.OpenPrice,
		position.CreatedAt,
		position.Buy,
		position.Open,
	)
	return err
}

func (p *postgres) Deleete(ctx context.Context, operID uuid.UUID) error {
	_, err := p.dbpool.Exec(ctx, "DELETE FROM trading.positions WHERE operation_id = $1", operID)
	return err
}

// TODO: Add created_at?
func (p *postgres) Get(ctx context.Context, userID uuid.UUID) (res []model.Position, err error) {
	rows, err := p.dbpool.Query(ctx, "SELECT (operation_id, user_id, symbol, open_price, close_price, buy, open) FROM trading.positions WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tmpPos := model.Position{}

		if err := rows.Scan(
			&tmpPos,
		); err != nil {
			return nil, err
		}

		res = append(res, tmpPos)
	}

	return res, nil
}

func (p *postgres) Update(ctx context.Context, pos model.Position) error {
	_, err := p.dbpool.Exec(ctx, "UPDATE trading.positions SET close_price = $1, open = $2 WHERE operation_id = $3", pos.ClosePrice, pos.Open, pos.OperationID)
	return err
}

func (p *postgres) GetLaterThen(ctx context.Context, t time.Time) (res []model.Position, err error) {
	rows, err := p.dbpool.Query(ctx, "SELECT (operation_id, symbol, open_price, buy) FROM trading.positions WHERE open_time > $1", t)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tmpPos := struct {
			OperationID uuid.UUID
			Symbol      string
			OpenPrice   decimal.Decimal
			Buy         bool
		}{}

		if err := rows.Scan(
			&tmpPos,
		); err != nil {
			return nil, err
		}

		res = append(res, model.Position{
			OperationID: tmpPos.OperationID,
			Symbol:      tmpPos.Symbol,
			OpenPrice:   tmpPos.OpenPrice,
			Buy:         tmpPos.Buy,
		})
	}

	return res, nil
}

func (p *postgres) GetOneState(ctx context.Context, operID uuid.UUID) (open bool, err error) {
	err = p.dbpool.QueryRow(ctx, "SELECT (open) FROM trading.positions WHERE operation_id = $1", operID).Scan(&open)
	return
}
