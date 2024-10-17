package mistapi

import (
	"errors"
	"github.com/apimatic/go-core-runtime/https"
	"net/http"
)

// CsrfTokenCredentials represents the credentials required for `csrfToken` authentication.
type CsrfTokenCredentials struct {
	xCSRFToken string
}

// NewCsrfTokenCredentials creates a new instance of CsrfTokenCredentials with provided parameters.
func NewCsrfTokenCredentials(xCSRFToken string) CsrfTokenCredentials {
	return CsrfTokenCredentials{
		xCSRFToken: xCSRFToken,
	}
}

// WithXCSRFToken sets xCSRFToken in CsrfTokenCredentials.
func (c CsrfTokenCredentials) WithXCSRFToken(xCSRFToken string) CsrfTokenCredentials {
	c.xCSRFToken = xCSRFToken
	return c
}

// XCSRFToken returns the xCSRFToken associated with the CsrfTokenCredentials.
func (c CsrfTokenCredentials) XCSRFToken() string {
	return c.xCSRFToken
}

// Validate function returns validation error associated with the CsrfTokenCredentials.
func (c CsrfTokenCredentials) Validate() error {
	if c.xCSRFToken == "" {
		return errors.New("csrfToken : missing auth credentials -> XCSRFToken.")
	}
	return nil
}

// Authenticator function returns HttpInterceptor function that provides authentication for API calls.
func (c CsrfTokenCredentials) Authenticator() https.HttpInterceptor {
	return func(req *http.Request,
		next https.HttpCallExecutor,
	) https.HttpContext {
		req.Header.Set("X-CSRFToken", c.XCSRFToken())
		return next(req)
	}
}
