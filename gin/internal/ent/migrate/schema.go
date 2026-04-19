package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "body", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "author_id", Type: field.TypeUUID},
	}

	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}

	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
	}

	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}

	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}

	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}

	PostTagsColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeUUID},
		{Name: "tag_id", Type: field.TypeUUID},
	}

	PostTagsTable = &schema.Table{
		Name:       "post_tags",
		Columns:    PostTagsColumns,
		PrimaryKey: []*schema.Column{PostTagsColumns[0], PostTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_tags_post_id",
				Columns:    []*schema.Column{PostTagsColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_tags_tag_id",
				Columns:    []*schema.Column{PostTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}

	Tables = []*schema.Table{
		PostsTable,
		TagsTable,
		UsersTable,
		PostTagsTable,
	}
)

func init() {
	PostsTable.ForeignKeys[0].RefTable = UsersTable
	PostTagsTable.ForeignKeys[0].RefTable = PostsTable
	PostTagsTable.ForeignKeys[1].RefTable = TagsTable
}
