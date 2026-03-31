package presenter

type PaginateRequest struct {
	Page int `query:"page" validate:"min=1"`
	Size int `query:"size" validate:"min=10,max=100"`
}
