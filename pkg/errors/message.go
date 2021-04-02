package errors

var MsgMap = map[ErrorCode]string{
	ECOK:               "OK",
	ECCanceled:         "Canceled",
	ECUnknown:          "Unknown",
	ECDeadlineExceeded: "Deadline exceeded",

	ECInvalidEmail:                  "Invalid email",
	ECInvalidPassword:               "Invalid password",
	ECEmailNotExists:                "Email does not exists",
	ECEmailAlreadyExists:            "Email already exists",
	ECPasswordMismatch:              "Password and confirm password must match",
	ECUpdateUser:                    "Cannot update user profile",
	ECChangePasswordSame:            "New password and old password are the same",
	ECChangePasswordOldPwdNotSame:   "Old password doesn't match",
	ECForgotPasswordCode:            "Invalid forgot password code",
	ECPOSServerNotFound:             "POS server not found",
	ECPOSTokenInvalid:               "POS's token is invalid",
	ECPOSTokenParamsNotFound:        "POS's token params not found",
	ECInvalidVerificationCode:       "Invalid verification code",
	ECUserVerificationAlreadyExists: "User verification already exists",
	ECUserVerificationNotFound:      "User verification not found",
	ECUserVerificationStatusInvalid: "User verification status invalid",
	ECInvalidPhone:                  "Invalid phone",
	ECPhoneAlreadyExists:            "Phone already exists",
	ECNotAllowUpdateEmail:           "Not allow update email",

	ECInvalidMessage:      "Invalid message",
	ECInvalidArgument:     "Invalid argument",
	ECInternalServerError: "Internal server error",
	ECInvalidLimit:        "Invalid limit",
	ECInvalidPage:         "Invalid page",
	ECSystemError:         "System error",
	ECPermissionDenied:    "Permission denied",
	ECInvalidAdminKey:     "Admin key is invalid",
	ECInvalidJSON:         "JSON is invalid",
	ECFileNotFound:        "File is not found",
	ECInvalidMIME:         "Invalid MIME type",
	ECUploadFileFail:      "Upload file is fail",
	ECMarshal:             "Marshal fail",
	ECUnmarshal:           "Unmarshal fail",
	ECInvalidDateTime:     "Datetime is invalid",

	ECMongoConnection:         "Mongodb connection error",
	ECMongoCreate:             "Mongodb create model error",
	ECMongoRead:               "Mongodb read model error",
	ECMongoUpdate:             "Mongodb update model error",
	ECMongoDelete:             "Mongodb delete model error",
	ECPostgresqlConnection:    "Postgresql connection error",
	ECPostgresqlDBEmpty:       "Postgresql db is empty",
	ECPostgresqlDBAutoMigrate: "Postgresql db auto migrate error",
	ECPostgresqlCreate:        "Postgresql create model error",
	ECPostgresqlRead:          "Postgresql read model error",
	ECPostgresqlUpdate:        "Postgresql update model error",
	ECPostgresqlDelete:        "Postgresql delete model error",
	ECPostgresqlAfterSave:     "Postgresql after_save hook model error",
	ECSolrConnection:          "Solr connection error",
	ECSolrRead:                "Solr read error",
	ECSolrCreate:              "Solr create error",
	ECSolrUpdate:              "Solr update error",
	ECSolrDelete:              "Solr delete error",
	ECSolrCommit:              "Solr commit error",
	ECRedisConnection:         "Redis connection error",
	ECRedisGet:                "Redis get error",
	ECRedisSet:                "Redis set error",
	ECRedisRemove:             "Redis remove error",
}

func GetMsg(status int) string {
	msg, ok := MsgMap[ErrorCode(status)]
	if ok {
		return msg
	}
	return MsgMap[ECUnknown]
}
