package facts

type CreateFactRequest struct {
	Body string
	Info string
}

type UpdateFactRequest struct {
	Body    string
	Info    string
	Updated string
}
