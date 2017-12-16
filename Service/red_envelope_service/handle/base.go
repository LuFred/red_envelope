package handle

//Server 红包服务
type Server struct{}

//HandleError records an error and the operation.
type HandleError struct {
	Op  string
	Err error
}

func (e *HandleError) Error() string { return e.Op + " : " + e.Err.Error() }
