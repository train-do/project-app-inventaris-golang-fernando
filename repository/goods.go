package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/train-do/project-app-inventaris-golang-fernando/collection"
)

type GoodRepository struct {
	db *sql.DB
}
type GoodsResponse struct {
	Id                  int       `json:"id,omitempty"`
	Name                string    `json:"name,omitempty"`
	Category            string    `json:"category,omitempty"`
	PhotoUrl            string    `json:"photo_url,omitempty"`
	Price               int       `json:"price,omitempty"`
	PurchaseDate        time.Time `json:"purchase_date,omitempty"`
	TotalUsageDays      int       `json:"total_usage_days,omitempty"`
	ReplacementRequired bool      `json:"replacement_required,omitempty"`
}
type Search struct {
	Key   string
	Value int
}

func NewGoodsRepository(db *sql.DB) *GoodRepository {
	return &GoodRepository{db}
}

func (r *GoodRepository) FindAll(search []Search) ([]GoodsResponse, int, error) {
	query := `select g.id , g."name" , c."name" , g.photo_url , g.price , g.purchase_date , g.total_usage_days , total_items from "Goods" g join "Category" c on g.category_id = c.id join "total_items" ti on true `
	var arrGoods []GoodsResponse
	var rows *sql.Rows
	var err error
	page := search[2].Value
	if page != 0 {
		page = (page - 1) * 2
	}
	if search[0].Value != 0 && search[1].Value != 0 {
		query = `with "total_items" as (select count(*) as total_items from "Goods" g join "Category" c on g.category_id = c.id where c.id = $1 and g.total_usage_days >= $2) ` + query
		query += ` where c.id = $1 and g.total_usage_days >= $2 order by g.id limit 2 offset $3;`
		rows, err = r.db.Query(query, search[0].Value, search[1].Value, page)
	} else if search[0].Value != 0 {
		query = `with "total_items" as (select count(*) as total_items from "Goods" g join "Category" c on g.category_id = c.id where c.id = $1)` + query
		query += ` where c.id = $2 order by g.id limit 2 offset $3;`
		rows, err = r.db.Query(query, search[0].Value, search[0].Value, page)
	} else if search[1].Value != 0 {
		query = `with "total_items" as (select count(*) as total_items from "Goods" g join "Category" c on g.category_id = c.id where g.total_usage_days >= $1) ` + query
		query += ` where g.total_usage_days >= $1 order by g.id limit 2 offset $2;`
		rows, err = r.db.Query(query, search[1].Value, page)
	} else {
		query = `with "total_items" as (select count(*) as total_items from "Goods" g join "Category" c on g.category_id = c.id) ` + query
		query += ` order by g.id limit 2 offset $1;`
		rows, err = r.db.Query(query, page)
	}
	if err != nil {
		fmt.Println(err, "REPO FINDALL")
		return []GoodsResponse{}, 0, fmt.Errorf("internal server error")
	}
	var totalItems int
	for rows.Next() {
		var goods GoodsResponse
		err = rows.Scan(&goods.Id, &goods.Name, &goods.Category, &goods.PhotoUrl, &goods.Price, &goods.PurchaseDate, &goods.TotalUsageDays, &totalItems)
		if err != nil {
			// fmt.Println(err, "REPO FINDALL NEXT")
			return []GoodsResponse{}, 0, fmt.Errorf("internal server error")
		}
		arrGoods = append(arrGoods, goods)
	}
	return arrGoods, totalItems, nil
}
func (r *GoodRepository) FindById(id int) (GoodsResponse, error) {
	query := `select g.id , g."name" , c."name" , g.photo_url , g.price , g.purchase_date , g.total_usage_days from "Goods" g join "Category" c on g.category_id = c.id where g.id = $1;`
	var goods GoodsResponse
	err := r.db.QueryRow(query, id).Scan(&goods.Id, &goods.Name, &goods.Category, &goods.PhotoUrl, &goods.Price, &goods.PurchaseDate, &goods.TotalUsageDays)
	if err != nil {
		// fmt.Println(err, "REPO FINDBYID")
		return GoodsResponse{}, errors.New("barang tidak ditemukan")
	}
	return goods, nil
}
func (r *GoodRepository) Insert(form collection.FormGoods, totalUsageDays int) (GoodsResponse, error) {
	query := `insert into "Goods" ("category_id", "name", "photo_url", "price", "purchase_date", "total_usage_days") values ($1, $2, $3, $4, $5, $6) returning id;`
	var goods GoodsResponse
	err := r.db.QueryRow(query, form.CategoryId, form.Name, form.PhotoUrl, form.Price, form.PurchaseDate, totalUsageDays).Scan(&goods.Id)
	if err != nil {
		// fmt.Println(err.Error(), "REPO INSERT")
		return GoodsResponse{}, errors.New("kategori tidak ditemukan")
	}
	return goods, nil
}
func (r *GoodRepository) Update(id int, form collection.FormGoods, totalUsageDays int) error {
	query := `update "Goods" set "category_id"=$1, "name"=$2, "photo_url"=$3, "price"=$4, "purchase_date"=$5, "total_usage_days"=$6 where "id"=$7;`
	_, err := r.db.Exec(query, form.CategoryId, form.Name, form.PhotoUrl, form.Price, form.PurchaseDate, totalUsageDays, id)
	if err != nil {
		// fmt.Println(err.Error(), "REPO UPDATE")
		if strings.Contains(err.Error(), "foreign key") {
			return errors.New("kategori tidak ditemukan")
		}
		return errors.New("barang tidak ditemukan")
	}
	return nil
}
func (r *GoodRepository) Delete(id int) error {
	query := `delete from "Goods" where id = $1;`
	_, err := r.db.Exec(query, id)
	if err != nil {
		// fmt.Println(err.Error(), "REPO DELETE")
		return errors.New("barang tidak ditemukan")
	}
	return nil
}
func (r *GoodRepository) FindAllGoods() ([]GoodsResponse, error) {
	fmt.Println("MASUK REPO FINDALL")
	query := `select g.id , g."name" , c."name" , g.photo_url , g.price , g.purchase_date , g.total_usage_days from "Goods" g join "Category" c on g.category_id = c.id;`
	var arrGoods []GoodsResponse
	rows, _ := r.db.Query(query)
	for rows.Next() {
		var goods GoodsResponse
		err := rows.Scan(&goods.Id, &goods.Name, &goods.Category, &goods.PhotoUrl, &goods.Price, &goods.PurchaseDate, &goods.TotalUsageDays)
		if err != nil {
			return []GoodsResponse{}, fmt.Errorf("internal server error")
		}
		arrGoods = append(arrGoods, goods)
	}
	return arrGoods, nil
	// return []GoodsResponse{}, nil
}
