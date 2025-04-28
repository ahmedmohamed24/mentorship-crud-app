package document

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateDocument(*gin.Context)
	UpdateDocument(*gin.Context)
	DeleteDocument(*gin.Context)
	GetDocument(*gin.Context)
	GetAllDocuments(*gin.Context)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateDocument(c *gin.Context) {
}
func (s *service) UpdateDocument(c *gin.Context) {
}
func (s *service) DeleteDocument(c *gin.Context) {
}
func (s *service) GetDocument(c *gin.Context) {
}
func (s *service) GetAllDocuments(c *gin.Context) {
}
