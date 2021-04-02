package errors

var MsgMap = map[ErrorType]string{
	OK:               "OK",
	Canceled:         "Canceled",
	Unknown:          "Unknown",
	DeadlineExceeded: "Deadline exceeded",

	InvalidEmail:                  "Invalid email",
	InvalidPassword:               "Invalid password",
	EmailNotExists:                "Email does not exists",
	EmailAlreadyExists:            "Email already exists",
	PasswordMismatch:              "Password and confirm password must match",
	InvalidCredentials:            "invalid credentials",
	UpdateUser:                    "Cannot update user profile",
	ChangePasswordSame:            "New password and old password are the same",
	ChangePasswordOldPwdNotSame:   "Old password doesn't match",
	ForgotPasswordCode:            "Invalid forgot password errorType",
	InvalidVerificationCode:       "Invalid verification errorType",
	UserVerificationAlreadyExists: "User verification already exists",
	UserVerificationNotFound:      "User verification not found",
	UserVerificationStatusInvalid: "User verification status invalid",
	InvalidPhone:                  "Invalid phone",
	PhoneAlreadyExists:            "Phone already exists",
	NotAllowUpdateEmail:           "Not allow update email",

	InvalidMessage:      "Invalid message",
	InvalidArgument:     "Invalid argument",
	InternalServerError: "Internal server error",
	InvalidLimit:        "Invalid limit",
	InvalidPage:         "Invalid page",
	SystemError:         "System error",
	PermissionDenied:    "Permission denied",
	InvalidAdminKey:     "Admin key is invalid",
	InvalidJSON:         "JSON is invalid",
	FileNotFound:        "File is not found",
	InvalidMIME:         "Invalid MIME type",
	UploadFileFail:      "Upload file is fail",
	Marshal:             "Marshal fail",
	Unmarshal:           "Unmarshal fail",
	InvalidDateTime:     "Datetime is invalid",

	MongoConnection:    "Mongodb connection error",
	MongoCreate:        "Mongodb create model error",
	MongoRead:          "Mongodb read model error",
	MongoUpdate:        "Mongodb update model error",
	MongoDelete:        "Mongodb delete model error",
	MysqlConnection:    "Mysql connection error",
	MysqlDBEmpty:       "Mysql db is empty",
	MysqlDBAutoMigrate: "Mysql db auto migrate error",
	MysqlCreate:        "Mysql create model error",
	MysqlRead:          "Mysql read model error",
	MysqlUpdate:        "Mysql update model error",
	MysqlDelete:        "Mysql delete model error",
	MysqlAfterSave:     "Mysql after_save hook model error",
	SolrConnection:     "Solr connection error",
	SolrRead:           "Solr read error",
	SolrCreate:         "Solr create error",
	SolrUpdate:         "Solr update error",
	SolrDelete:         "Solr delete error",
	SolrCommit:         "Solr commit error",
	RedisConnection:    "Redis connection error",
	RedisGet:           "Redis get error",
	RedisSet:           "Redis set error",
	RedisRemove:        "Redis remove error",
}

func GetMsg(status int) string {
	msg, ok := MsgMap[ErrorType(status)]
	if ok {
		return msg
	}
	return MsgMap[Unknown]
}
