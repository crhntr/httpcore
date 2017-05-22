package httpcore

type StatusGetter interface {
	Status() int
}

type Identifier interface {
	GetID() string
	GetType() string
}
