package document

import (
	"net/http"
	"strconv"

	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/db"
	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

type CreateDocumentRequest struct {
	Title   string `json:"title" binding:"required,min=3,max=100"`
	Author  string `json:"author" binding:"required,min=3,max=100"`
	Content string `json:"content" binding:"required,min=3,max=500"`
}

func (s *service) CreateDocument(c *gin.Context) {
	var req CreateDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.repo.CreateDocument(c, &Document{
		Title:   req.Title,
		Author:  req.Author,
		Content: req.Content,
	})
	if err != nil {
		log.WithField("IP", c.ClientIP()).Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternal.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

type UpdateDocumentRequest struct {
	Title   string `json:"title" binding:"required,min=3,max=100"`
	Author  string `json:"author" binding:"required,min=3,max=100"`
	Content string `json:"content" binding:"required,min=3,max=500"`
}

func (s *service) UpdateDocument(c *gin.Context) {
	id := c.Param("id")
	var req UpdateDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	documentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrDocumentNotFound.Error()})
		return
	}
	err = s.repo.UpdateDocument(c, &Document{
		ID:      uint(documentID),
		Title:   req.Title,
		Author:  req.Author,
		Content: req.Content,
	})
	if err != nil {
		log.WithField("IP", c.ClientIP()).Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternal.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
func (s *service) DeleteDocument(c *gin.Context) {
	id := c.Param("id")
	documentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrDocumentNotFound.Error()})
		return
	}
	err = s.repo.DeleteDocument(c, uint(documentID))
	if err != nil {
		log.WithField("IP", c.ClientIP()).Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternal.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func (s *service) GetDocument(c *gin.Context) {
	id := c.Param("id")
	documentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrDocumentNotFound.Error()})
		return
	}
	doc, err := s.repo.GetById(c, uint(documentID))
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrDocumentNotFound.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternal.Error()})
		return
	}
	c.JSON(http.StatusOK, doc)
}

type ListDocumentResponse struct {
	db.Pagination
	Data []Document `json:"data"`
}

func (s *service) GetAllDocuments(c *gin.Context) {
	pagination := db.GetRequestPagination(c)
	paginated, err := s.repo.GetAllDocuments(c, pagination)
	if err != nil {
		log.WithField("IP", c.ClientIP()).Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternal.Error()})
		return
	}
	c.JSON(http.StatusOK, paginated)
}
