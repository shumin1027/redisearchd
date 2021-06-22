package http

// HTTP methods and their unique INTs
func MethodInt(s string) int {
	switch s {
	case MethodGet:
		return 0
	case MethodHead:
		return 1
	case MethodPost:
		return 2
	case MethodPut:
		return 3
	case MethodDelete:
		return 4
	case MethodConnect:
		return 5
	case MethodOptions:
		return 6
	case MethodTrace:
		return 7
	case MethodPatch:
		return 8
	default:
		return -1
	}
}

// HTTP methods were copied from net/http.
// Common HTTP methods.
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)
