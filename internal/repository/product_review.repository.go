package repository

import (
	"context"
	"rezafauzan/koda-b6-golang/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductReviewRepository struct {
	db *pgxpool.Pool
}

func NewProductReviewRepository(db *pgxpool.Pool) *ProductReviewRepository {
	return &ProductReviewRepository{db: db}
}

func (r *ProductReviewRepository) GetLatestReviews() ([]models.ProductReview, error) {
	rows, err := r.db.Query(context.Background(), `SELECT id, product_id, user_id, rating, messages, created_at, updated_at FROM product_reviews ORDER BY created_at DESC LIMIT 4`)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductReview])
}

func (r *ProductReviewRepository) GetPopularProducts() ([]map[string]any, error) {
	rows, err := r.db.Query(context.Background(), `SELECT products.id, products.category_id, products.name, products.description, products.price, products.stock, products.created_at, products.updated_at, product_reviews.total_reviews FROM products JOIN (SELECT product_id, COUNT(product_id) AS total_reviews FROM product_reviews GROUP BY product_id ORDER BY total_reviews DESC LIMIT 4) AS product_reviews ON products.id = product_reviews.product_id ORDER BY product_reviews.total_reviews DESC`)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToMap)
}
