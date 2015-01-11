package loafer

type Adaptor interface {
	Write([]byte) (int, error)
}
