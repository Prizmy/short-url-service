package storage

type Storage interface {
	Post(url string) (string, error)
	Get(shortURL string) (string, error)
}
