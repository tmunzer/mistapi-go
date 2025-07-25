
# HttpConfiguration

The following parameters are configurable for the HttpConfiguration:

## Properties

| Name | Type | Description | Setter | Getter |
|  --- | --- | --- | --- | --- |
| timeout | `float64` | Timeout in milliseconds.<br>*Default*: `0` | `WithTimeout` | `Timeout()` |
| transport | `httpRoundTripper` | Establishes network connection and caches them for reuse.<br>*Default*: `http.DefaultTransport` | `WithTransport` | `Transport()` |
| retryConfiguration | [`mistapiRetryConfiguration`](../doc/retry-configuration.md) | Configurations to retry requests.<br>*Default*: `mistapi.DefaultRetryConfiguration()` | `WithRetryConfiguration` | `RetryConfiguration()` |

The httpConfiguration can be initialized as follows:

```go
package main

import (
    "mistapi"
    "net/http"
)

func main() {
    httpConfiguration := mistapi.CreateHttpConfiguration(
        mistapi.WithTimeout(0),
        mistapi.WithTransport(http.DefaultTransport),
        mistapi.WithRetryConfiguration(mistapi.DefaultRetryConfiguration()),
    )
}
```

