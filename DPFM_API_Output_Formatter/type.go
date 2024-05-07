package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	User             	*[]User					`json:"User"`
	SMSAuth            	*[]SMSAuth				`json:"SMSAuth"`
	GoogleAccountAuth	*[]GoogleAccountAuth	`json:"GoogleAccountAuth"`
	InstagramAuth		*[]InstagramAuth		`json:"InstagramAuth"`
}

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
