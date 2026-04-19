package post

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

func (h *Handler) Create(ctx context.Context, c *app.RequestContext) {
	uidStr, _ := c.Get("user_id")
	authorID, _ := uuid.Parse(uidStr.(string))

	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.H{"error": err.Error()})
		return
	}

	res, err := h.svc.CreatePost(ctx, authorID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) Get(ctx context.Context, c *app.RequestContext) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{"error": "invalid id"})
		return
	}

	res, err := h.svc.GetPost(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) List(ctx context.Context, c *app.RequestContext) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	res, err := h.svc.ListPosts(ctx, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Delete(ctx context.Context, c *app.RequestContext) {
	uidStr, _ := c.Get("user_id")
	authorID, _ := uuid.Parse(uidStr.(string))

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{"error": "invalid id"})
		return
	}

	if err := h.svc.DeletePost(ctx, id, authorID); err != nil {
		c.JSON(http.StatusForbidden, utils.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
