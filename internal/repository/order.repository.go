package repository

import (
	"context"
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o *OrderRepository) CheckoutOrder(userID int, payload dto.CreatePaymentRequestDTO) (map[string]any, error) {
	tx, err := o.db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	var cartID int
	if err := tx.QueryRow(context.Background(), `SELECT id FROM carts WHERE user_id = $1 LIMIT 1`, userID).Scan(&cartID); err != nil {
		return nil, errors.New("Cart not found")
	}

	rows, err := tx.Query(context.Background(), `
		SELECT ci.product_id, ci.quantity, p.name, p.price,
			COALESCE(v.variant_name, '') AS variant_name, COALESCE(v.additional_price, 0) AS variant_price,
			COALESCE(s.portion_size, '') AS portion_size, COALESCE(s.additional_price, 0) AS portion_price,
			COALESCE(pi.image, '') AS product_image
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.id
		LEFT JOIN product_variants v ON ci.variant_id = v.id
		LEFT JOIN product_portions s ON ci.size_id = s.id
		LEFT JOIN LATERAL (SELECT image FROM product_images WHERE product_id = p.id LIMIT 1) pi ON true
		WHERE ci.cart_id = $1`, cartID)
	if err != nil {
		return nil, err
	}
	items, err := pgx.CollectRows(rows, pgx.RowToMap)
	if err != nil || len(items) == 0 {
		return nil, errors.New("Cart is empty")
	}

	now := time.Now()
	var orderID int
	var status int
	if err := tx.QueryRow(context.Background(), `
		INSERT INTO orders (cart_id, total, status, fullname, phone, email, address, delivery, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		RETURNING id, status`,
		cartID, 0, 1, payload.Fullname, payload.Phone, payload.Email, payload.Address, payload.Delivery, now, now).
		Scan(&orderID, &status); err != nil {
		return nil, err
	}

	total := 0
	orderItems := make([]map[string]any, 0, len(items))
	for _, item := range items {
		qty := int(item["quantity"].(int32))
		base := int(item["price"].(int32))
		variantPrice := int(item["variant_price"].(int32))
		portionPrice := int(item["portion_price"].(int32))
		unitFinal := base + variantPrice + portionPrice
		totalPrice := unitFinal * qty
		total += totalPrice

		itemMap := map[string]any{}
		var itemID, itemOrderID, itemProductID, itemQuantity, itemTotalPrice int
		var itemProductName string
		err := tx.QueryRow(context.Background(), `
			INSERT INTO order_items (
				order_id, product_id, product_name, product_image, variant_name, portion_size, quantity,
				unit_base_price, variant_price, portion_price, unit_final_price, total_price, created_at, updated_at
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
			RETURNING id, order_id, product_id, product_name, quantity, total_price`,
			orderID, item["product_id"], item["name"], item["product_image"], item["variant_name"], item["portion_size"], qty,
			base, variantPrice, portionPrice, unitFinal, totalPrice, now, now,
		).Scan(&itemID, &itemOrderID, &itemProductID, &itemProductName, &itemQuantity, &itemTotalPrice)
		if err != nil {
			return nil, err
		}
		itemMap["id"] = itemID
		itemMap["order_id"] = itemOrderID
		itemMap["product_id"] = itemProductID
		itemMap["product_name"] = itemProductName
		itemMap["quantity"] = itemQuantity
		itemMap["total_price"] = itemTotalPrice
		orderItems = append(orderItems, itemMap)
	}

	if _, err := tx.Exec(context.Background(), `UPDATE orders SET total = $1, updated_at = $2 WHERE id = $3`, total, now, orderID); err != nil {
		return nil, err
	}
	if _, err := tx.Exec(context.Background(), `DELETE FROM cart_items WHERE cart_id = $1`, cartID); err != nil {
		return nil, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return map[string]any{
		"id":         orderID,
		"cart_id":    cartID,
		"total":      total,
		"status":     status,
		"fullname":   payload.Fullname,
		"phone":      payload.Phone,
		"email":      payload.Email,
		"address":    payload.Address,
		"delivery":   payload.Delivery,
		"items":      orderItems,
		"created_at": now,
		"updated_at": now,
	}, nil
}
