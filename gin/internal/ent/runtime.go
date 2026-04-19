package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/post"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/schema"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/tag"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/user"
)

func init() {
	postFields := schema.Post{}.Fields()
	_ = postFields

	postDescTitle := postFields[1].Descriptor()

	post.TitleValidator = postDescTitle.Validators[0].(func(string) error)

	postDescBody := postFields[2].Descriptor()

	post.BodyValidator = postDescBody.Validators[0].(func(string) error)

	postDescCreatedAt := postFields[4].Descriptor()

	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)

	postDescUpdatedAt := postFields[5].Descriptor()

	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)

	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)

	postDescID := postFields[0].Descriptor()

	post.DefaultID = postDescID.Default.(func() uuid.UUID)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields

	tagDescName := tagFields[1].Descriptor()

	tag.NameValidator = tagDescName.Validators[0].(func(string) error)

	tagDescID := tagFields[0].Descriptor()

	tag.DefaultID = tagDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields

	userDescName := userFields[1].Descriptor()

	user.NameValidator = userDescName.Validators[0].(func(string) error)

	userDescEmail := userFields[2].Descriptor()

	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)

	userDescCreatedAt := userFields[4].Descriptor()

	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)

	userDescUpdatedAt := userFields[5].Descriptor()

	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)

	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)

	userDescID := userFields[0].Descriptor()

	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
