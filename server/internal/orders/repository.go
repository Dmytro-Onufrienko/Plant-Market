package orders

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Create(
	ctx context.Context,
	pool *pgxpool.Pool,
	userID string,
	items []ItemInput,
	address ShippingAddress,
) (*Order, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	total := 0
	type itemData struct {
		plantID int
		qty     int
		price   int
	}
	resolved := make([]itemData, 0, len(items))

	for _, item := range items {
		var price, stock int
		err := tx.QueryRow(ctx,
			`SELECT price, stock FROM plants WHERE id = $1 FOR UPDATE`, item.PlantID).
			Scan(&price, &stock)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("plant %d not found", item.PlantID)
		}
		if err != nil {
			return nil, err
		}
		if stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for plant %d", item.PlantID)
		}
		total += price * item.Quantity
		resolved = append(resolved, itemData{item.PlantID, item.Quantity, price})
	}

	addrJSON, err := json.Marshal(address)
	if err != nil {
		return nil, err
	}

	var order Order
	var returnedAddr []byte
	err = tx.QueryRow(ctx, `
		INSERT INTO orders (user_id, total_price, shipping_address)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, status, total_price, shipping_address, created_at`,
		userID, total, addrJSON).
		Scan(&order.ID, &order.UserID, &order.Status, &order.TotalPrice,
			&returnedAddr, &order.CreatedAt)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(returnedAddr, &order.ShippingAddress); err != nil {
		return nil, err
	}

	for _, item := range resolved {
		_, err = tx.Exec(ctx, `
			INSERT INTO order_items (order_id, plant_id, quantity, price_at_purchase)
			VALUES ($1, $2, $3, $4)`,
			order.ID, item.plantID, item.qty, item.price)
		if err != nil {
			return nil, err
		}

		_, err = tx.Exec(ctx,
			`UPDATE plants SET stock = stock - $1 WHERE id = $2`, item.qty, item.plantID)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return &order, nil
}

func GetByUser(ctx context.Context, pool *pgxpool.Pool, userID string) ([]Order, error) {
	rows, err := pool.Query(ctx, `
		SELECT id, user_id, status, total_price, shipping_address, created_at
		FROM orders WHERE user_id = $1
		ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanOrders(rows)
}

func GetByID(ctx context.Context, pool *pgxpool.Pool, id int, userID string) (*OrderWithItemsDetailed, error) {
	order := &OrderWithItemsDetailed{}
	var addrRaw []byte
	err := pool.QueryRow(ctx, `
		SELECT id, user_id, status, total_price, shipping_address, created_at
		FROM orders WHERE id = $1 AND user_id = $2`, id, userID).
		Scan(&order.ID, &order.UserID, &order.Status, &order.TotalPrice,
			&addrRaw, &order.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(addrRaw, &order.ShippingAddress); err != nil {
		return nil, err
	}

	rows, err := pool.Query(ctx, `
		SELECT oi.id, oi.order_id, oi.plant_id, oi.quantity, oi.price_at_purchase,
		       p.name, p.slug, p.images
		FROM order_items oi
		JOIN plants p ON oi.plant_id = p.id
		WHERE oi.order_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item OrderItemWithPlant
		var images []string
		if err := rows.Scan(&item.ID, &item.OrderID, &item.PlantID,
			&item.Quantity, &item.PriceAtPurchase,
			&item.PlantName, &item.PlantSlug, &images); err != nil {
			return nil, err
		}
		if images == nil {
			images = []string{}
		}
		item.PlantImages = images
		order.Items = append(order.Items, item)
	}
	if order.Items == nil {
		order.Items = []OrderItemWithPlant{}
	}
	return order, rows.Err()
}

func GetAll(ctx context.Context, pool *pgxpool.Pool) ([]OrderWithUser, error) {
	rows, err := pool.Query(ctx, `
		SELECT o.id, o.user_id, o.status, o.total_price, o.shipping_address, o.created_at,
		       u.email, u.name
		FROM orders o
		JOIN users u ON o.user_id = u.id
		ORDER BY o.created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanOrdersWithUser(rows)
}

func GetRecent(ctx context.Context, pool *pgxpool.Pool, limit int) ([]OrderWithUser, error) {
	rows, err := pool.Query(ctx, `
		SELECT o.id, o.user_id, o.status, o.total_price, o.shipping_address, o.created_at,
		       u.email, u.name
		FROM orders o
		JOIN users u ON o.user_id = u.id
		ORDER BY o.created_at DESC
		LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanOrdersWithUser(rows)
}

func UpdateStatus(ctx context.Context, pool *pgxpool.Pool, id int, status string) error {
	tag, err := pool.Exec(ctx,
		`UPDATE orders SET status = $1 WHERE id = $2`, status, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errors.New("not found")
	}
	return nil
}

func scanOrders(rows pgx.Rows) ([]Order, error) {
	var result []Order
	for rows.Next() {
		var o Order
		var addrRaw []byte
		if err := rows.Scan(&o.ID, &o.UserID, &o.Status, &o.TotalPrice,
			&addrRaw, &o.CreatedAt); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(addrRaw, &o.ShippingAddress); err != nil {
			return nil, err
		}
		result = append(result, o)
	}
	return result, rows.Err()
}

func scanOrdersWithUser(rows pgx.Rows) ([]OrderWithUser, error) {
	var result []OrderWithUser
	for rows.Next() {
		var o OrderWithUser
		var addrRaw []byte
		if err := rows.Scan(&o.ID, &o.UserID, &o.Status, &o.TotalPrice,
			&addrRaw, &o.CreatedAt, &o.UserEmail, &o.UserName); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(addrRaw, &o.ShippingAddress); err != nil {
			return nil, err
		}
		result = append(result, o)
	}
	return result, rows.Err()
}
