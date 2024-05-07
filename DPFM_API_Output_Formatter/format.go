package dpfm_api_output_formatter

import (
	"data-platform-api-authenticator-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToUser(rows *sql.Rows) (*[]User, error) {
	defer rows.Close()
	user := make([]User, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.User{}

		err := rows.Scan(
			&pm.UserID,
			&pm.BusinessPartner,
			&pm.Password,
			&pm.Qos,
			&pm.IsEncrypt,
			&pm.Language,
			&pm.LastLoginDate,
			&pm.LastLoginTime,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &user, err
		}

		data := pm
		user = append(user, User{
			UserID:					data.UserID,
			BusinessPartner:		data.BusinessPartner,
			Password:				data.Password,
			Qos:					data.Qos,
			IsEncrypt:				data.IsEncrypt,
			Language:				data.Language,
			LastLoginDate:			data.LastLoginDate,
			LastLoginTime:			data.LastLoginTime,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &user, nil
	}

	return &user, nil
}

func ConvertToSMSAuth(rows *sql.Rows) (*[]SMSAuth, error) {
	defer rows.Close()
	sMSAuth := make([]SMSAuth, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SMSAuth{}

		err := rows.Scan(
			&pm.UserID,
			&pm.MobilePhoneNumber,
			&pm.AuthenticationCode,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &sMSAuth, err
		}

		data := pm
		sMSAuth = append(sMSAuth, SMSAuth{
			UserID:                 data.UserID,
			MobilePhoneNumber:		data.MobilePhoneNumber,
			AuthenticationCode:		data.AuthenticationCode,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &sMSAuth, nil
	}

	return &sMSAuth, nil
}

func ConvertToGoogleAccountAuth(rows *sql.Rows) (*[]GoogleAccountAuth, error) {
	defer rows.Close()
	googleAccountAuth := make([]GoogleAccountAuth, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.GoogleAccountAuth{}

		err := rows.Scan(
			&pm.UserID,
			&pm.EmailAddress,
			&pm.GoogleID,
			&pm.AccessToken,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &googleAccountAuth, err
		}

		data := pm
		googleAccountAuth = append(googleAccountAuth, GoogleAccountAuth{
			UserID:                 data.UserID,
			EmailAddress:			data.EmailAddress,
			GoogleID:				data.GoogleID,
			AccessToken:			data.AccessToken,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &googleAccountAuth, nil
	}

	return &googleAccountAuth, nil
}

func ConvertToInstagramAuth(rows *sql.Rows) (*[]InstagramAuth, error) {
	defer rows.Close()
	instagramAuth := make([]InstagramAuth, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.InstagramAuth{}

		err := rows.Scan(
			&pm.UserID,
			&pm.InstagramID,
			&pm.AccessToken,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &instagramAuth, err
		}

		data := pm
		instagramAuth = append(instagramAuth, InstagramAuth{
			UserID:                 data.UserID,
			InstagramID:			data.InstagramID,
			AccessToken:			data.AccessToken,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &instagramAuth, nil
	}

	return &instagramAuth, nil
}
