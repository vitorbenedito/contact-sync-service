package domains

type List struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Contact struct {
		Company  string `json:"company"`
		Address1 string `json:"address1"`
		City     string `json:"city"`
		State    string `json:"state"`
		Zip      string `json:"zip"`
		Country  string `json:"country"`
		Phone    string `json:"phone"`
	} `json:"contact"`
	PermissionReminder string `json:"permission_reminder"`
	CampaignDefaults   struct {
		FromName  string `json:"from_name"`
		FromEmail string `json:"from_email"`
		Subject   string `json:"subject"`
		Language  string `json:"language"`
	} `json:"campaign_defaults"`
	EmailTypeOption bool `json:"email_type_option"`
}
