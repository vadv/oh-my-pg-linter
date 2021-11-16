package rules

type response struct {
	passed  bool
	message []byte
}

func (r *response) Passed() bool {
	return r.passed
}

func (r *response) Message() []byte {
	return r.message
}
