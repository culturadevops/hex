package infra

type Driver interface {
	newDriver([]string) error
}
