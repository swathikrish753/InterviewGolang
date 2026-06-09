// Run this at: https://go.dev/play/
// Simulates a RESTful Book API using net/http/httptest
// No real server needed — works perfectly in Go Playground.
//
// SOLID Principles Applied:
//   S — Single Responsibility  : Repository (data), Service (business logic), Handlers (HTTP)
//   O — Open/Closed            : Add DB-backed repo without changing service or handlers
//   L — Liskov Substitution    : InMemoryBookRepository satisfies BookRepository anywhere
//   I — Interface Segregation  : BookRepository has only focused, relevant methods
//   D — Dependency Inversion   : Service & handlers depend on the interface, not the struct

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

// ─────────────────────────────────────────────
// DOMAIN
// ─────────────────────────────────────────────

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// ─────────────────────────────────────────────
// REPOSITORY LAYER  (I) narrow interface  (O) open for new impls  (D) abstraction
// ─────────────────────────────────────────────

type BookRepository interface {
	FindAll() []Book
	FindByID(id int) (Book, error)
	Save(b Book) Book
	Update(b Book) (Book, error)
	Delete(id int) error
}

// InMemoryBookRepository — (L) can substitute BookRepository anywhere
type InMemoryBookRepository struct {
	books  map[int]Book
	nextID int
}

func NewInMemoryBookRepository() *InMemoryBookRepository {
	r := &InMemoryBookRepository{books: make(map[int]Book), nextID: 1}
	r.Save(Book{Title: "The Go Programming Language", Author: "Donovan & Kernighan"})
	r.Save(Book{Title: "Clean Code", Author: "Robert C. Martin"})
	return r
}

func (r *InMemoryBookRepository) FindAll() []Book {
	all := make([]Book, 0, len(r.books))
	for _, b := range r.books {
		all = append(all, b)
	}
	return all
}

func (r *InMemoryBookRepository) FindByID(id int) (Book, error) {
	b, ok := r.books[id]
	if !ok {
		return Book{}, errors.New("book not found")
	}
	return b, nil
}

func (r *InMemoryBookRepository) Save(b Book) Book {
	b.ID = r.nextID
	r.nextID++
	r.books[b.ID] = b
	return b
}

func (r *InMemoryBookRepository) Update(b Book) (Book, error) {
	if _, ok := r.books[b.ID]; !ok {
		return Book{}, errors.New("book not found")
	}
	r.books[b.ID] = b
	return b, nil
}

func (r *InMemoryBookRepository) Delete(id int) error {
	if _, ok := r.books[id]; !ok {
		return errors.New("book not found")
	}
	delete(r.books, id)
	return nil
}

// ─────────────────────────────────────────────
// SERVICE LAYER  (S) business logic only  (D) depends on interface
// ─────────────────────────────────────────────

type BookService struct {
	repo BookRepository // depends on abstraction, not concrete type
}

