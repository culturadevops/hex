package infra

type Repository interface {
	SetDriver(d Driver) error
	CreateField(filter interface{}, v interface{}) error
}
