package document

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	CreateDocument(ctx context.Context, document *Document) error
	UpdateDocument(ctx context.Context, document *Document) error
	DeleteDocument(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*Document, error)
	GetAllDocuments(ctx context.Context) ([]Document, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateDocument(ctx context.Context, document *Document) error {
	return nil
}

func (r *repository) UpdateDocument(ctx context.Context, document *Document) error {
	return nil
}

func (r *repository) DeleteDocument(ctx context.Context, id string) error {
	return nil
}

func (r *repository) GetById(ctx context.Context, id string) (*Document, error) {
	return &Document{}, nil
}

func (r *repository) GetAllDocuments(ctx context.Context) ([]Document, error) {
	return []Document{}, nil
}
