package adapter

// Adapter write log to underly writer
type Adapter interface {
	Write([]byte) error
}
