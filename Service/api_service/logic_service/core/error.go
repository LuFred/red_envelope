package core

// LogicError .
type LogicError struct {
	Code         int
	Message      string
	Type         int32 //0:系统错误|1：业务逻辑错误
	SeriviceName string
	Op           string
	Err          error
}

func (e *LogicError) Error() string { return e.SeriviceName + "  " + e.Op + " : " + e.Err.Error() }
