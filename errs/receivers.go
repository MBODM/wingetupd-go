package errs

// Implement error interface, so ExpectedError becomes a typical error.

func (e *ExpectedError) Error() string {
	if e.Err != nil {
		return e.Msg + ": " + e.Err.Error()
	}
	return e.Msg
}

func (e *ExpectedError) Unwrap() error {
	return e.Err
}
