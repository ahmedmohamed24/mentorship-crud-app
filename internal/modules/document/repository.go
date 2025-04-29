package document

import (
	"context"
	"time"

	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/db"
	"gorm.io/gorm"
)

type Repository interface {
	CreateDocument(ctx context.Context, document *Document) error
	UpdateDocument(ctx context.Context, document *Document) error
	DeleteDocument(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*Document, error)
	GetAllDocuments(ctx context.Context, pagination db.Pagination) (*db.Pagination, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateDocument(ctx context.Context, document *Document) error {
	return r.db.Create(document).Error
}

func (r *repository) UpdateDocument(ctx context.Context, document *Document) error {
	return r.db.Model(&Document{}).Where("id = ?", document.ID).Updates(Document{
		Title:     document.Title,
		Author:    document.Author,
		Content:   document.Content,
		UpdatedAt: time.Now().UTC(),
	}).Error
}

func (r *repository) DeleteDocument(ctx context.Context, id uint) error {
	return r.db.Model(&Document{}).Where("id = ?", id).Delete(&Document{}).Error
}

func (r *repository) GetById(ctx context.Context, id uint) (*Document, error) {
	var doc Document
	if err := r.db.First(&doc, id).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}

func (r *repository) GetAllDocuments(ctx context.Context, pagination db.Pagination) (*db.Pagination, error) {
	var categories []*Document
	err := r.db.Scopes(db.Paginate(&categories, &pagination, r.db)).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = categories

	return &pagination, nil
}
