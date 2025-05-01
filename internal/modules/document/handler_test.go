package document

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	CreateDocumentFunc  func(ctx context.Context, document *Document) error
	UpdateDocumentFunc  func(ctx context.Context, document *Document) error
	DeleteDocumentFunc  func(ctx context.Context, id uint) error
	GetByIdFunc         func(ctx context.Context, id uint) (*Document, error)
	GetAllDocumentsFunc func(ctx context.Context, pagination db.Pagination) (*db.Pagination, error)
}


func (m *mockRepository) CreateDocument(ctx context.Context, document *Document) error {
	return m.CreateDocumentFunc(ctx, document)
}

func (m *mockRepository) UpdateDocument(ctx context.Context, document *Document) error {
	return m.UpdateDocumentFunc(ctx, document)
}

func (m *mockRepository) DeleteDocument(ctx context.Context, id uint) error {
	return m.DeleteDocumentFunc(ctx, id)
}

func (m *mockRepository) GetById(ctx context.Context, id uint) (*Document, error) {
	return m.GetByIdFunc(ctx, id)
}

func (m *mockRepository) GetAllDocuments(ctx context.Context, pagination db.Pagination) (*db.Pagination, error) {
	return m.GetAllDocumentsFunc(ctx, pagination)
}


func TestCreateDocumentHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := &mockRepository{
		CreateDocumentFunc: func(ctx context.Context, doc *Document) error {
			assert.Equal(t, "Test Title", doc.Title)
			assert.Equal(t, "Tester", doc.Author)
			assert.Equal(t, "Test content", doc.Content)
			return nil
		},
	}

	svc := &service{repo: mockRepo}

	router := gin.Default()
	router.POST("/api/v1/document", svc.CreateDocument)

	body := `{"title":"Test Title","author":"Tester","content":"Test content"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/document", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}


func TestUpdateDocumentHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := &mockRepository{
		UpdateDocumentFunc: func(ctx context.Context, doc *Document) error {
			assert.Equal(t, uint(1), doc.ID) // ID of the document to update
			assert.Equal(t, "Updated Title", doc.Title)
			assert.Equal(t, "Updated Author", doc.Author)
			assert.Equal(t, "Updated Content", doc.Content)
			return nil
		},
	}

	svc := &service{repo: mockRepo}

	router := gin.Default()
	router.PUT("/api/v1/document/:id", svc.UpdateDocument)

	body := `{"title":"Updated Title","author":"Updated Author","content":"Updated Content"}`
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/document/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestDeleteDocumentHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := &mockRepository{
		DeleteDocumentFunc: func(ctx context.Context, id uint) error {
			assert.Equal(t, uint(1), id) 
			return nil
		},
	}

	svc := &service{repo: mockRepo}

	router := gin.Default()
	router.DELETE("/api/v1/document/:id", svc.DeleteDocument)

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/document/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}


func TestGetDocumentByIdHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := &mockRepository{
		GetByIdFunc: func(ctx context.Context, id uint) (*Document, error) {
			assert.Equal(t, uint(1), id) 
			return &Document{
				ID:      1,
				Title:   "Test Title",
				Author:  "Test Author",
				Content: "Test content",
			}, nil
		},
	}

	svc := &service{repo: mockRepo}

	router := gin.Default()
	router.GET("/api/v1/document/:id", svc.GetDocument)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/document/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Test Title")
}
