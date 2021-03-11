package infra

type Service interface {
	NewService(r Repository) *Service
	ListAllFor(v interface{}, fields interface{}) error
	GetField(v interface{}) error
	CreateField(filter interface{}, v interface{}) error
	Get(v interface{}, where string) error
	Create(v interface{}, where string) error
	Update(filter interface{}, update interface{}) error
}
