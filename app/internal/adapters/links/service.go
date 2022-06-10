package links

type Service interface {
	CreateLink(url string) (key string, err error)
	GetLink(key string) (url string, err error)
}
