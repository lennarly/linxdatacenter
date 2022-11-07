package product

type Storage struct {
	file string
}

func New(file string) *Storage {
	return &Storage{file: file}
}
