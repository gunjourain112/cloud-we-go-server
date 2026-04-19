package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Handler struct {
	svc      Service
	validate *validator.Validate
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		svc:      svc,
		validate: validator.New(),
	}
}

func (h *Handler) Create(c *gin.Context) {
	postID, _ := uuid.Parse(c.Param("id"))
	uidStr, _ := c.Get("user_id")
	authorID, _ := uuid.Parse(uidStr.(string))

	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.svc.CreateComment(c.Request.Context(), postID, authorID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) List(c *gin.Context) {
	postID, _ := uuid.Parse(c.Param("id"))

	res, err := h.svc.GetCommentsByPost(c.Request.Context(), postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Reply(c *gin.Context) {
	commentID, _ := bson.ObjectIDFromHex(c.Param("cid"))
	uidStr, _ := c.Get("user_id")
	authorID, _ := uuid.Parse(uidStr.(string))

	var req ReplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.AddReply(c.Request.Context(), commentID, authorID, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) Delete(c *gin.Context) {
	commentID, _ := bson.ObjectIDFromHex(c.Param("cid"))
	uidStr, _ := c.Get("user_id")
	authorID, _ := uuid.Parse(uidStr.(string))

	if err := h.svc.DeleteComment(c.Request.Context(), commentID, authorID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
