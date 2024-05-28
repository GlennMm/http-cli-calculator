package utils

type Err struct {
  Message string 
  Code int 
}

func NewErr(message string, code int) *Err {
  return &Err{Message: message, Code: code}
}

func (e *Err) Error() string {
  return e.Message
}
