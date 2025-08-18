// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"os"
)

// ConfigurationOptions represents a function type that can be used to apply options to the Configuration struct.
type ConfigurationOptions func(*Configuration)

// Configuration holds configuration settings.
type Configuration struct {
	accept               string
	environment          Environment
	httpConfiguration    HttpConfiguration
	apiTokenCredentials  ApiTokenCredentials
	basicAuthCredentials BasicAuthCredentials
	csrfTokenCredentials CsrfTokenCredentials
	loggerConfiguration  LoggerConfiguration
}

// newConfiguration creates a new Configuration with the provided options.
func newConfiguration(options ...ConfigurationOptions) Configuration {
	config := Configuration{}

	for _, option := range options {
		option(&config)
	}
	return config
}

// cloneWithOptions provides configuration with the provided options.
func (config Configuration) cloneWithOptions(options ...ConfigurationOptions) Configuration {
	for _, option := range options {
		option(&config)
	}
	return config
}

// WithAccept is an option that sets the Accept in the Configuration.
func WithAccept(accept string) ConfigurationOptions {
	return func(c *Configuration) {
		c.accept = accept
	}
}

// WithEnvironment is an option that sets the Environment in the Configuration.
func WithEnvironment(environment Environment) ConfigurationOptions {
	return func(c *Configuration) {
		c.environment = environment
	}
}

// WithHttpConfiguration is an option that sets the HttpConfiguration in the Configuration.
func WithHttpConfiguration(httpConfiguration HttpConfiguration) ConfigurationOptions {
	return func(c *Configuration) {
		c.httpConfiguration = httpConfiguration
	}
}

// WithApiTokenCredentials is an option that sets the ApiTokenCredentials in the Configuration.
func WithApiTokenCredentials(apiTokenCredentials ApiTokenCredentials) ConfigurationOptions {
	return func(c *Configuration) {
		c.apiTokenCredentials = apiTokenCredentials
	}
}

// WithBasicAuthCredentials is an option that sets the BasicAuthCredentials in the Configuration.
func WithBasicAuthCredentials(basicAuthCredentials BasicAuthCredentials) ConfigurationOptions {
	return func(c *Configuration) {
		c.basicAuthCredentials = basicAuthCredentials
	}
}

// WithCsrfTokenCredentials is an option that sets the CsrfTokenCredentials in the Configuration.
func WithCsrfTokenCredentials(csrfTokenCredentials CsrfTokenCredentials) ConfigurationOptions {
	return func(c *Configuration) {
		c.csrfTokenCredentials = csrfTokenCredentials
	}
}

// WithLoggerConfiguration is an option that sets the LoggerConfiguration in the Configuration.
func WithLoggerConfiguration(options ...LoggerOptions) ConfigurationOptions {
	return func(c *Configuration) {
		c.loggerConfiguration = NewLoggerConfiguration(options...)
	}
}

// Accept returns the accept from the Configuration.
func (c Configuration) Accept() string {
	return c.accept
}

// Environment returns the environment from the Configuration.
func (c Configuration) Environment() Environment {
	return c.environment
}

// HttpConfiguration returns the httpConfiguration from the Configuration.
func (c Configuration) HttpConfiguration() HttpConfiguration {
	return c.httpConfiguration
}

// ApiTokenCredentials returns the apiTokenCredentials from the Configuration.
func (c Configuration) ApiTokenCredentials() ApiTokenCredentials {
	return c.apiTokenCredentials
}

// BasicAuthCredentials returns the basicAuthCredentials from the Configuration.
func (c Configuration) BasicAuthCredentials() BasicAuthCredentials {
	return c.basicAuthCredentials
}

// CsrfTokenCredentials returns the csrfTokenCredentials from the Configuration.
func (c Configuration) CsrfTokenCredentials() CsrfTokenCredentials {
	return c.csrfTokenCredentials
}

// CreateConfigurationFromEnvironment creates a new Configuration with default settings.
// It also configures various Configuration options.
func CreateConfigurationFromEnvironment(options ...ConfigurationOptions) Configuration {
	config := DefaultConfiguration()

	accept := os.Getenv("MISTAPI_ACCEPT")
	if accept != "" {
		config.accept = accept
	}
	environment := os.Getenv("MISTAPI_ENVIRONMENT")
	if environment != "" {
		config.environment = Environment(environment)
	}
	authorization := os.Getenv("MISTAPI_AUTHORIZATION")
	if authorization != "" {
		config.apiTokenCredentials.authorization = authorization
	}
	username := os.Getenv("MISTAPI_USERNAME")
	if username != "" {
		config.basicAuthCredentials.username = username
	}
	password := os.Getenv("MISTAPI_PASSWORD")
	if password != "" {
		config.basicAuthCredentials.password = password
	}
	xCSRFToken := os.Getenv("MISTAPI_X_CSRF_TOKEN")
	if xCSRFToken != "" {
		config.csrfTokenCredentials.xCSRFToken = xCSRFToken
	}
	for _, option := range options {
		option(&config)
	}
	return config
}

// Server represents available servers.
type Server string

const (
	APIHOST Server = "API Host"
)

// Environment represents available environments.
type Environment string

const (
	MIST_GLOBAL_01 Environment = "Mist Global 01"
	MIST_GLOBAL_02 Environment = "Mist Global 02"
	MIST_GLOBAL_03 Environment = "Mist Global 03"
	MIST_GLOBAL_04 Environment = "Mist Global 04"
	MIST_GLOBAL_05 Environment = "Mist Global 05"
	MIST_EMEA_01   Environment = "Mist EMEA 01"
	MIST_EMEA_02   Environment = "Mist EMEA 02"
	MIST_EMEA_03   Environment = "Mist EMEA 03"
	MIST_EMEA_04   Environment = "Mist EMEA 04"
	MIST_APAC_01   Environment = "Mist APAC 01"
	AWS_STAGING    Environment = "AWS Staging"
	GOV_CLOUD      Environment = "GOV CLOUD"
	MIST_APAC_03   Environment = "Mist APAC 03"
)

// CreateRetryConfiguration creates a new RetryConfiguration with the provided options.
func CreateRetryConfiguration(options ...RetryConfigurationOptions) RetryConfiguration {
	config := DefaultRetryConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}

// CreateHttpConfiguration creates a new HttpConfiguration with the provided options.
func CreateHttpConfiguration(options ...HttpConfigurationOptions) HttpConfiguration {
	config := DefaultHttpConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}

// CreateConfiguration creates a new Configuration with the provided options.
func CreateConfiguration(options ...ConfigurationOptions) Configuration {
	config := DefaultConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}
