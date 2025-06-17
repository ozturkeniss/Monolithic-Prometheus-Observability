package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"go-observability-demo/internal/metrics"
	"go-observability-demo/internal/model"
	"go-observability-demo/internal/repository"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	Repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.HttpRequestDuration.WithLabelValues("POST", "/products").Observe(duration)
		metrics.HttpRequestsTotal.WithLabelValues("POST", "/products", strconv.Itoa(http.StatusCreated)).Inc()
		metrics.ProductOperationsTotal.WithLabelValues("create").Inc()
	}()

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("POST", "/products", strconv.Itoa(http.StatusBadRequest)).Inc()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Repo.Create(&product); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("POST", "/products", strconv.Itoa(http.StatusInternalServerError)).Inc()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.HttpRequestDuration.WithLabelValues("GET", "/products").Observe(duration)
		metrics.HttpRequestsTotal.WithLabelValues("GET", "/products", strconv.Itoa(http.StatusOK)).Inc()
		metrics.ProductOperationsTotal.WithLabelValues("list").Inc()
	}()

	products, err := h.Repo.GetAll()
	if err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("GET", "/products", strconv.Itoa(http.StatusInternalServerError)).Inc()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.HttpRequestDuration.WithLabelValues("GET", "/products/{id}").Observe(duration)
	}()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("GET", "/products/{id}", strconv.Itoa(http.StatusBadRequest)).Inc()
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	product, err := h.Repo.GetByID(uint(id))
	if err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("GET", "/products/{id}", strconv.Itoa(http.StatusNotFound)).Inc()
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	metrics.HttpRequestsTotal.WithLabelValues("GET", "/products/{id}", strconv.Itoa(http.StatusOK)).Inc()
	metrics.ProductOperationsTotal.WithLabelValues("get").Inc()
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.HttpRequestDuration.WithLabelValues("PUT", "/products/{id}").Observe(duration)
	}()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("PUT", "/products/{id}", strconv.Itoa(http.StatusBadRequest)).Inc()
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("PUT", "/products/{id}", strconv.Itoa(http.StatusBadRequest)).Inc()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product.ID = uint(id)
	if err := h.Repo.Update(&product); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("PUT", "/products/{id}", strconv.Itoa(http.StatusInternalServerError)).Inc()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	metrics.HttpRequestsTotal.WithLabelValues("PUT", "/products/{id}", strconv.Itoa(http.StatusOK)).Inc()
	metrics.ProductOperationsTotal.WithLabelValues("update").Inc()
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metrics.HttpRequestDuration.WithLabelValues("DELETE", "/products/{id}").Observe(duration)
	}()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("DELETE", "/products/{id}", strconv.Itoa(http.StatusBadRequest)).Inc()
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.Repo.Delete(uint(id)); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("DELETE", "/products/{id}", strconv.Itoa(http.StatusInternalServerError)).Inc()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	metrics.HttpRequestsTotal.WithLabelValues("DELETE", "/products/{id}", strconv.Itoa(http.StatusNoContent)).Inc()
	metrics.ProductOperationsTotal.WithLabelValues("delete").Inc()
	w.WriteHeader(http.StatusNoContent)
}
