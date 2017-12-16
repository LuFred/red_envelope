package core

// DbError records an error and the operation.
type DbError struct {
	Op     string
	Entity string
	Err    error
}

func (e *DbError) Error() string { return e.Op + "  " + e.Entity + " : " + e.Err.Error() }
