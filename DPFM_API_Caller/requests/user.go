package requests

type User struct {
	UserID				string	`json:"UserID"`
	BusinessPartner		int		`json:"BusinessPartner"`
	Password			string	`json:"Password"`
	Qos					string	`json:"Qos"`
	IsEncrypt			bool	`json:"IsEncrypt"`
	Language			string	`json:"Language"`
	LastLoginDate		*string	`json:"LastLoginDate"`
	LastLoginTime		*string	`json:"LastLoginTime"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
