package services

import (
	"catalog-products/internal/database/repository"
	"catalog-products/internal/models"
	"context"
	"time"
)

// ProductService. Сервіс товарів
type ProductService struct {
	repo *repository.ProductRepo
}

// NewProductService. Конструктор сервіса товарів
func NewProductService(repo *repository.ProductRepo) ProductService {
	return ProductService{
		repo: repo,
	}
}

// List. Повертає список товарів
func (s *ProductService) List(params map[string]string) []models.ProductDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	products := s.repo.List(ctx, params)
	var productsDto []models.ProductDTO
	for _, p := range products {
		productsDto = append(productsDto, p.DTO())
	}
	return productsDto
}

// One. Повертає товар за ідентифікатором
func (s *ProductService) One(id int) models.ProductDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	p := s.repo.One(ctx, id)
	return p.DTO()
}

// OneUnscoped. Повертає м'яко видалений товар за ідентифікатором
func (s *ProductService) OneUnscoped(id int) models.ProductDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	p := s.repo.OneUnscoped(ctx, id)
	return p.DTO()
}

// Create. Створює новий товар
func (s *ProductService) Create(dto models.ProductDTO) models.ProductDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var model models.Product
	dto.FillModel(&model)
	s.repo.Save(ctx, &model)
	return model.DTO()
}

// Update. Оновлює товар
func (s *ProductService) Update(dto models.ProductDTO) models.ProductDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	model := s.repo.One(ctx, dto.ID)
	dto.FillModel(&model)
	s.repo.Save(ctx, &model)
	return model.DTO()
}

// SoftDelete. М'яке видалення товара
func (s *ProductService) SoftDelete(dto models.ProductDTO) models.ProductDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	model := s.repo.One(ctx, int(dto.ID))
	s.repo.SoftDelete(ctx, &model)
	return model.DTO()
}

// Recover. Відновлення м'яко видаленного товара
func (s *ProductService) Recover(dto models.ProductDTO) models.ProductDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	model := s.repo.OneUnscoped(ctx, int(dto.ID))
	s.repo.Recover(ctx, &model)
	return model.DTO()
}

// Delete. Остаточне видалення товара
func (s *ProductService) Delete(dto models.ProductDTO) models.ProductDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	model := s.repo.One(ctx, dto.ID)
	dto.FillModel(&model)
	s.repo.Delete(ctx, &model)
	return model.DTO()
}

// Exists. Перевірка існування товара за ідентифікатором
func (s *ProductService) Exists(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return s.repo.Exists(ctx, id)
}

// Exists. Перевірка існування м'яко видаленного товара за ідентифікатором
func (s *ProductService) ExistsUnscoped(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return s.repo.ExistsUnscoped(ctx, id)
}
