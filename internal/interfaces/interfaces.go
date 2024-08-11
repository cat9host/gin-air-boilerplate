package interfaces

// common
type RouteInterface struct {
	Method string
	Path   string
}

type RoutineResult struct {
	Error  error
	Result bool
}

type StatusCode int

type Result int

const (
	ResultOk  Result = 0
	ResultErr Result = 1
)

const (
	Success         StatusCode = 0
	ValidationError StatusCode = 1
	InternalError   StatusCode = 2
	ExternalError   StatusCode = 3
	AuthError       StatusCode = 4
	CriticalError   StatusCode = 9
)

type RestSuccess struct {
	Code   StatusCode `json:"code"`
	Result any        `json:"result"`
}

type RestError struct {
	Message string     `json:"message"`
	Error   string     `json:"error"`
	Code    StatusCode `json:"code"`
}

type HCResponse struct {
	Status bool   `json:"status"`
	MySql  Result `json:"mysql"`
}
