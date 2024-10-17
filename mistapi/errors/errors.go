package errors

import (
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

// ErrorDeleteFailed is a custom error.
type ErrorDeleteFailed struct {
	https.ApiError
	Detail string    `json:"detail"`
	OrgId  uuid.UUID `json:"org_id"`
}

// NewErrorDeleteFailed is a constructor for ErrorDeleteFailed.
// It creates and returns a pointer to a new ErrorDeleteFailed instance with the given statusCode and body.
func NewErrorDeleteFailed(apiError https.ApiError) error {
	return &ErrorDeleteFailed{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ErrorDeleteFailed.
func (e ErrorDeleteFailed) Error() string {
	return fmt.Sprintf("ErrorDeleteFailed occured: %v", e.Message)
}

// ResponseDetailString is a custom error.
type ResponseDetailString struct {
	https.ApiError
	Detail *string `json:"detail,omitempty"`
}

// NewResponseDetailString is a constructor for ResponseDetailString.
// It creates and returns a pointer to a new ResponseDetailString instance with the given statusCode and body.
func NewResponseDetailString(apiError https.ApiError) error {
	return &ResponseDetailString{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseDetailString.
func (r ResponseDetailString) Error() string {
	return fmt.Sprintf("ResponseDetailString occured: %v", r.Message)
}

// ResponseHttp400 is a custom error.
type ResponseHttp400 struct {
	https.ApiError
	Detail *string `json:"detail,omitempty"`
}

// NewResponseHttp400 is a constructor for ResponseHttp400.
// It creates and returns a pointer to a new ResponseHttp400 instance with the given statusCode and body.
func NewResponseHttp400(apiError https.ApiError) error {
	return &ResponseHttp400{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseHttp400.
func (r ResponseHttp400) Error() string {
	return fmt.Sprintf("ResponseHttp400 occured: %v", r.Message)
}

// ResponseHttp404 is a custom error.
type ResponseHttp404 struct {
	https.ApiError
	Id *string `json:"id,omitempty"`
}

// NewResponseHttp404 is a constructor for ResponseHttp404.
// It creates and returns a pointer to a new ResponseHttp404 instance with the given statusCode and body.
func NewResponseHttp404(apiError https.ApiError) error {
	return &ResponseHttp404{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseHttp404.
func (r ResponseHttp404) Error() string {
	return fmt.Sprintf("ResponseHttp404 occured: %v", r.Message)
}

// ResponseLoginFailure is a custom error.
type ResponseLoginFailure struct {
	https.ApiError
	Detail     string  `json:"detail"`
	ForwardUrl *string `json:"forward_url,omitempty"`
}

// NewResponseLoginFailure is a constructor for ResponseLoginFailure.
// It creates and returns a pointer to a new ResponseLoginFailure instance with the given statusCode and body.
func NewResponseLoginFailure(apiError https.ApiError) error {
	return &ResponseLoginFailure{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseLoginFailure.
func (r ResponseLoginFailure) Error() string {
	return fmt.Sprintf("ResponseLoginFailure occured: %v", r.Message)
}

// ResponseSelfOauthLinkFailure is a custom error.
type ResponseSelfOauthLinkFailure struct {
	https.ApiError
	MError           string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// NewResponseSelfOauthLinkFailure is a constructor for ResponseSelfOauthLinkFailure.
// It creates and returns a pointer to a new ResponseSelfOauthLinkFailure instance with the given statusCode and body.
func NewResponseSelfOauthLinkFailure(apiError https.ApiError) error {
	return &ResponseSelfOauthLinkFailure{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseSelfOauthLinkFailure.
func (r ResponseSelfOauthLinkFailure) Error() string {
	return fmt.Sprintf("ResponseSelfOauthLinkFailure occured: %v", r.Message)
}

// ResponseHttp401Error is a custom error.
type ResponseHttp401Error struct {
	https.ApiError
	Detail *string `json:"detail,omitempty"`
}

// NewResponseHttp401Error is a constructor for ResponseHttp401Error.
// It creates and returns a pointer to a new ResponseHttp401Error instance with the given statusCode and body.
func NewResponseHttp401Error(apiError https.ApiError) error {
	return &ResponseHttp401Error{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseHttp401Error.
func (r ResponseHttp401Error) Error() string {
	return fmt.Sprintf("ResponseHttp401Error occured: %v", r.Message)
}

// ResponseHttp403Error is a custom error.
type ResponseHttp403Error struct {
	https.ApiError
	Detail *string `json:"detail,omitempty"`
}

// NewResponseHttp403Error is a constructor for ResponseHttp403Error.
// It creates and returns a pointer to a new ResponseHttp403Error instance with the given statusCode and body.
func NewResponseHttp403Error(apiError https.ApiError) error {
	return &ResponseHttp403Error{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseHttp403Error.
func (r ResponseHttp403Error) Error() string {
	return fmt.Sprintf("ResponseHttp403Error occured: %v", r.Message)
}

// ResponseHttp429Error is a custom error.
type ResponseHttp429Error struct {
	https.ApiError
	Detail *string `json:"detail,omitempty"`
}

// NewResponseHttp429Error is a constructor for ResponseHttp429Error.
// It creates and returns a pointer to a new ResponseHttp429Error instance with the given statusCode and body.
func NewResponseHttp429Error(apiError https.ApiError) error {
	return &ResponseHttp429Error{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseHttp429Error.
func (r ResponseHttp429Error) Error() string {
	return fmt.Sprintf("ResponseHttp429Error occured: %v", r.Message)
}

// ResponseInventoryError is a custom error.
type ResponseInventoryError struct {
	https.ApiError
	Added               []string                                           `json:"added,omitempty"`
	Duplicated          []string                                           `json:"duplicated,omitempty"`
	MError              []string                                           `json:"error,omitempty"`
	InventoryAdded      []models.ResponseInventoryInventoryAddedItems      `json:"inventory_added,omitempty"`
	InventoryDuplicated []models.ResponseInventoryInventoryDuplicatedItems `json:"inventory_duplicated,omitempty"`
	Reason              []string                                           `json:"reason,omitempty"`
}

// NewResponseInventoryError is a constructor for ResponseInventoryError.
// It creates and returns a pointer to a new ResponseInventoryError instance with the given statusCode and body.
func NewResponseInventoryError(apiError https.ApiError) error {
	return &ResponseInventoryError{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseInventoryError.
func (r ResponseInventoryError) Error() string {
	return fmt.Sprintf("ResponseInventoryError occured: %v", r.Message)
}
