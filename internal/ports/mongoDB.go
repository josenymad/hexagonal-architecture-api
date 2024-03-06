package ports

type Database interface {
	PostData(data interface{}) error
}
