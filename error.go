package acars

/* this type of error is thrown by the acars service */
type AcarsError struct {
	errorMessage	string
}

func (err *AcarsError) Error() string {
	return err.errorMessage
}