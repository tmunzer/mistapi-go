
# LoggerConfiguration

Represents the logging configurations for API calls.

## Methods

| Name | Type | Description | Setter |
|  --- | --- | --- | --- |
| logger | `LoggerInterface` | Takes in your custom implementation of the LoggerInterface. **Default: `ConsoleLogger`** | `WithLogger` |
| level | `Level` | Defines the log message severity (e.g., DEBUG, INFO, etc). **Default : `Level_INFO`** | `WithLevel` |
| maskSensitiveHeaders | `bool` | Toggles the global setting to mask sensitive HTTP headers in both requests and responses before logging, safeguarding confidential data. **Default : `true`** | `WithMaskSensitiveHeaders` |
| request | [`RequestLoggerConfiguration`](../doc/request-logger-configuration.md) | The logging configurations for an API request. | `WithRequestConfiguration` |
| response | [`ResponseLoggerConfiguration`](../doc/response-logger-configuration.md) | The logging configurations for an API response. | `WithResponseConfiguration` |

## Usage Example

### Client Initialization with Custom Logger

In order to provide custom logger, any implementation of the `LoggerInterface` can be used so that you can override the `Log` behavior and provide its instance directly in the SDK client initialization.

The following example uses [`Logrus`](https://pkg.go.dev/github.com/sirupsen/logrus) implementation of `LoggerInterface` for MistAPIClient initialization.

```go
import (
	log "github.com/sirupsen/logrus"
	"os"
)

type Logrus struct{}

func (c Logrus) Log(level log.Level, message string, params map[string]any) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	log.WithFields(params).Debug(message)
}
config := CreateConfigurationFromEnvironment(
	WithLoggerConfiguration(
		WithLevel("debug"),
		WithLogger(Logrus{}),
		WithRequestConfiguration(
			WithRequestBody(true),
		),
		WithResponseConfiguration(
			WithResponseHeaders(true),
		),
	),
)
```

