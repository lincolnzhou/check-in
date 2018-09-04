package errors

const (
	// Ring
	RetRingEmpty = 2000
	RetRingFull  = 2001
)

var (
	ErrorRingEmpty = Error(RetRingEmpty)
	ErrorRingFull  = Error(RetRingFull)
)

type Error int

func (e Error) Error() string {
	return errorMsg[int(e)]
}

var errorMsg = map[int]string{
	//Ring
	RetRingEmpty: "check ring buffer empty",
	RetRingFull:  "check ring buffer full",
}
