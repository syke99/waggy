package mime

type MultipartForm struct {
	parts map[string]*Part
}

func GetMultipartForm() *MultipartForm {
	m := MultipartForm{parts: make(map[string]*Part)}

	return &m
}

func (m *MultipartForm) Get(key string) *Part {
	return m.parts[key]
}

func (m *MultipartForm) Set(key string, part []byte) {
	m.parts[key] = ParsePart(part)
}
