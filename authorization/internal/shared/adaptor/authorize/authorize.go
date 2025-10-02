package authorize

type Authorize interface {
	AddPolicy(policy ...interface{}) (bool, error)
	RemovePolicy(policy ...interface{}) (bool, error)
	GetAllPolicy() ([][]string, error)
	AddGroup(policy ...interface{}) (bool, error)
	RemoveGroup(policy ...interface{}) (bool, error)
	GetAllGroup() ([][]string, error)
	Verify(policy ...interface{}) (bool, error)
}
