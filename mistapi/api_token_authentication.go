// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"errors"
	"github.com/apimatic/go-core-runtime/https"
	"net/http"
)

// ApiTokenCredentials represents the credentials required for `apiToken` authentication.
type ApiTokenCredentials struct {
	authorization string
}

// NewApiTokenCredentials creates a new instance of ApiTokenCredentials with provided parameters.
func NewApiTokenCredentials(authorization string) ApiTokenCredentials {
	return ApiTokenCredentials{
		authorization: authorization,
	}
}

// WithAuthorization sets authorization in ApiTokenCredentials.
func (a ApiTokenCredentials) WithAuthorization(authorization string) ApiTokenCredentials {
	a.authorization = authorization
	return a
}

// Authorization returns the authorization associated with the ApiTokenCredentials.
func (a ApiTokenCredentials) Authorization() string {
	return a.authorization
}

// Validate function returns validation error associated with the ApiTokenCredentials.
func (a ApiTokenCredentials) Validate() error {
	if a.authorization == "" {
		return errors.New("apiToken : missing auth credentials -> Authorization.")
	}
	return nil
}

// Authenticator function returns HttpInterceptor function that provides authentication for API calls.
func (a ApiTokenCredentials) Authenticator() https.HttpInterceptor {
	return func(req *http.Request,
		next https.HttpCallExecutor,
	) https.HttpContext {
		req.Header.Set("Authorization", a.Authorization())
		return next(req)
	}
}
