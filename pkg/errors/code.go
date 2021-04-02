package errors

var CodeMap = map[ErrorType]string{
	OK:               "OK",
	Canceled:         "CANCELED",
	Unknown:          "UNKNOWN",
	DeadlineExceeded: "DEADLINE_EXCEEDED",

	InvalidEmail:                  "INVALID_EMAIL",
	InvalidPassword:               "INVALID_PASSWORD",
	EmailNotExists:                "EMAIL_NOT_EXISTS",
	EmailAlreadyExists:            "EMAIL_ALREADY_EXISTS",
	PasswordMismatch:              "PASSWORD_MISMATCH",
	InvalidCredentials:            "INVALID_CREDENTIALS",
	UpdateUser:                    "UPDATE_USER_FAIL",
	ChangePasswordSame:            "CHANGE_PASSWORD_SAME",
	ChangePasswordOldPwdNotSame:   "CHANGE_PASSWORD_OLD_PWD_NOT_SAME",
	ForgotPasswordCode:            "FORGOT_PASSWORD_CODE",
	InvalidVerificationCode:       "INVALID_VERIFICATION_CODE",
	UserVerificationAlreadyExists: "USER_VERIFICATION_ALREADY_EXISTS",
	UserVerificationNotFound:      "USER_VERIFICATION_NOTFOUND",
	UserVerificationStatusInvalid: "USER_VERIFICATION_STATUS_INVALID",
	InvalidPhone:                  "INVALID_PHONE",
	PhoneAlreadyExists:            "PHONE_ALREADY_EXISTS",
	NotAllowUpdateEmail:           "NOT_ALLOW_UPDATE_EMAIL",

	InvalidMessage:      "INVALID_MESSAGE",
	InvalidArgument:     "INVALID_ARGUMENT",
	InternalServerError: "INTERNAL_SERVER_ERROR",
	InvalidLimit:        "INVALID_LIMIT",
	InvalidPage:         "INVALID_PAGE",
	SystemError:         "SYSTEM_ERROR",
	PermissionDenied:    "PERMISSION_DENIED",
	InvalidAdminKey:     "INVALID_ADMIN_KEY",
	InvalidJSON:         "INVALID_JSON",
	FileNotFound:        "FILE_NOT_FOUND",
	InvalidMIME:         "INVALID_MIME",
	UploadFileFail:      "UPLOAD_FILE_FAIL",
	Marshal:             "MARSHAL",
	Unmarshal:           "UNMARSHAL",
	InvalidDateTime:     "INVALID_DATETIME",

	MongoConnection:    "MONGO_CONNECTION",
	MongoCreate:        "MONGO_CREATE",
	MongoRead:          "MONGO_READ",
	MongoUpdate:        "MONGO_UPDATE",
	MongoDelete:        "MONGO_DELETE",
	MysqlConnection:    "MYSQL_CONNECTION",
	MysqlDBEmpty:       "MYSQL_DBEMPTY",
	MysqlDBAutoMigrate: "MYSQL_DBAUTOMIGRATE",
	MysqlCreate:        "MYSQL_CREATE",
	MysqlRead:          "MYSQL_READ",
	MysqlUpdate:        "MYSQL_UPDATE",
	MysqlDelete:        "MYSQL_DELETE",
	MysqlAfterSave:     "MYSQL_AFTERSAVE",
	SolrConnection:     "SOLR_CONNECTION",
	SolrRead:           "SOLR_READ",
	SolrCreate:         "SOLR_CREATE",
	SolrUpdate:         "SOLR_UPDATE",
	SolrDelete:         "SOLR_DELETE",
	SolrCommit:         "SOLR_COMMIT",
	RedisConnection:    "REDIS_CONNECTION",
	RedisGet:           "REDIS_GET",
	RedisSet:           "REDIS_SET",
	RedisRemove:        "REDIS_REMOVE",
}

func GetCode(status int) string {
	code, ok := CodeMap[ErrorType(status)]
	if ok {
		return code
	}
	return CodeMap[Unknown]
}
