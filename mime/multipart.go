package mime

type MultipartForm struct {
	parts map[string][]byte
}

func GetMultipartForm() *MultipartForm {
	m := MultipartForm{parts: make(map[string][]byte)}

	return &m
}

func (m *MultipartForm) Get(key string) []byte {
	return m.parts[key]
}

func (m *MultipartForm) Set(key string, part []byte) {
	m.parts[key] = part
}
