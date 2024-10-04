package requests

const (
	default_page = 1
	default_size = 10
)

type PaginationRequestInterface interface {
	GetPage() int
	GetSize() int
	GetOffset() int
}

type PaginationRequest struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func (r PaginationRequest) GetPage() int {
	if r.Page < default_page {
		r.Page = default_page
	}
	return r.Page
}
func (r PaginationRequest) GetSize() int {
	if r.Size < 1 {
		r.Size = default_size
	}
	return r.Size
}
func (r PaginationRequest) GetOffset() int {
	return (r.GetPage() - 1) * r.GetSize()
}
