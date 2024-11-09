package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/train-do/project-app-inventaris-golang-fernando/collection"
	"github.com/train-do/project-app-inventaris-golang-fernando/service"
	"github.com/train-do/project-app-inventaris-golang-fernando/utils"
	"github.com/train-do/project-app-inventaris-golang-fernando/validator"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("MASUK HANDLE CATEGORY")
	categories, err := h.service.GetAllCategory()
	if err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusNotFound, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, true, categories, http.StatusOK, "")
	json.NewEncoder(w).Encode(response)
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := utils.ToInt(chi.URLParam(r, "id"))
	goods, err := h.service.GetCategoryById(id)
	if err != nil {
		response := utils.SetResponse(w, false, goods, http.StatusNotFound, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, true, goods, http.StatusOK, "")
	json.NewEncoder(w).Encode(response)
}
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	body := collection.FormCategory{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusInternalServerError, "Internal Server Error")
		json.NewEncoder(w).Encode(response)
		return
	}
	if err := validator.ValidatorFormCategory(body); err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	goods, err := h.service.CreateCategory(body)
	if err != nil {
		response := utils.SetResponse(w, false, goods, http.StatusNotFound, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, true, goods, http.StatusOK, "Kategori berhasil ditambahkan")
	json.NewEncoder(w).Encode(response)
}
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := utils.ToInt(chi.URLParam(r, "id"))
	body := collection.FormCategory{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusInternalServerError, "Internal Server Error")
		json.NewEncoder(w).Encode(response)
		return
	}
	if err := validator.ValidatorFormCategory(body); err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusBadRequest, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	goods, err := h.service.UpdateCategory(id, body)
	if err != nil {
		response := utils.SetResponse(w, false, goods, http.StatusNotFound, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, true, goods, http.StatusOK, "Kategori berhasil diperbarui")
	json.NewEncoder(w).Encode(response)
}
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := utils.ToInt(chi.URLParam(r, "id"))
	err := h.service.DeleteCategory(id)
	if err != nil {
		response := utils.SetResponse(w, false, collection.Response{}, http.StatusNotFound, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, true, collection.Response{}, http.StatusOK, "Kategori berhasil dihapus")
	json.NewEncoder(w).Encode(response)
}
