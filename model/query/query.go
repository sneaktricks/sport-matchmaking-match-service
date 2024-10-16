package query

type PaginationParams struct {
	Page  uint `query:"page" validate:"omitempty,min=1"`
	Limit uint `query:"limit" validate:"omitempty,min=1"`
}
