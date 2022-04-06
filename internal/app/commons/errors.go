package commons

import (
	"errors"
)

// ErrDBConn error type for Error DB Connection
var ErrDBConn = errors.New("ErrDBConn")

// ErrCacheConn error type for Error Cache Connection
var ErrCacheConn = errors.New("ErrCacheConn")

var ErrMapping = errors.New("MappingError")

var ErrEmailExists = errors.New("EmailExistError")

var ErrInvalidToken = errors.New("ErrInvalidToken")

var ErrPhoneExistError = errors.New("PhoneExistError")

var ErrHashPassword = errors.New("HashPasswordError")

var ErrAuthorization = errors.New("AuthorizationError")

var ErrParsingBody = errors.New("ParsingBodyError")

var ErrInvalidCredential = errors.New("InvalidCredentialError")

var ErrInvalidData = errors.New("InvalidValueError")

var ErrJWTGenerate = errors.New("JWTGenerateError")

var ErrUpdate = errors.New("Update Error")

var ErrCreateUser = errors.New("Createe User Error")

var ErrGetAll = errors.New("Get All Error")

var ErrGetUserByID = errors.New("Get User By id Error")

var ErrDeleteUser = errors.New("Delete User Error")
