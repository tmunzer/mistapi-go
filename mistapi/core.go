// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/logger"
)

// HttpConfiguration holds the configuration options for the client.
type HttpConfiguration = https.HttpConfiguration

// RetryConfiguration holds the configuration options for the client.
type RetryConfiguration = https.RetryConfiguration

// LoggerInterface holds the configuration options for the client.
type LoggerInterface = logger.LoggerInterface

// SdkLoggerInterface holds the configuration options for the client.
type SdkLoggerInterface = logger.SdkLoggerInterface

// LoggerConfiguration holds the configuration options for the client.
type LoggerConfiguration = logger.LoggerConfiguration

// Level holds the configuration options for the client.
type Level = logger.Level

// LoggerOptions holds the configuration options for the client.
type LoggerOptions = logger.LoggerOptions

// RetryConfigurationOptions is a function type that modifies the RetryConfiguration.
type RetryConfigurationOptions = https.RetryConfigurationOptions

// HttpConfigurationOptions is a function type that modifies the HttpConfiguration.
type HttpConfigurationOptions = https.HttpConfigurationOptions

// NewRetryConfiguration creates a new RetryConfiguration.
var NewRetryConfiguration = https.NewRetryConfiguration

// NewHttpConfiguration creates a new HttpConfiguration.
var NewHttpConfiguration = https.NewHttpConfiguration

// NewAndAuth creates a new AndAuth.
var NewAndAuth = https.NewAndAuth

// NewOrAuth creates a new OrAuth.
var NewOrAuth = https.NewOrAuth

// NewAuth creates a new Auth.
var NewAuth = https.NewAuth

// NewSdkLogger creates a new SdkLogger.
var NewSdkLogger = logger.NewSdkLogger

// NewLoggerConfiguration creates a new LoggerConfiguration.
var NewLoggerConfiguration = logger.NewLoggerConfiguration

// WithLevel creates a new hLevel.
var WithLevel = logger.WithLevel

// WithLogger creates a new hLogger.
var WithLogger = logger.WithLogger

// WithMaskSensitiveHeaders creates a new hMaskSensitiveHeaders.
var WithMaskSensitiveHeaders = logger.WithMaskSensitiveHeaders

// WithRequestConfiguration creates a new hRequestConfiguration.
var WithRequestConfiguration = logger.WithRequestConfiguration

// WithResponseConfiguration creates a new hResponseConfiguration.
var WithResponseConfiguration = logger.WithResponseConfiguration

// WithRequestBody creates a new hRequestBody.
var WithRequestBody = logger.WithRequestBody

// WithRequestHeaders creates a new hRequestHeaders.
var WithRequestHeaders = logger.WithRequestHeaders

// WithExcludeRequestHeaders creates a new hExcludeRequestHeaders.
var WithExcludeRequestHeaders = logger.WithExcludeRequestHeaders

// WithIncludeRequestHeaders creates a new hIncludeRequestHeaders.
var WithIncludeRequestHeaders = logger.WithIncludeRequestHeaders

// WithWhitelistRequestHeaders creates a new hWhitelistRequestHeaders.
var WithWhitelistRequestHeaders = logger.WithWhitelistRequestHeaders

// WithIncludeQueryInPath creates a new hIncludeQueryInPath.
var WithIncludeQueryInPath = logger.WithIncludeQueryInPath

// WithResponseBody creates a new hResponseBody.
var WithResponseBody = logger.WithResponseBody

// WithResponseHeaders creates a new hResponseHeaders.
var WithResponseHeaders = logger.WithResponseHeaders

// WithExcludeResponseHeaders creates a new hExcludeResponseHeaders.
var WithExcludeResponseHeaders = logger.WithExcludeResponseHeaders

// WithIncludeResponseHeaders creates a new hIncludeResponseHeaders.
var WithIncludeResponseHeaders = logger.WithIncludeResponseHeaders

// WithWhitelistResponseHeaders creates a new hWhitelistResponseHeaders.
var WithWhitelistResponseHeaders = logger.WithWhitelistResponseHeaders

// WithMaxRetryAttempts sets the MaxRetryAttempts.
var WithMaxRetryAttempts = https.WithMaxRetryAttempts

// WithRetryOnTimeout sets the RetryOnTimeout.
var WithRetryOnTimeout = https.WithRetryOnTimeout

// WithRetryInterval sets the RetryInterval.
var WithRetryInterval = https.WithRetryInterval

// WithMaximumRetryWaitTime sets the MaximumRetryWaitTime.
var WithMaximumRetryWaitTime = https.WithMaximumRetryWaitTime

// WithBackoffFactor sets the BackoffFactor.
var WithBackoffFactor = https.WithBackoffFactor

// WithHttpStatusCodesToRetry sets the HttpStatusCodesToRetry.
var WithHttpStatusCodesToRetry = https.WithHttpStatusCodesToRetry

// WithHttpMethodsToRetry sets the HttpMethodsToRetry.
var WithHttpMethodsToRetry = https.WithHttpMethodsToRetry

// WithTimeout sets the Timeout.
var WithTimeout = https.WithTimeout

// WithTransport sets the Transport.
var WithTransport = https.WithTransport

// WithRetryConfiguration sets the RetryConfiguration.
var WithRetryConfiguration = https.WithRetryConfiguration

// FromFastHttpRequest creates http.Request using provided values and Closures.
var FromFastHttpRequest = https.FromFastHttpRequest
