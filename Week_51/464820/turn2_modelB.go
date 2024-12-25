package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model
type Product struct {
	ID          int
	Name        string
	Price       float64
	ArchivedAt time.Time // Added field for archiving
}

var products []Product

// View
func renderProduct(w http.ResponseWriter, p Product) {
	fmt.Fprintf(w, "Product ID: %d, Name: %s, Price: %.2f, Archived At: %s\n", p.ID, p.Name, p.Price, p.ArchivedAt.Format(time.RFC3339))
}

// Controller - ProductRepository
type ProductRepository interface {
	GetProduct(id int) (*Product, error)
	ArchiveProduct(id int) error
}

type InMemoryProductRepository struct {
	products map[int]*Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[int]*Product),
	}
}

func (r *InMemoryProductRepository) GetProduct(id int) (*Product, error) {
	p, ok := r.products[id]
	if !ok {
		return nil, fmt.Errorf("product not found")
	}
	return p, nil
}

func (r *InMemoryProductRepository) ArchiveProduct(id int) error {
	p, ok := r.products[id]
	if !ok {
		return fmt.Errorf("product not found")
	}
	// Mark the product as archived by setting the ArchivedAt field
	p.ArchivedAt = time.Now()
	return nil
}

// Controller - ProductController
type ProductController struct {
	repo ProductRepository
}

func NewProductController(repo ProductRepository) *ProductController {
	return &ProductController{repo: repo}
}

func (c *ProductController) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Validate and find the product
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := c.repo.GetProduct(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Render the product
	renderProduct(w, *product)
}

func (c *ProductController) ArchiveProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Validate and find the product
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = c.repo.ArchiveProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Product archived successfully")
}

func main() {
	// Initialize products and repository
	repo := NewInMemoryProductRepository()
	repo.products[1] = &Product{ID: 1, Name: "Laptop", Price: 1000.0}
	repo.products[2] = &Product{ID: 2, Name: "Smartphone", Price: 500.0}