package errors

var MsgMap = map[ErrorCode]string{
	ECOK:               "ok",
	ECCanceled:         "canceled",
	ECUnknown:          "unknown",
	ECDeadlineExceeded: "deadline exceeded",

	ECInvalidEmail:                  "invalid email",
	ECInvalidPassword:               "invalid password",
	ECEmailNotExists:                "email does not exists",
	ECEmailAlreadyExists:            "email already exists",
	ECPasswordMismatch:              "password and confirm password must match",
	ECUpdateUser:                    "cannot update user profile",
	ECChangePasswordSame:            "new password and old password are the same",
	ECChangePasswordOldPwdNotSame:   "old password doesn't match",
	ECForgotPasswordCode:            "invalid forgot password code",
	ECPOSServerNotFound:             "POS server not found",
	ECPOSTokenInvalid:               "POS's token is invalid",
	ECPOSTokenParamsNotFound:        "POS's token params not found",
	ECInvalidVerificationCode:       "invalid verification code",
	ECUserVerificationAlreadyExists: "user verification already exists",
	ECUserVerificationNotFound:      "user verification not found",
	ECUserVerificationStatusInvalid: "user verification status invalid",
	ECInvalidPhone:                  "invalid phone",
	ECPhoneAlreadyExists:            "phone already exists",
	ECNotAllowUpdateEmail:           "not allow update email",

	ECInvalidMessage:      "invalid message",
	ECInvalidArgument:     "invalid argument",
	ECInternalServerError: "internal server error",
	ECInvalidLimit:        "invalid limit",
	ECInvalidPage:         "invalid page",
	ECSystemError:         "system error",
	ECPermissionDenied:    "permission denied",
	ECInvalidAdminKey:     "admin key is invalid",
	ECInvalidJSON:         "JSON is invalid",
	ECFileNotFound:        "file is not found",
	ECInvalidMIME:         "invalid MIME type",
	ECUploadFileFail:      "upload file is fail",
	ECMarshal:             "marshal fail",
	ECUnmarshal:           "unmarshal fail",
	ECInvalidDateTime:     "datetime is invalid",

	ECMongoConnection:         "mongodb connection error",
	ECMongoCreate:             "mongodb create model error",
	ECMongoRead:               "mongodb read model error",
	ECMongoUpdate:             "mongodb update model error",
	ECMongoDelete:             "mongodb delete model error",
	ECPostgresqlConnection:    "postgresql connection error",
	ECPostgresqlDBEmpty:       "postgresql db is empty",
	ECPostgresqlDBAutoMigrate: "postgresql db auto migrate error",
	ECPostgresqlCreate:        "postgresql create model error",
	ECPostgresqlRead:          "postgresql read model error",
	ECPostgresqlUpdate:        "postgresql update model error",
	ECPostgresqlDelete:        "postgresql delete model error",
	ECPostgresqlAfterSave:     "postgresql after_save hook model error",
	ECSolrConnection:          "solr connection error",
	ECSolrRead:                "solr read error",
	ECSolrCreate:              "solr create error",
	ECSolrUpdate:              "solr update error",
	ECSolrDelete:              "solr delete error",
	ECSolrCommit:              "solr commit error",
	ECRedisConnection:         "redis connection error",
	ECRedisGet:                "redis get error",
	ECRedisSet:                "redis set error",
	ECRedisRemove:             "redis remove error",
}

func GetMsg(status int) string {
	msg, ok := MsgMap[ErrorCode(status)]
	if ok {
		return msg
	}
	return MsgMap[ECUnknown]
}
