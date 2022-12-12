package models

type SMSMessage struct {
	AccountSid  string `json:"AccountSid"`
	Body        string `json:"Body"`
	From        string `json:"From"`
	MessageSid  string `json:"MessageSid"`
	NumMedia    string `json:"NumMedia"`
	NumSegments string `json:"NumSegments"`
	SmsSid      string `json:"SmsSid"`
	To          string `json:"To"`
}
