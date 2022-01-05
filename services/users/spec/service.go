package spec

type (
	UserService interface {
		Find() (interface{}, error)
	}
)
