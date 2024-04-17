package persistence

type RecordNotFoundError struct{}

func (RecordNotFoundError) Error() string {
	return "Record not found"
}
