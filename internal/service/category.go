package service

import (
	"context"

	"github.com/curso_grpc/project_grpc/internal/database"
	"github.com/curso_grpc/project_grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CategoryCreateResquest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, blank *pb.Blank) (*pb.Categories, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var ListCategories pb.Categories

	for i := range categories {
		categoryResponse := &pb.Category{
			Id:          categories[i].ID,
			Name:        categories[i].Name,
			Description: categories[i].Description,
		}

		ListCategories.Categories = append(ListCategories.Categories, categoryResponse)
	}

	return &ListCategories, nil
}
