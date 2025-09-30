package validation

type Validation interface {
	Validate(field any) error
}
