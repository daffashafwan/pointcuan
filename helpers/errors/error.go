package errors

import "errors"

var (
	ErrNotFound                 = errors.New("data not found")
	ErrIDNotFound               = errors.New("id not found")
	ErrInvalidId                = errors.New("invalid id, id not numeric")
	ErrUserIdNotFound           = errors.New("user id not found")
	ErrEmailHasBeenRegister     = errors.New("email has been used")
	ErrUsernameAlreadyExisted   = errors.New("username already exist")
	ErrUsernameRequired         = errors.New("username is required")
	ErrPasswordRequired         = errors.New("password is required")
	ErrEmailNotValid            = errors.New("email is not valid")
	ErrEmailRequired            = errors.New("email is required")
	ErrInvalidDate              = errors.New("invalid date, date must be formed : yyyy-mm-dd")
	ErrFavoriteIsAlready        = errors.New("this coin is already favorited by user")
	ErrUsernamePasswordNotFound = errors.New("username or password empty")
	ErrInvalidAuthentication    = errors.New("authentication failed: invalid user credentials")
	ErrInvalidTokenCredential   = errors.New("token not found")
	ErrBadRequest               = errors.New("bad requests")
	ErrPasswordDidntMatch       = errors.New("password didnt match")
)
