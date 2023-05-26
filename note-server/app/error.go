package app

type AppStatus struct {
	Message string `json:"errmsg"`
	Code    int    `json:"errcode"`
}

type AppError struct {
	Error   error
	Message string
	Code    int
}

func (a *AppError) Status() *AppStatus {
	return &AppStatus{
		Message: a.Message,
		Code:    a.Code,
	}
}
