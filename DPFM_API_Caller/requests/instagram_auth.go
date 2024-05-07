package requests

type InstagramAuth struct {
	UserID				string	`json:"UserID"`
	InstagramID			string	`json:"InstagramID"`
	AccessToken			string	`json:"AccessToken"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
