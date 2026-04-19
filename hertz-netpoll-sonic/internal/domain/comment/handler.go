package comment

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var testUserID = uuid.MustParse("00000000-0000-0000-0000-000000000001")

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
	postID, _ := uuid.Parse(c.Param("id"))
	authorID := testUserID

	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.H{"error": err.Error()})
		return
	}

	res, err := h.svc.CreateComment(ctx, postID, authorID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) List(ctx context.Context, c *app.RequestContext) {
	postID, _ := uuid.Parse(c.Param("id"))

	res, err := h.svc.GetCommentsByPost(ctx, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Reply(ctx context.Context, c *app.RequestContext) {
	commentID, _ := bson.ObjectIDFromHex(c.Param("cid"))
	authorID := testUserID

	var req ReplyRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.H{"error": err.Error()})
		return
	}

	if err := h.svc.AddReply(ctx, commentID, authorID, &req); err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) Delete(ctx context.Context, c *app.RequestContext) {
	commentID, _ := bson.ObjectIDFromHex(c.Param("cid"))
	authorID := testUserID

	if err := h.svc.DeleteComment(ctx, commentID, authorID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
