package persistence

type Storage interface {
	Set(alias, url string) error
	Get(alias string) (url string, found bool)
	Count() int
	Increment() error
}
