package errors

type (
	// ErrorType : error type
	ErrorType int32
)

// common errors
const (
	// OK is returned on success.
	// Error errorType starts from 100 in order not to duplicate grpc error codes which starts from 0 to 15.
	OK ErrorType = iota + 100
	// Canceled indicates the operation was canceled (typically by the caller).
	Canceled
	// Unknown error. An example of where this error may be returned is
	// if a Status value received from another address space belongs to
	// an error-space that is not known in this address space. Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	Unknown
	// DeadlineExceeded means operation expired before completion.
	// For operations that change the state of the system, this error may be
	// returned even if the operation has completed successfully. For
	// example, a successful response from a server could have been delayed
	// long enough for the deadline to expire.
	DeadlineExceeded
)

// user errors
const (
	// InvalidEmail : invalid email
	InvalidEmail ErrorType = iota + 200
	// InvalidPassword : invalid password
	InvalidPassword
	// EmailNotExists : email doesn't exist
	EmailNotExists
	// EmailAlreadyExists : email already exists
	EmailAlreadyExists
	// PasswordMismatch : password and confirm password must match
	PasswordMismatch
	// InvalidCredentials : invalid credentials
	InvalidCredentials
	// UpdateUser : cannot update user profile
	UpdateUser
	// ChangePasswordSame : new password and old password are the same
	ChangePasswordSame
	// ChangePasswordOldPwdNotSame : old password doesn't match
	ChangePasswordOldPwdNotSame
	// ForgotPasswordCode : invalid forgot password errorType
	ForgotPasswordCode
	// InvalidVerificationCode : invalid verification errorType
	InvalidVerificationCode
	// UserVerificationAlreadyExists : user verification already exists
	UserVerificationAlreadyExists
	// UserVerificationNotFound : user verification not found
	UserVerificationNotFound
	// UserVerificationStatusInvalid : user verification status invalid
	UserVerificationStatusInvalid
	// InvalidPhone : invalid phone
	InvalidPhone
	// PhoneAlreadyExists : phone already exists
	PhoneAlreadyExists
	// NotAllowUpdateEmail : not allow update email
	NotAllowUpdateEmail
)

// server errors
const (
	// InvalidMessage : invalid message
	InvalidMessage ErrorType = iota + 500
	// InvalidArgument : invalid argument
	InvalidArgument
	// InternalServerError : internal server error
	InternalServerError
	// InvalidLimit : invalid limit
	InvalidLimit
	// InvalidPage : invalid page
	InvalidPage
	// SystemError : system error
	SystemError
	// PermissionDenied : permission denied
	PermissionDenied

	// InvalidAdminKey : admin key is invalid
	InvalidAdminKey

	// InvalidJSON : JSON is invalid
	InvalidJSON
	// FileNotFound : file is not found
	FileNotFound
	// InvalidMIME : invalid MIME type
	InvalidMIME
	// UploadFileFail : upload file is fail
	UploadFileFail
	// Marshal : marshal fail
	Marshal
	// Unmarshal : unmarshal fail
	Unmarshal

	// InvalidDateTime : DateTime is invalid
	InvalidDateTime

	// HashPasswordFail : DateTime is invalid
	HashPasswordFail
)

// third - party errors
const (
	// MongoConnection : mongodb connection error
	MongoConnection ErrorType = iota + 1000
	// MongoCreate : mongodb create model error
	MongoCreate
	// MongoRead : mongodb read model error
	MongoRead
	// MongoUpdate : mongodb update model error
	MongoUpdate
	// MongoDelete : mongodb delete model error
	MongoDelete

	// MysqlConnection : mysql connection error
	MysqlConnection
	// MysqlDBEmpty : mysql db is empty
	MysqlDBEmpty
	// MysqlDBAutoMigrate : mysql db auto migrate error
	MysqlDBAutoMigrate
	// MysqlCreate : mysql create model error
	MysqlCreate
	// MysqlRead : mysql read model error
	MysqlRead
	// MysqlUpdate : mysql update model error
	MysqlUpdate
	// MysqlDelete : mysql delete model error
	MysqlDelete
	// MysqlAfterSave : mysql after_save hook model error
	MysqlAfterSave

	// SolrConnection : solr connection error
	SolrConnection
	// SolrRead : solr read error
	SolrRead
	// SolrCreate : solr create error
	SolrCreate
	// SolrUpdate : solr update error
	SolrUpdate
	// SolrDelete : solr delete error
	SolrDelete
	// SolrCommit : solr commit error
	SolrCommit

	// RedisConnection : redis connection error
	RedisConnection
	// RedisGet : redis get error
	RedisGet
	// RedisSet : redis set error
	RedisSet
	// RedisRemove : redis remove error
	RedisRemove
)
