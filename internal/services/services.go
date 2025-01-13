package services

import (
	"catalog-products/internal/database/repository"
	"catalog-products/internal/logger"
	"catalog-products/internal/models"
	"context"
	"errors"
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
func (s *ProductService) List(params map[string]string) ([]models.ProductDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	products, err := s.repo.List(ctx, params)
	if err != nil {
		logger.Log().Warn(err)
		return nil, errors.New("service not available")
	}
	var productsDto []models.ProductDTO
	for _, p := range products {
		productsDto = append(productsDto, p.DTO())
	}
	return productsDto, nil
}

// One. Повертає товар за ідентифікатором
func (s *ProductService) One(id int) (models.ProductDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	p, err := s.repo.One(ctx, id)
	if err != nil {
		logger.Log().Warn(err)
		return models.ProductDTO{}, errors.New("service not available")
	}
	return p.DTO(), nil
}

// OneUnscoped. Повертає м'яко видалений товар за ідентифікатором
func (s *ProductService) OneUnscoped(id int) (models.ProductDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	p, err := s.repo.OneUnscoped(ctx, id)
	if err != nil {
		logger.Log().Warn(err)
		return models.ProductDTO{}, errors.New("service not available")
	}
	return p.DTO(), nil
}

// Create. Створює новий товар
func (s *ProductService) Create(dto models.ProductDTO) (models.ProductDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var model models.Product
	dto.FillModel(&model)
	err := s.repo.Save(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return models.ProductDTO{}, errors.New("service not available")
	}
	return model.DTO(), nil
}

// Update. Оновлює товар
func (s *ProductService) Update(dto models.ProductDTO) (models.ProductDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	model, err := s.repo.One(ctx, dto.ID)
	if err != nil {
		logger.Log().Warn(err)
		return models.ProductDTO{}, errors.New("service not available")
	}
	dto.FillModel(&model)
	err = s.repo.Save(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return models.ProductDTO{}, errors.New("service not available")
	}
	return model.DTO(), nil
}

// SoftDelete. М'яке видалення товара
func (s *ProductService) SoftDelete(dto models.ProductDTO) (models.ProductDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	model, err := s.repo.One(ctx, int(dto.ID))
	if err != nil {
		logger.Log().Warn(err)
		return models.ProductDTO{}, errors.New("service not available")
	}
	err = s.repo.SoftDelete(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return models.ProductDTO{}, errors.New("service not available")
	}
	return model.DTO(), nil
}

// Recover. Відновлення м'яко видаленного товара
func (s *ProductService) Recover(dto models.ProductDTO) (models.ProductDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	model, err := s.repo.OneUnscoped(ctx, int(dto.ID))
	if err != nil {
		logger.Log().Warn(err)
		return models.ProductDTO{}, errors.New("service not available")
	}
	err = s.repo.Recover(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return models.ProductDTO{}, errors.New("service not available")
	}
	return model.DTO(), nil
}

// Delete. Остаточне видалення товара
func (s *ProductService) Delete(dto models.ProductDTO) (models.ProductDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	model, err := s.repo.One(ctx, dto.ID)
	if err != nil {
		logger.Log().Warn(err)
		return models.ProductDTO{}, errors.New("service not available")
	}
	dto.FillModel(&model)
	err = s.repo.Delete(ctx, &model)
	if err != nil {
		logger.Log().Warn(err)
		return models.ProductDTO{}, errors.New("service not available")
	}
	return model.DTO(), nil
}

// Exists. Перевірка існування товара за ідентифікатором
func (s *ProductService) Exists(id int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	exists, err := s.repo.Exists(ctx, id)
	if err != nil {
		logger.Log().Warn(err)
		return exists, errors.New("service not available")
	}
	return exists, nil
}

// Exists. Перевірка існування м'яко видаленного товара за ідентифікатором
func (s *ProductService) ExistsUnscoped(id int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	exists, err := s.repo.ExistsUnscoped(ctx, id)
	if err != nil {
		logger.Log().Warn(err)
		return exists, errors.New("service not available")
	}
	return exists, nil
}
