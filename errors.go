package go_pkg

type DomainError struct {
	Msg string
}

func (e *DomainError) Error() string {
	return e.Msg
}

func NewDomainError(msg string) error {
	return &DomainError{Msg: msg}
}
