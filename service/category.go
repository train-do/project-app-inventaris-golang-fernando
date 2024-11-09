package service

import (
	"github.com/train-do/project-app-inventaris-golang-fernando/collection"
	"github.com/train-do/project-app-inventaris-golang-fernando/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) GetAllCategory() (collection.Response, error) {
	// fmt.Println("MASUK Service CATEGORY")
	data, err := s.repo.FindAll()
	if err != nil {
		return collection.Response{}, err
	}
	return collection.Response{Data: data}, nil
}

func (s *CategoryService) GetCategoryById(id int) (collection.Response, error) {
	data, err := s.repo.FindById(id)
	if err != nil {
		return collection.Response{}, err
	}
	return collection.Response{Data: data}, nil
}

func (s *CategoryService) CreateCategory(form collection.FormCategory) (collection.Response, error) {
	data, err := s.repo.Insert(form)
	if err != nil {
		return collection.Response{}, err
	}
	data.Name = form.Name
	data.Description = form.Description
	return collection.Response{Data: data}, nil
}

func (s *CategoryService) UpdateCategory(id int, form collection.FormCategory) (collection.Response, error) {
	if _, err := s.repo.FindById(id); err != nil {
		return collection.Response{}, err
	}
	if _, err := s.repo.Update(id, form); err != nil {
		return collection.Response{}, err
	}
	data, _ := s.repo.FindById(id)
	data.Name = form.Name
	data.Description = form.Description
	return collection.Response{Data: data}, nil
}
func (s *CategoryService) DeleteCategory(id int) error {
	if _, err := s.repo.FindById(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
