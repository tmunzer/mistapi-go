
# Gateway Wan Probe Override Http

HTTP probe settings for a WAN probe override

## Structure

`GatewayWanProbeOverrideHttp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcceptedStatusCodes` | `[]int` | Optional | HTTP response status codes that indicate a successful probe. Defaults to 200 if not specified. |
| `Urls` | `[]string` | Optional | HTTP or HTTPS URLs to probe |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayWanProbeOverrideHttp := models.GatewayWanProbeOverrideHttp{
        AcceptedStatusCodes:  []int{
            204,
        },
        Urls:                 []string{
            "http://www.google.com/generate_204",
            "https://www.google.com/generate_204",
        },
    }

}
```

