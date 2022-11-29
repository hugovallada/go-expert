package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hugovallada/go-expert/api/internal/dto"
	"github.com/hugovallada/go-expert/api/internal/entity"
	"github.com/hugovallada/go-expert/api/internal/infra/database"
	entityPkg "github.com/hugovallada/go-expert/api/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductModel
}

func NewProductHandler(db database.ProductModel) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

// Create Product godoc
// @Sumary Create Product
// @Description Create Product with Authentication
// @Tags products
// @Accept json
// @Produce json
// @Param request body dto.CreateProductInput true "product request"
// @Success 201
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 500 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.CreateProductInput
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := entity.NewProduct(productDto.Name, productDto.Price)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(product)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get Product godoc
// @Sumary Get One Product
// @Description Get only one product with Authentication
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product ID" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Update Product godoc
// @Sumary Update Product
// @Description Update Product with Authentication
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product ID" Format(uuid)
// @Param request body dto.CreateProductInput true "product request"
// @Success 202
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

// Delete Product godoc
// @Sumary Delete Product
// @Description Delete Product with Authentication
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product ID" Format(uuid)
// @Success 202
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

// Get Products godoc
// @Sumary Get All Products
// @Description Get all products with Authentication
// @Tags products
// @Accept json
// @Produce json
// @Param page query string false "page number"
// @Param size query string false "size number"
// @Success 200 {array} entity.Product
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(limitStr)
	if err != nil {
		size = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(page, size, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}