func NewBookService(repo BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAllBooks() []Book          { return s.repo.FindAll() }
func (s *BookService) GetBook(id int) (Book, error) { return s.repo.FindByID(id) }

func (s *BookService) CreateBook(b Book) (Book, error) {
	if strings.TrimSpace(b.Title) == "" {
		return Book{}, errors.New("title is required")
	}
	if strings.TrimSpace(b.Author) == "" {
		return Book{}, errors.New("author is required")
	}
	return s.repo.Save(b), nil
}

func (s *BookService) UpdateBook(b Book) (Book, error) {
	if strings.TrimSpace(b.Title) == "" {
		return Book{}, errors.New("title is required")
	}
	return s.repo.Update(b)
}

func (s *BookService) DeleteBook(id int) error { return s.repo.Delete(id) }

// ─────────────────────────────────────────────
// HANDLER LAYER  (S) HTTP concerns only  (D) depends on service interface
// ─────────────────────────────────────────────

type BookHandler struct {
	svc *BookService
}

func NewBookHandler(svc *BookService) *BookHandler {
	return &BookHandler{svc: svc}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// Router: /books → list/create   /books/{id} → get/update/delete
func (h *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	switch {
	case len(parts) == 1 && parts[0] == "books":
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, http.StatusOK, h.svc.GetAllBooks())
		case http.MethodPost:
			var b Book
			if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
				writeError(w, http.StatusBadRequest, "invalid JSON")
				return
			}
			created, err := h.svc.CreateBook(b)
			if err != nil {
				writeError(w, http.StatusBadRequest, err.Error())
				return
			}
			writeJSON(w, http.StatusCreated, created)
		default:
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		}

	case len(parts) == 2 && parts[0] == "books":
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid id")
			return
		}
		switch r.Method {
		case http.MethodGet:
			book, err := h.svc.GetBook(id)
			if err != nil {
				writeError(w, http.StatusNotFound, err.Error())
				return
			}
			writeJSON(w, http.StatusOK, book)
		case http.MethodPut:
			var b Book
			if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
				writeError(w, http.StatusBadRequest, "invalid JSON")
				return
			}
			b.ID = id
			updated, err := h.svc.UpdateBook(b)
			if err != nil {
				writeError(w, http.StatusNotFound, err.Error())
				return
			}
			writeJSON(w, http.StatusOK, updated)
		case http.MethodDelete:
			if err := h.svc.DeleteBook(id); err != nil {
				writeError(w, http.StatusNotFound, err.Error())
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		}

	default:
		writeError(w, http.StatusNotFound, "route not found")
	}
}

// ─────────────────────────────────────────────
// HELPERS — simulate HTTP calls via httptest
// ─────────────────────────────────────────────

func call(handler http.Handler, method, path, body string) (int, string) {
	var reqBody *strings.Reader
	if body != "" {
		reqBody = strings.NewReader(body)
	} else {
		reqBody = strings.NewReader("")
	}
	req := httptest.NewRequest(method, path, reqBody)
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec.Code, strings.TrimSpace(rec.Body.String())
}

func section(title string) {
	fmt.Printf("\n══════════════════════════════════════\n  %s\n══════════════════════════════════════\n", title)
}

func show(label, method, path string, status int, body string) {
	fmt.Printf("%-22s %s %-20s  →  %d  %s\n", label, method, path, status, body)
}

// ─────────────────────────────────────────────
// MAIN — wire up and run demo calls
// ─────────────────────────────────────────────

func main() {
	// Dependency injection: concrete repo → service → handler
	repo := NewInMemoryBookRepository() // implements BookRepository
	service := NewBookService(repo)     // depends on interface
	handler := NewBookHandler(service)  // depends on service

	section("GET /books — list all (2 seeded)")
	status, body := call(handler, http.MethodGet, "/books", "")
	show("List all books", "GET", "/books", status, body)

	section("POST /books — create a new book")
	status, body = call(handler, http.MethodPost, "/books",
		`{"title":"Clean Architecture","author":"Robert C. Martin"}`)
	show("Create book", "POST", "/books", status, body)

	section("GET /books/3 — fetch the new book")
	status, body = call(handler, http.MethodGet, "/books/3", "")
	show("Get book id=3", "GET", "/books/3", status, body)

	section("PUT /books/3 — update title")
	status, body = call(handler, http.MethodPut, "/books/3",
		`{"title":"Clean Architecture (2nd Ed)","author":"Robert C. Martin"}`)
	show("Update book id=3", "PUT", "/books/3", status, body)

	section("DELETE /books/1 — remove first book")
	status, body = call(handler, http.MethodDelete, "/books/1", "")
	show("Delete book id=1", "DELETE", "/books/1", status, body)

	section("GET /books — list after delete (2 books left)")
	status, body = call(handler, http.MethodGet, "/books", "")
	show("List after delete", "GET", "/books", status, body)

	section("POST /books — validation: missing title")
	status, body = call(handler, http.MethodPost, "/books",
		`{"author":"Someone"}`)
	show("Bad create (no title)", "POST", "/books", status, body)

	section("GET /books/99 — not found")
	status, body = call(handler, http.MethodGet, "/books/99", "")
	show("Get missing book", "GET", "/books/99", status, body)
}
