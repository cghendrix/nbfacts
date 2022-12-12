package facts

type CreateFactRequest struct {
	Body string
	Info string
}

type CreateFactFromSMSRequest struct {
	Body  string
	Info  string
	SmsId string
}

type UpdateFactRequest struct {
	Body    string
	Info    string
	Updated string
}
