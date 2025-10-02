package messages

// Error response constants
const (
	ERR_ADD_GROUP_FAILED  = "Failed to add group"
	ERR_ADD_POLICY_FAILED = "Failed to add policy"
	ERR_INTERNAL_SERVER   = "Internal server error"
	ERR_VALIDATION        = "Field validation failed"
	ERR_FORBIDDEN         = "forbidden"
)

// Success message constants
const (
	SUCCESS_GROUP_ADDED    = "Group added successfully"
	SUCCESS_GROUP_DELETED  = "Group deleted successfully"
	SUCCESS_GROUP_GET_LIST = "Group get list successfully"

	SUCCESS_POLICY_ADDED    = "Policy added successfully"
	SUCCESS_POLICY_DELETED  = "Policy deleted successfully"
	SUCCESS_POLICY_GET_LIST = "Policy get list successfully"

	SUCCESS_VERIFY = "Verify"
)
