package requests

type GoogleAccountAuth struct {
	UserID				string	`json:"UserID"`
	EmailAddress		string	`json:"EmailAddress"`
	GoogleID			string	`json:"GoogleID"`
	AccessToken			string	`json:"AccessToken"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
