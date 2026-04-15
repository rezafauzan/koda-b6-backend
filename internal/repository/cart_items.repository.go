package repository

import (
	"context"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CartItemRepository struct {
	db *pgxpool.Pool
}

func NewCartItemRepository(db *pgxpool.Pool) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (c *CartItemRepository) AddItem(item dto.CreateCartItemRequestDTO) (models.CartItem, error) {
	sql := `INSERT INTO cart_items (cart_id, product_id, size, hotice, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id, cart_id, product_id, size, hotice, quantity, created_at, updated_at`

	rows, err := c.db.Query(context.Background(), sql, item.CartId, item.ProductId, item.Size, item.Hotice, item.Quantity)
	if err != nil {
		return models.CartItem{}, err
	}

	cartItems, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.CartItem])
	if err != nil {
		return models.CartItem{}, err
	}

	return cartItems, nil
}

func (c *CartItemRepository) GetCartByUserId(userId int) (models.Cart, error) {
	sql := `SELECT id, user_id, created_at, updated_at FROM carts WHERE user_id = $1`

	rows, err := c.db.Query(context.Background(), sql, userId)
	if err != nil {
		return models.Cart{}, err
	}

	cart, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Cart])
	if err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}

func (c *CartItemRepository) GetCartItemsByCartId(cartId int) ([]models.CartItem, error) {
	sql := `SELECT id, cart_id, product_id, size, hotice, quantity, created_at, updated_at FROM cart_items WHERE cart_id = $1 ORDER BY created_at ASC`

	rows, err := c.db.Query(context.Background(), sql, cartId)
	if err != nil {
		return []models.CartItem{}, err
	}
	defer rows.Close()

	cartItems, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.CartItem])
	if err != nil {
		return []models.CartItem{}, err
	}

	return cartItems, nil
}

func (c *CartItemRepository) DeleteItem(id int) (models.CartItem, error) {
	sql := `DELETE FROM cart_items WHERE id = $1 RETURNING id, cart_id, product_id, size, hotice, quantity, created_at, updated_at`

	rows, err := c.db.Query(context.Background(), sql, id)
	if err != nil {
		return models.CartItem{}, err
	}

	deletedItems, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.CartItem])
	if err != nil {
		return models.CartItem{}, err
	}

	return deletedItems, nil
}

func (c *CartItemRepository) ClearCartItem(cartId int) ([]models.CartItem, error) {
	sql := `DELETE FROM cart_items WHERE cart_id = $1 RETURNING id, cart_id, product_id, size, hotice, quantity, created_at, updated_at`

	rows, err := c.db.Query(context.Background(), sql, cartId)
	if err != nil {
		return []models.CartItem{}, err
	}

	deletedItems, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.CartItem])
	if err != nil {
		return []models.CartItem{}, err
	}

	return deletedItems, nil
}
