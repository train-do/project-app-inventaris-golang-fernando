package service

import (
	"fmt"
	"math"
	"time"

	"github.com/train-do/project-app-inventaris-golang-fernando/collection"
	"github.com/train-do/project-app-inventaris-golang-fernando/repository"
	"github.com/train-do/project-app-inventaris-golang-fernando/utils"
)

type GoodsService struct {
	repo *repository.GoodRepository
}

func NewGoodsService(repo *repository.GoodRepository) *GoodsService {
	return &GoodsService{repo}
}

func (s *GoodsService) GetAllGoods(search []repository.Search) (collection.Response, error) {
	// fmt.Printf("%+v------\n", search)
	// return []repository.GoodsResponse{}, nil
	data, totalItem, err := s.repo.FindAll(search)
	if err != nil {
		fmt.Println(err, "Service")
		return collection.Response{}, err
	}
	page := search[2].Value
	if page == 0 {
		page = 1
	}
	if len(data) == 0 {
		return collection.Response{}, fmt.Errorf("gagal mengambil data barang inventaris")
	}
	response := collection.Response{
		Page:       page,
		Limit:      2,
		TotalItem:  totalItem,
		TotalPages: int(math.Ceil(float64(totalItem) / float64(2))),
		Data:       data,
	}
	return response, nil
}
func (s *GoodsService) GetGoodsById(id int) (collection.Response, error) {
	data, err := s.repo.FindById(id)
	if err != nil {
		return collection.Response{}, err
	}
	response := collection.Response{Data: data}
	return response, nil
}
func (s *GoodsService) CreateGoods(form collection.FormGoods) (collection.Response, error) {
	totalUsageDays := utils.CalculateTotalUsageDays(form.PurchaseDate)
	feedback, err := s.repo.Insert(form, totalUsageDays)
	if err != nil {
		return collection.Response{}, err
	}
	data, err := s.repo.FindById(feedback.Id)
	if err != nil {
		return collection.Response{}, err
	}
	response := collection.Response{Data: data}
	return response, nil
}
func (s *GoodsService) UpdateGoods(id int, form collection.FormGoods) (collection.Response, error) {
	totalUsageDays := utils.CalculateTotalUsageDays(form.PurchaseDate)
	if _, err := s.repo.FindById(id); err != nil {
		return collection.Response{}, err
	}
	err := s.repo.Update(id, form, totalUsageDays)
	if err != nil {
		return collection.Response{}, err
	}
	data, err := s.repo.FindById(id)
	if err != nil {
		return collection.Response{}, err
	}
	response := collection.Response{Data: data}
	return response, nil
}
func (s *GoodsService) DeleteGoods(id int) error {
	_, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
func (s *GoodsService) GetInvestments() (collection.Response, error) {
	goods, err := s.repo.FindAllGoods()
	if err != nil {
		return collection.Response{}, err
	}
	var total_investment, depreciated_value int
	for _, v := range goods {
		total_investment += v.Price
		depreciated_value += utils.CalculateDepreciation(v.Price, v.TotalUsageDays, collection.RateDepreciation)
	}
	data := struct {
		TotalInvestment  int `json:"total_invesment"`
		DepreciatedValue int `json:"depreciated_value"`
	}{
		TotalInvestment:  total_investment,
		DepreciatedValue: depreciated_value,
	}
	return collection.Response{Data: data}, err
}
func (s *GoodsService) GetInvestmentById(id int) (collection.Response, error) {
	goods, err := s.repo.FindById(id)
	if err != nil {
		return collection.Response{}, err
	}
	data := struct {
		Id               int    `json:"item_id"`
		Name             string `json:"name"`
		InitialPrice     int    `json:"initial_price"`
		DepreciatedValue int    `json:"depreciated_value"`
		DepreciatedRate  int    `json:"depreciated_rate"`
	}{
		Id:               goods.Id,
		Name:             goods.Name,
		InitialPrice:     goods.Price,
		DepreciatedValue: utils.CalculateDepreciation(goods.Price, goods.TotalUsageDays, collection.RateDepreciation),
		DepreciatedRate:  collection.RateDepreciation,
	}
	response := collection.Response{Data: data}
	return response, nil
}
func (s *GoodsService) GetReplacementNeeded() (collection.Response, error) {
	fmt.Printf("MASUK SERVICE ReplacementRequired\n")
	goods, err := s.repo.FindAllGoods()
	if err != nil {
		return collection.Response{}, err
	}
	// fmt.Printf("%+v", goods)
	type IsReplace struct {
		Id                  int
		Name                string
		Category            string
		PurchaseDate        time.Time
		TotalUsageDays      int
		ReplacementRequired bool
	}
	var data []IsReplace
	for _, v := range goods {
		isReplace := IsReplace{}
		if v.TotalUsageDays > 100 {
			isReplace.Id = v.Id
			isReplace.Name = v.Name
			isReplace.Category = v.Category
			isReplace.PurchaseDate = v.PurchaseDate
			isReplace.TotalUsageDays = v.TotalUsageDays
			isReplace.ReplacementRequired = true
			data = append(data, isReplace)
		}
	}
	return collection.Response{Data: data}, err
	// return collection.Response{}, nil
}
