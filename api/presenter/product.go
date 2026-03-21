package presenter

import "mime/multipart"

type CreateUpdateRequestProduct struct {
	Name  string                `form:"name" validate:"required,min=3"`
	Photo *multipart.FileHeader `form:"photo"`
}
