package jsons

var jsonData = `{
    "name": "Vitor Benedito",
    "contact": {
        "company": "VNB",
        "address1": "Rua Bartolomeu",
        "city": "Sao Jose Dos Campos",
        "state": "SP",
        "zip": "12459-590",
        "country": "BR",
        "phone": "+5515338901092"
    },
    "permission_reminder": "You are receiving this email because you signed up on our website.",
    "campaign_defaults": {
        "from_name": "VITOR",
        "from_email": "vitorbenedito@gmail.com",
        "subject": "",
        "language": "en"
    },
    "email_type_option": false
}`

func GetJson() string {
	return jsonData
}
