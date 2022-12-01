package mime

// MultipartForm is used for accessing and setting parts of a MultipartForm request by each part's name value
type MultipartForm struct {
	parts map[string]*Part
}

// GetMultipartForm initializes a *MultipartForm
func GetMultipartForm() *MultipartForm {
	m := MultipartForm{parts: make(map[string]*Part)}

	return &m
}

// Get returns the form Part stored in the MultipartForm with the given key
func (m *MultipartForm) Get(key string) *Part {
	return m.parts[key]
}

// Set parses a new *Part and sets it equal to the key given. The key must equal the value of the
// name value of the Content-Disposition header.Header of the provided part
func (m *MultipartForm) Set(key string, part []byte) {
	m.parts[key] = ParsePart(part)
}
