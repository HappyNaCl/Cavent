package errors

import "errors"

var (
	// Register Errors
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidInviteCode = errors.New("invalid invite code")
	ErrInviteCodeLength = errors.New("invite code must be 6 characters")
	ErrNameLength = errors.New("name must be 4 to 24 characters")
	ErrPasswordLength = errors.New("password must be 8 to 24 characters")
	ErrInvalidPassword = errors.New("password must contain a uppercase letter and a number")
	ErrConfirmPasswordMismatch = errors.New("password and confirm password is not the same")
	ErrDuplicateEmail = errors.New("email already exists")
)