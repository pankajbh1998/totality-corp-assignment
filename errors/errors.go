package errors

type ConstError string

func (e ConstError) Error() string {
	return string(e)
}

const (
	ErrContentNotFound        = ConstError("Content not found while fetching the user data")
	ErrIdCantBeZeroOrNegative = ConstError("ID must be integer and greater than zero")
)
