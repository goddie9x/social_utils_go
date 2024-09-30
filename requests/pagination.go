package requests

type PaginationRequestInterface interface {
	GetPage() int
	GetSize() int
	GetOffset() int
}

type PaginationRequest struct {
	page int
	size int
}

func (r *PaginationRequest) GetPage() int {
	return r.page
}
func (r *PaginationRequest) GetSize() int {
	return r.size
}
func (r *PaginationRequest) GetOffset() int {
	return (r.GetPage() - 1) * r.GetSize()
}
