package domains

type Member struct {
	Email  string  `json:"email_address"`
	Fields *Fields `json:"merge_fields"`
	Status string  `json:"status"`
}

type Fields struct {
	FirstName string `json:"FNAME"`
	LastName  string `json:"LNAME"`
}
