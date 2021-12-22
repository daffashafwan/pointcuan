package errors

import "errors"

var (
	ErrNotFound                 = errors.New("data not found")
	ErrIDNotFound               = errors.New("id not found")
	ErrInvalidId                = errors.New("invalid id, id not numeric")
	ErrUserIdNotFound           = errors.New("user id not found")
	ErrEmailHasBeenRegister     = errors.New("email has been used")
	ErrPasswordRequired         = errors.New("password is required")
	ErrEmailNotValid            = errors.New("email is not valid")
	ErrEmailRequired            = errors.New("email is required")
	ErrCoinNotFound             = errors.New("coin not found")
	ErrInvalidDate              = errors.New("invalid date, date must be formed : yyyy-mm-dd")
	ErrFavoriteIsAlready        = errors.New("this coin is already favorited by user")
	ErrUsernamePasswordNotFound = errors.New("username or password empty")
	ErrInvalidAuthentication    = errors.New("authentication failed: invalid user credentials")
	ErrInvalidTokenCredential   = errors.New("token not found or expired")
	ErrBadRequest               = errors.New("bad requests")
	ErrInvalidPayload           = errors.New("invalid payload")
	ErrHasBeenVerified          = errors.New("transaction already verified")
	ErrExpiredConfirm           = errors.New("expired payment confirmation")
	ErrTypeTransaction          = errors.New("type transaction not valid")
	ErrCoinRequired             = errors.New("coin symbol is required")
	ErrCoinNotEnough            = errors.New("coin not enough to sell")
	ErrQtyRequired              = errors.New("qty of coin is required")
	ErrWalletNotEnough          = errors.New("wallet is not enough")
	ErrCoinNotEnoughVerify      = errors.New("coin is not enough to sell")
	ErrWalletNotEnoughVerify    = errors.New("wallet is not enough, let's top up before expired")
	ErrUserIdRequired           = errors.New("qty of coin is required")
	ErrNothingDestroy           = errors.New("no data found to delete")
	ErrTypeActivity             = errors.New("type must be only 'income' and 'expense'")
	ErrInsufficientPermission   = errors.New("insufficient permission")
)