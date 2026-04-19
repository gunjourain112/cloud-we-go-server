package auth

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/go-playground/validator/v10"
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

func (h *Handler) Register(ctx context.Context, c *app.RequestContext) {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.H{"error": err.Error()})
		return
	}

	if err := h.svc.Register(ctx, &req); err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) LoginInternal(ctx context.Context, req *LoginRequest) (*LoginInternalResponse, error) {
	if err := h.validate.Struct(req); err != nil {
		return nil, err
	}
	return h.svc.LoginInternal(ctx, req)
}

type LoginInternalResponse struct {
	UserID string
}
