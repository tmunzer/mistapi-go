
# RequestLoggerConfiguration

Represents the logging configurations for API requests.

## Methods

| Name | Type | Description | Setter |
|  --- | --- | --- | --- |
| includeQueryInPath | `bool` | Toggles the inclusion of query parameters in the logged request path. **Default : `false`** | `WithIncludeQueryInPath` |
| body | `bool` | Toggles the logging of the request body. **Default : `false`** | `WithRequestBody` |
| headers | `bool` | Toggles the logging of the request headers. **Default : `false`** | `WithRequestHeaders` |
| includeHeaders | `[]string` | Includes only specified request headers (provided as variadic string arguments) in the log output. **Default : `[]string{}`** | `WithIncludeRequestHeaders(...string)` |
| excludeHeaders | `[]string` | Excludes specified request headers (provided as variadic string arguments) from the log output. **Default : `[]string{}`** | `WithExcludeRequestHeaders(...string)` |
| whitelistHeaders | `[]string` | Logs specified request headers (provided as variadic string arguments) without masking, revealing their actual values. **Default : `[]string{}`** | `WithWhitelistRequestHeaders(...string)` |

