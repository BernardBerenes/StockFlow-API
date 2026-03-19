package presenter

type CreateRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}
