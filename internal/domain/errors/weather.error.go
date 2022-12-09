package errors

type NotFoundLocationError struct {
}

func (e *NotFoundLocationError) Error() string {
	return "Not found location"
}
