package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/train-do/project-app-inventaris-golang-fernando/collection"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (r *CategoryRepository) FindAll() ([]collection.Category, error) {
	query := `select * from "Category";`
	rows, err := r.db.Query(query)
	var categories []collection.Category
	for rows.Next() {
		var category collection.Category
		rows.Scan(&category.Id, &category.Name, &category.Description)
		if err != nil {
			fmt.Println(err, "REPO FINDALL")
			return []collection.Category{}, errors.New("kategori tidak ditemukan")
		}
		categories = append(categories, category)
	}
	return categories, nil
}
func (r *CategoryRepository) FindById(id int) (collection.Category, error) {
	query := `select * from "Category" where id = $1;`
	var category collection.Category
	err := r.db.QueryRow(query, id).Scan(&category.Id, &category.Name, &category.Description)
	if err != nil {
		fmt.Println(err, "REPO FINDBYID")
		return collection.Category{}, errors.New("kategori tidak ditemukan")
	}
	return category, nil
}
func (r *CategoryRepository) Insert(form collection.FormCategory) (collection.Category, error) {
	query := `insert into "Category" ("name", "description") values ($1, $2) returning id;`
	var category collection.Category
	err := r.db.QueryRow(query, form.Name, form.Description).Scan(&category.Id)
	if err != nil {
		// fmt.Println(err.Error(), "REPO INSERT")
		return collection.Category{}, errors.New("bad request")
	}
	return category, nil
}
func (r *CategoryRepository) Update(id int, form collection.FormCategory) (collection.Category, error) {
	query := `update "Category" set "name"=$1, "description"=$2 where "id"=$3;`
	var category collection.Category
	_, err := r.db.Exec(query, form.Name, form.Description, id)
	if err != nil {
		// fmt.Println(err.Error(), "REPO UPDATE")
		return collection.Category{}, errors.New("bad request")
	}
	return category, nil
}
func (r *CategoryRepository) Delete(id int) error {
	query := `delete from "Category" where "id"=$1;`
	_, err := r.db.Exec(query, id)
	if err != nil {
		// fmt.Println(err.Error(), "REPO INSERT")
		return errors.New("category tidak ditemukan")
	}
	return nil
}
