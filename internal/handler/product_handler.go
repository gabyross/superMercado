package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gabyross/superMercado/internal/domain"
	"github.com/gabyross/superMercado/internal/service"
	"github.com/go-chi/chi"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ph.service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toWriteBody(w, http.StatusOK, products)
}

func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := ph.service.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	toWriteBody(w, http.StatusOK, product)
}

func (ph *ProductHandler) SearchProductByPriceGreaterThan(w http.ResponseWriter, r *http.Request) {
	priceStr := r.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("precio %v", price)

	products := ph.service.SearchProductByPriceGreaterThan(price)

	toWriteBody(w, http.StatusOK, products)
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct domain.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ph.service.CreateProduct(newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var product domain.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = ph.service.UpdateProduct(id, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	toWriteBody(w, http.StatusOK, product)

}
func (ph *ProductHandler) PatchProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var product domain.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = ph.service.PatchProduct(id, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	toWriteBody(w, http.StatusOK, product)
}
func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	panic("implement")
}

func toWriteBody(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
