// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "net/http"
)

// DefaultRetryConfiguration returns the default RetryConfiguration for HTTP requests.
// It also configures various retry options.
func DefaultRetryConfiguration() RetryConfiguration {
    return NewRetryConfiguration(
        WithMaxRetryAttempts(0),
        WithRetryOnTimeout(true),
        WithRetryInterval(1),
        WithMaximumRetryWaitTime(0),
        WithBackoffFactor(2),
        WithHttpStatusCodesToRetry([]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524, 408, 413, 429, 500, 502, 503, 504, 521, 522, 524}),
        WithHttpMethodsToRetry([]string{"GET", "PUT", "GET", "PUT"}),
    )
}

// DefaultHttpConfiguration returns the default HttpConfiguration for HTTP requests.
// It also configures various HttpConfiguration options.
func DefaultHttpConfiguration() HttpConfiguration {
    return NewHttpConfiguration(
        WithTimeout(0),
        WithTransport(http.DefaultTransport),
        WithRetryConfiguration(DefaultRetryConfiguration()),
    )
}

// DefaultConfiguration returns the default Configuration.
func DefaultConfiguration() Configuration {
    return newConfiguration(
        WithAccept("application/json, application/vnd.api+json"),
        WithEnvironment(MIST_GLOBAL_01),
        WithHttpConfiguration(DefaultHttpConfiguration()),
    )
}
