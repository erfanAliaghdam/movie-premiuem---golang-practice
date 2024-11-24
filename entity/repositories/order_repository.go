package repositories

import (
	"database/sql"
	"movie_premiuem/entity"
)

type OrderRepository interface {
	CreateOrder(order entity.Order) (entity.Order, error)
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order entity.Order) (entity.Order, error) {
	query := `INSERT INTO orders (user_id, paid, paid_price) VALUES (?, ?, ?)`

	result, err := r.db.Exec(query, order.UserID, order.Paid, order.PaidPrice)
	if err != nil {
		return entity.Order{}, err
	}

	id, getIDErr := result.LastInsertId()
	if getIDErr != nil {
		return entity.Order{}, getIDErr
	}

	order.ID = id
	return order, nil
}
