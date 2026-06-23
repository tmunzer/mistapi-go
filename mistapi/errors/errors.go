// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package errors

import (
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"strings"
)

// ErrorDeleteFailed is a custom error.
// Error response returned when an organization delete request is blocked
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

// String implements the fmt.Stringer interface for ErrorDeleteFailed,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e ErrorDeleteFailed) String() string {
	suffixTrimmed := strings.TrimSuffix(e.ApiError.String(), "]")
	prefixTrimmed := strings.TrimPrefix(suffixTrimmed, "ApiError[")

	return fmt.Sprintf(
		"ErrorDeleteFailed[%v, Detail=%v, OrgId=%v]",
		prefixTrimmed, e.Detail, e.OrgId)
}

// ResponseDetailString is a custom error.
// Response containing a human-readable detail message
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

// String implements the fmt.Stringer interface for ResponseDetailString,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseDetailString) String() string {
	suffixTrimmed := strings.TrimSuffix(r.ApiError.String(), "]")
	prefixTrimmed := strings.TrimPrefix(suffixTrimmed, "ApiError[")

	return fmt.Sprintf(
		"ResponseDetailString[%v, Detail=%v]",
		prefixTrimmed, r.Detail)
}

// ResponseHttp400 is a custom error.
// Standard HTTP 400 bad request error response
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

// String implements the fmt.Stringer interface for ResponseHttp400,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseHttp400) String() string {
	suffixTrimmed := strings.TrimSuffix(r.ApiError.String(), "]")
	prefixTrimmed := strings.TrimPrefix(suffixTrimmed, "ApiError[")

	return fmt.Sprintf(
		"ResponseHttp400[%v, Detail=%v]",
		prefixTrimmed, r.Detail)
}

