package messages

// Error response constants
const (
	ERR_INVALID_FIELDS      = "Invalid fields provided"
	ERR_VALIDATION          = "Field validation failed"
	ERR_INTERNAL_SERVER     = "Internal server error"
	ERR_INVALID_CREDENTIALS = "Invalid username or password"
	ERR_INVALID_TOKEN       = "Invalid authentication token"
	ERR_INVALID_GOOGLE_CODE = "Invalid Google verification code"
)

// Success message constants
const (
	SUCCESS_LOGIN            = "Login successful"
	SUCCESS_LOGIN_TWO_FACTOR = "Login successful with two-factor authentication"

	SUCCESS_REGISTER = "Register successful"

	SUCCESS_PASSWORD_RESET_EMAIL = "Password reset email sent"
	SUCCESS_PASSWORD_RESET       = "Password successfully reset"

	SUCCESS_EMAIL_TWO_FACTOR      = "Email two-factor authentication successful"
	SUCCESS_SEND_EMAIL_TWO_FACTOR = "Email sent for two-factor authentication"

	SUCCESS_GOOGLE_TWO_FACTOR = "Google two-factor authentication successful"

	SUCCESS_ACCOUNT_CREATED   = "Account created successfully"
	SUCCESS_ACCOUNT_UPDATED   = "Account updated successfully"
	SUCCESS_ACCOUNT_DELETED   = "Account deleted successfully"
	SUCCESS_ACCOUNT_RETRIEVED = "Account details retrieved successfully"
)
