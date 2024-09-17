package storage

type Log interface {
	Error(args ...interface{})
}