// ResponseHttp400Webhook is a custom error.
// Webhook-specific HTTP 400 bad request error response
type ResponseHttp400Webhook struct {
	https.ApiError
	Detail *string `json:"detail,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// NewResponseHttp400Webhook is a constructor for ResponseHttp400Webhook.
// It creates and returns a pointer to a new ResponseHttp400Webhook instance with the given statusCode and body.
func NewResponseHttp400Webhook(apiError https.ApiError) error {
	return &ResponseHttp400Webhook{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseHttp400Webhook.
func (r ResponseHttp400Webhook) Error() string {
	return fmt.Sprintf("ResponseHttp400Webhook occured: %v", r.Message)
}

// String implements the fmt.Stringer interface for ResponseHttp400Webhook,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseHttp400Webhook) String() string {
	suffixTrimmed := strings.TrimSuffix(r.ApiError.String(), "]")
	prefixTrimmed := strings.TrimPrefix(suffixTrimmed, "ApiError[")

	return fmt.Sprintf(
		"ResponseHttp400Webhook[%v, Detail=%v, Reason=%v]",
		prefixTrimmed, r.Detail, r.Reason)
}

// ResponseHttp401 is a custom error.
// Standard HTTP 401 authentication error response
type ResponseHttp401 struct {
	https.ApiError
	Detail *string `json:"detail,omitempty"`
}

// NewResponseHttp401 is a constructor for ResponseHttp401.
// It creates and returns a pointer to a new ResponseHttp401 instance with the given statusCode and body.
func NewResponseHttp401(apiError https.ApiError) error {
	return &ResponseHttp401{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseHttp401.
func (r ResponseHttp401) Error() string {
	return fmt.Sprintf("ResponseHttp401 occured: %v", r.Message)
}

// String implements the fmt.Stringer interface for ResponseHttp401,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseHttp401) String() string {
	suffixTrimmed := strings.TrimSuffix(r.ApiError.String(), "]")
	prefixTrimmed := strings.TrimPrefix(suffixTrimmed, "ApiError[")

	return fmt.Sprintf(
		"ResponseHttp401[%v, Detail=%v]",
		prefixTrimmed, r.Detail)
}

// ResponseHttp403 is a custom error.
// Standard HTTP 403 permission error response
type ResponseHttp403 struct {
	https.ApiError
	Detail *string `json:"detail,omitempty"`
}

// NewResponseHttp403 is a constructor for ResponseHttp403.
// It creates and returns a pointer to a new ResponseHttp403 instance with the given statusCode and body.
func NewResponseHttp403(apiError https.ApiError) error {
	return &ResponseHttp403{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseHttp403.
func (r ResponseHttp403) Error() string {
	return fmt.Sprintf("ResponseHttp403 occured: %v", r.Message)
}

// String implements the fmt.Stringer interface for ResponseHttp403,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseHttp403) String() string {
	suffixTrimmed := strings.TrimSuffix(r.ApiError.String(), "]")
	prefixTrimmed := strings.TrimPrefix(suffixTrimmed, "ApiError[")

	return fmt.Sprintf(
		"ResponseHttp403[%v, Detail=%v]",
		prefixTrimmed, r.Detail)
}

// ResponseHttp404 is a custom error.
// Standard HTTP 404 not found error response
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

// String implements the fmt.Stringer interface for ResponseHttp404,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseHttp404) String() string {
	suffixTrimmed := strings.TrimSuffix(r.ApiError.String(), "]")
	prefixTrimmed := strings.TrimPrefix(suffixTrimmed, "ApiError[")

	return fmt.Sprintf(
		"ResponseHttp404[%v, Id=%v]",
		prefixTrimmed, r.Id)
}

// ResponseHttp429 is a custom error.
// Standard HTTP 429 rate limit error response
type ResponseHttp429 struct {
	https.ApiError
	Detail *string `json:"detail,omitempty"`
}

// NewResponseHttp429 is a constructor for ResponseHttp429.
// It creates and returns a pointer to a new ResponseHttp429 instance with the given statusCode and body.
func NewResponseHttp429(apiError https.ApiError) error {
	return &ResponseHttp429{
		ApiError: apiError,
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ResponseHttp429.
func (r ResponseHttp429) Error() string {
	return fmt.Sprintf("ResponseHttp429 occured: %v", r.Message)
}

// String implements the fmt.Stringer interface for ResponseHttp429,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseHttp429) String() string {
	suffixTrimmed := strings.TrimSuffix(r.ApiError.String(), "]")
	prefixTrimmed := strings.TrimPrefix(suffixTrimmed, "ApiError[")

	return fmt.Sprintf(
		"ResponseHttp429[%v, Detail=%v]",
		prefixTrimmed, r.Detail)
}

// ResponseLoginFailure is a custom error.
// Login failure response returned when authentication cannot continue locally
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

// String implements the fmt.Stringer interface for ResponseLoginFailure,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseLoginFailure) String() string {
	suffixTrimmed := strings.TrimSuffix(r.ApiError.String(), "]")
	prefixTrimmed := strings.TrimPrefix(suffixTrimmed, "ApiError[")

	return fmt.Sprintf(
		"ResponseLoginFailure[%v, Detail=%v, ForwardUrl=%v]",
		prefixTrimmed, r.Detail, r.ForwardUrl)
}

// ResponseSelfOauthLinkFailure is a custom error.
// OAuth2 account-linking failure response
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

// String implements the fmt.Stringer interface for ResponseSelfOauthLinkFailure,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSelfOauthLinkFailure) String() string {
	suffixTrimmed := strings.TrimSuffix(r.ApiError.String(), "]")
	prefixTrimmed := strings.TrimPrefix(suffixTrimmed, "ApiError[")

	return fmt.Sprintf(
		"ResponseSelfOauthLinkFailure[%v, MError=%v, ErrorDescription=%v]",
		prefixTrimmed, r.MError, r.ErrorDescription)
}

// ResponseInventoryError is a custom error.
// Result of adding device claim codes to organization inventory
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

// String implements the fmt.Stringer interface for ResponseInventoryError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseInventoryError) String() string {
	suffixTrimmed := strings.TrimSuffix(r.ApiError.String(), "]")
	prefixTrimmed := strings.TrimPrefix(suffixTrimmed, "ApiError[")

	return fmt.Sprintf(
		"ResponseInventoryError[%v, Added=%v, Duplicated=%v, MError=%v, InventoryAdded=%v, InventoryDuplicated=%v, Reason=%v]",
		prefixTrimmed, r.Added, r.Duplicated, r.MError, r.InventoryAdded, r.InventoryDuplicated, r.Reason)
}
