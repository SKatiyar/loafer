package loafer

type Formatter interface {
	Format(Fields) ([]byte, error)
}
