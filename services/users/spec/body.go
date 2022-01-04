package spec

type (
	UserBody struct {
		Firstname string `json:"firstname" validate:"required"`
		Lastname  string `json:"lastname" validate:"required,oneof=1 2"`
	}
)
