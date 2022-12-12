package sms

type CreateSMSRequest struct {
	Body       string
	From       string
	DateSent   string
	MessageSID string
}
