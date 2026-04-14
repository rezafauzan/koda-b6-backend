package services

import "rezafauzan/koda-b6-golang/internal/repository"

type ProductReviewService struct {
	repo *repository.ProductReviewRepository
}

func NewProductReviewService(repo *repository.ProductReviewRepository) *ProductReviewService {
	return &ProductReviewService{repo: repo}
}

func (s *ProductReviewService) GetLatestReviews() (any, error) {
	return s.repo.GetLatestReviews()
}

func (s *ProductReviewService) GetPopularProducts() (any, error) {
	return s.repo.GetPopularProducts()
}
