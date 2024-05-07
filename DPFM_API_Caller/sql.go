package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-authenticator-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-authenticator-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var user *[]dpfm_api_output_formatter.User
	var sMSAuth *[]dpfm_api_output_formatter.SMSAuth
	var googleAccountAuth *[]dpfm_api_output_formatter.GoogleAccountAuth
	var instagramAuth *[]dpfm_api_output_formatter.InstagramAuth

	for _, fn := range accepter {
		switch fn {
		case "User":
			func() {
				user = c.User(mtx, input, output, errs, log)
			}()
		case "Users":
			func() {
				user = c.Users(mtx, input, output, errs, log)
			}()
		case "SMSAuth":
			func() {
				sMSAuth = c.SMSAuth(mtx, input, output, errs, log)
			}()
		case "GoogleAccountAuth":
			func() {
				googleAccountAuth = c.GoogleAccountAuth(mtx, input, output, errs, log)
			}()
		case "InstagramAuth":
			func() {
				instagramAuth = c.InstagramAuth(mtx, input, output, errs, log)
			}()

		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		User:              user,
		SMSAuth:           sMSAuth,
		GoogleAccountAuth: googleAccountAuth,
		InstagramAuth:     instagramAuth,
	}

	return data
}

func (c *DPFMAPICaller) User(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.User {
	where := fmt.Sprintf("WHERE user.UserID = \"%s\"", input.User.UserID)

	if input.User.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND user.IsMarkedForDeletion = %v", where, *input.User.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_authenticator_user_data AS user
		` + where + ` ORDER BY user.IsMarkedForDeletion ASC, user.UserID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToUser(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Users(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.User {
	where := "WHERE 1 = 1"
	if input.User.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND user.IsMarkedForDeletion = %v", where, *input.User.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_authenticator_user_data AS user
		` + where + ` ORDER BY user.IsMarkedForDeletion ASC, user.UserID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToUser(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) SMSAuth(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SMSAuth {
	var args []interface{}
	userID := input.User.UserID
	sMSAuth := input.User.SMSAuth

	cnt := 0
	args = append(args, userID)

	for _, v := range sMSAuth {
		args = append(args, v.UserID)
		cnt++
	}
	repeat := strings.Repeat("(?),", cnt-1) + "(?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_authenticator_sms_auth_data
		WHERE (UserID) IN ( `+repeat+` ) 
		ORDER BY UserID ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToSMSAuth(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) GoogleAccountAuth(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.GoogleAccountAuth {
	var args []interface{}
	userID := input.User.UserID
	googleAccountAuth := input.User.GoogleAccountAuth

	cnt := 0
	for _, _ = range googleAccountAuth {
		args = append(args, userID)
		cnt++
	}
	repeat := strings.Repeat("(?),", cnt-1) + "(?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_authenticator_google_account_auth_data
		WHERE (UserID) IN ( `+repeat+` ) 
		ORDER BY UserID ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGoogleAccountAuth(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) InstagramAuth(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.InstagramAuth {
	var args []interface{}
	userID := input.User.UserID
	instagramAuth := input.User.InstagramAuth

	cnt := 0
	for _, _ = range instagramAuth {
		args = append(args, userID)
		cnt++
	}
	repeat := strings.Repeat("(?),", cnt-1) + "(?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_authenticator_instagram_auth_data
		WHERE (UserID) IN ( `+repeat+` ) 
		ORDER BY UserID ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToInstagramAuth(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
