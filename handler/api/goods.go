package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/train-do/project-app-inventaris-golang-fernando/collection"
	"github.com/train-do/project-app-inventaris-golang-fernando/repository"
	"github.com/train-do/project-app-inventaris-golang-fernando/service"
	"github.com/train-do/project-app-inventaris-golang-fernando/utils"
	"github.com/train-do/project-app-inventaris-golang-fernando/validator"
)

type GoodsHandler struct {
	service *service.GoodsService
}

func NewGoodsHandler(service *service.GoodsService) *GoodsHandler {
	return &GoodsHandler{service}
}

func (h *GoodsHandler) GetAllGoods(w http.ResponseWriter, r *http.Request) {
	search := []repository.Search{}
	category := utils.ToInt(r.URL.Query().Get("category"))
	totalUsageDays := utils.ToInt(r.URL.Query().Get("total_usage_days"))
	page := utils.ToInt(r.URL.Query().Get("page"))
	search = append(search, repository.Search{Key: "category", Value: category})
	search = append(search, repository.Search{Key: "total_usage_days", Value: totalUsageDays})
	search = append(search, repository.Search{Key: "total_usage_days", Value: page})
	goods, err := h.service.GetAllGoods(search)
	if err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusNotFound, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	// fmt.Printf("%+v\n", goods)
	response := utils.SetResponse(w, true, goods, http.StatusOK, "")
	json.NewEncoder(w).Encode(response)
}
func (h *GoodsHandler) GetGoods(w http.ResponseWriter, r *http.Request) {
	id := utils.ToInt(chi.URLParam(r, "id"))
	goods, err := h.service.GetGoodsById(id)
	if err != nil {
		response := utils.SetResponse(w, false, goods, http.StatusNotFound, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, true, goods, http.StatusOK, "")
	json.NewEncoder(w).Encode(response)
}
func (h *GoodsHandler) CreateGoods(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	file, fileHandler, err := r.FormFile("photo_url")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	ext := filepath.Ext(fileHandler.Filename)
	fileHandler.Filename = strings.ToLower(strings.ReplaceAll(r.FormValue("name"), " ", "")) + r.FormValue("purchase_date") + ext
	defer file.Close()

	photoUrl := r.FormValue("photo_url")
	if photoUrl == "" {
		photoUrl = "http://localhost:8080/image/" + fileHandler.Filename
	}

	// fmt.Println(photoUrl, fileHandler.Filename)
	body := collection.FormGoods{
		Name:         r.FormValue("name"),
		CategoryId:   utils.ToInt(r.FormValue("category_id")),
		PhotoUrl:     photoUrl,
		Price:        utils.ToInt(r.FormValue("price")),
		PurchaseDate: utils.ToTimeFormat(r.FormValue("purchase_date")),
	}
	if err := validator.ValidatorFormGoods(body); err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	goods, err := h.service.CreateGoods(body)
	if err != nil {
		response := utils.SetResponse(w, false, goods, http.StatusNotFound, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	out, err := os.Create(filepath.Join("images", fileHandler.Filename))
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	response := utils.SetResponse(w, true, goods, http.StatusCreated, "Barang berhasil ditambahkan")
	json.NewEncoder(w).Encode(response)
}
func (h *GoodsHandler) UpdateGoods(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	file, fileHandler, err := r.FormFile("photo_url")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	ext := filepath.Ext(fileHandler.Filename)
	fileHandler.Filename = strings.ToLower(strings.ReplaceAll(r.FormValue("name"), " ", "")) + time.Now().Format("2006-01-02") + ext
	defer file.Close()

	photoUrl := r.FormValue("photo_url")
	if photoUrl == "" {
		photoUrl = "http://localhost:8080/image/" + fileHandler.Filename
	}

	fmt.Println(photoUrl, fileHandler.Filename)
	body := collection.FormGoods{
		Name:         r.FormValue("name"),
		CategoryId:   utils.ToInt(r.FormValue("category_id")),
		PhotoUrl:     photoUrl,
		Price:        utils.ToInt(r.FormValue("price")),
		PurchaseDate: utils.ToTimeFormat(r.FormValue("purchase_date")),
	}
	if err := validator.ValidatorFormGoods(body); err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	id := utils.ToInt(chi.URLParam(r, "id"))
	goods, err := h.service.UpdateGoods(id, body)
	if err != nil {
		response := utils.SetResponse(w, false, goods, http.StatusNotFound, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	out, err := os.Create(filepath.Join("images", fileHandler.Filename))
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	response := utils.SetResponse(w, true, goods, http.StatusOK, "Barang berhasil diperbarui")
	json.NewEncoder(w).Encode(response)
}
func (h *GoodsHandler) DeleteGoods(w http.ResponseWriter, r *http.Request) {
	id := utils.ToInt(chi.URLParam(r, "id"))
	err := h.service.DeleteGoods(id)
	if err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, true, collection.Response{}, http.StatusOK, "Barang berhasil dihapus")
	json.NewEncoder(w).Encode(response)
}
func (h *GoodsHandler) GetAllInvestment(w http.ResponseWriter, r *http.Request) {
	investments, err := h.service.GetInvestments()
	if err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, true, investments, http.StatusOK, "")
	json.NewEncoder(w).Encode(response)
}
func (h *GoodsHandler) GetInvestmentById(w http.ResponseWriter, r *http.Request) {
	id := utils.ToInt(chi.URLParam(r, "id"))
	investments, err := h.service.GetInvestmentById(id)
	if err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, true, investments, http.StatusOK, "")
	json.NewEncoder(w).Encode(response)
}
func (h *GoodsHandler) GetReplacementNeeded(w http.ResponseWriter, r *http.Request) {
	replacementReq, err := h.service.GetReplacementNeeded()
	if err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, true, replacementReq, http.StatusOK, "")
	json.NewEncoder(w).Encode(response)
}
