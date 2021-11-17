package rules

type response struct {
	passed  bool
	message []byte
	query   string
}

func (r *response) Passed() bool {
	return r.passed
}
func (r *response) Query() *string {
	result := r.query
	if result == "" {
		return nil
	}
	return &result
}

func (r *response) Message() []byte {
	return r.message
}
