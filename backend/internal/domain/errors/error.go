package errors

import "errors"

var (
	// General Errors
	ErrMissingFields = errors.New("missing required fields")
	ErrInvalidUUID = errors.New("invalid UUID format")
	
	// Register Errors
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidInviteCode = errors.New("invalid invite code")
	ErrInviteCodeLength = errors.New("invite code must be 6 characters")
	ErrNameLength = errors.New("name must be 4 to 24 characters")
	ErrPasswordLength = errors.New("password must be 8 to 24 characters")
	ErrInvalidPassword = errors.New("password must contain a uppercase letter and a number")
	ErrConfirmPasswordMismatch = errors.New("password and confirm password is not the same")
	ErrDuplicateEmail = errors.New("email already exists")

	// Login Errors
	ErrInvalidCredentials = errors.New("invalid credentials")

	// Update Interests Errors
	ErrInterestLength = errors.New("please select at least one interest")

	// Create Event Errors
	ErrInvalidEventType = errors.New("event type must be 'single' or 'recurring'")
	ErrInvalidTicketType = errors.New("ticket type must be 'free' or 'ticketed'")
	ErrInvalidStartDate = errors.New("start date must be in the future")
	ErrInvalidStartTime = errors.New("start time must be in the future")
	ErrMissingBanner = errors.New("banner image is required")
	ErrInvalidEndTime = errors.New("end time must be after start time")
	ErrBannerMaxSize = errors.New("banner image size exceeds 5MB limit")
	ErrInvalidBannerFormat = errors.New("banner image must be a valid image format (jpg, jpeg, png, gif)")
	ErrInvalidTicketPrice = errors.New("ticket price must be a positive integer")
	ErrInvalidTicketQuantity = errors.New("ticket quantity must be more than or equal to 1")


	// Supabase Errors 
	ErrSupabaseRequestFailed = errors.New("supabase request failed")
	ErrSupabaseBucketMissing = errors.New("missing SUPABASE_BUCKET environment variable")
	ErrSupabaseUrlMissing = errors.New("missing SUPABASE_URL environment variable")
	ErrSupabaseKeyMissing  = errors.New("missing SUPABASE_KEY environment variable")	
)