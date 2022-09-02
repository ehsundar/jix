package jix

import "errors"

var (
	ErrInvalidArgument    = errors.New("INVALID_ARGUMENT")
	ErrAlreadyExists      = errors.New("ALREADY_EXISTS")
	ErrPermissionDenied   = errors.New("PERMISSION_DENIED")
	ErrResourceExhausted  = errors.New("RESOURCE_EXHAUSTED")
	ErrFailedPrecondition = errors.New("FAILED_PRECONDITION")
	ErrAborted            = errors.New("ABORTED")
	ErrOutOfRange         = errors.New("OUT_OF_RANGE")
	ErrUnimplemented      = errors.New("UNIMPLEMENTED")
	ErrInternal           = errors.New("INTERNAL")
	ErrUnavailable        = errors.New("UNAVAILABLE")
	ErrDataLoss           = errors.New("DATA_LOSS")
	ErrUnauthenticated    = errors.New("UNAUTHENTICATED")

	ErrBadRequest                  = errors.New("bad request")
	ErrUnauthorized                = errors.New("unauthorized")
	ErrPaymentRequired             = errors.New("payment required")
	ErrForbidden                   = errors.New("forbidden")
	ErrNotFound                    = errors.New("NOT_FOUND")
	ErrMethodNotAllowed            = errors.New("method not allowed")
	ErrNotAcceptable               = errors.New("not acceptable")
	ErrProxyAuthenticationRequired = errors.New("proxy authentication required")
	ErrRequestTimeout              = errors.New("request timeout")
	ErrConflict                    = errors.New("conflict")
	ErrGone                        = errors.New("gone")
	ErrLengthRequired              = errors.New("length required")
	ErrPreconditionFailed          = errors.New("precondition failed")
	ErrPayloadTooLarge             = errors.New("payload too large")

	ErrInternalServerError = errors.New("internal server error")
	ErrNotImplemented      = errors.New("not implemented")
	ErrBadGateway          = errors.New("bad gateway")
	ErrServiceUnavailable  = errors.New("service unavailable")
	ErrGatewayTimeout      = errors.New("gateway timeout")
)

var errorToStatusMap = map[error]int{
	ErrInvalidArgument:    400,
	ErrAlreadyExists:      400,
	ErrPermissionDenied:   401,
	ErrResourceExhausted:  503,
	ErrFailedPrecondition: 412,
	ErrAborted:            406,
	ErrOutOfRange:         416,
	ErrUnimplemented:      501,
	ErrInternal:           500,
	ErrUnavailable:        503,
	ErrDataLoss:           507,
	ErrUnauthenticated:    401,

	ErrBadRequest:                  400,
	ErrUnauthorized:                401,
	ErrPaymentRequired:             402,
	ErrForbidden:                   403,
	ErrNotFound:                    404,
	ErrMethodNotAllowed:            405,
	ErrNotAcceptable:               406,
	ErrProxyAuthenticationRequired: 407,
	ErrRequestTimeout:              408,
	ErrConflict:                    409,
	ErrGone:                        410,
	ErrLengthRequired:              411,
	ErrPreconditionFailed:          412,
	ErrPayloadTooLarge:             413,

	ErrInternalServerError: 500,
	ErrNotImplemented:      501,
	ErrBadGateway:          502,
	ErrServiceUnavailable:  503,
	ErrGatewayTimeout:      504,
}
