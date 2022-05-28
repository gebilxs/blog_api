package model

import "blog_api/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Paper *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
