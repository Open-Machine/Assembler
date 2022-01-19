package myerrors

type CustomError struct {
	s       string
	errType int
}

// Error in the assembler programming
func NewAssemblerError(e error) *CustomError {
	return &CustomError{s: e.Error(), errType: assemblerError}
}

// Error in the assembly code
func NewCodeError(e error) *CustomError {
	return &CustomError{s: e.Error(), errType: codeError}
}

func (e *CustomError) Error() string {
	return e.s
}

func (e *CustomError) IsCodeError() bool {
	return e.errType == codeError
}

const assemblerError = 500
const codeError = 400
