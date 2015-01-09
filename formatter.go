package loafer

type Formatter interface {
	Format(Fields) ([]byte, error)
}

type defaultFormatter struct{}

func (d *defaultFormatter) Format(f Fields) ([]byte, error) {
	return nil, nil
}
