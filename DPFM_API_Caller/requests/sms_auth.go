package requests

type SMSAuth struct {
	UserID				string	`json:"UserID"`
	MobilePhoneNumber	string	`json:"MobilePhoneNumber"`
	AuthenticationCode	int		`json:"AuthenticationCode"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
