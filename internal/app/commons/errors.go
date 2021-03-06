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

var ErrQueryDB = errors.New("QueryDBError")

var ErrOutOfStock = errors.New("OutOfStockError")

var ErrAddCart = errors.New("AddCartError")

var ErrUpdateCart = errors.New("UpdateCartError")

var ErrDeleteCart = errors.New("DeleteCartError")

var ErrEmptyCart = errors.New("EmptyCartError")

var ErrInvalidCartItem = errors.New("InvalidCartItemError")

var ErrAddTransaction = errors.New("AddTransactionError")

var ErrNotFound = errors.New("NotFoundError")

var ErrUpdate = errors.New("UpdateError")

var ErrCreateUser = errors.New("CreateeUserError")

var ErrGetAll = errors.New("GetAllError")

var ErrGetUserByID = errors.New("GetUserByIdError")

var ErrDeleteUser = errors.New("DeleteUserError")
