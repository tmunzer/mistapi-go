
# ResponseLoggerConfiguration

Represents the logging configurations for API responses.

## Methods

| Name | Type | Description | Setter |
|  --- | --- | --- | --- |
| body | `bool` | Toggles the logging of the request body. **Default : `false`** | `WithResponseBody` |
| headers | `bool` | Toggles the logging of the request headers. **Default : `false`** | `WithResponseHeaders` |
| includeHeaders | `[]string` | Includes only specified request headers (provided as variadic string arguments) in the log output. **Default : `[]string{}`** | `WithIncludeResponseHeaders(...string)` |
| excludeHeaders | `[]string` | Excludes specified request headers (provided as variadic string arguments) from the log output. **Default : `[]string{}`** | `WithExcludeResponseHeaders(...string)` |
| whitelistHeaders | `[]string` | Logs specified request headers (provided as variadic string arguments) without masking, revealing their actual values. **Default : `[]string{}`** | `WithWhitelistResponseHeaders(...string)` |

