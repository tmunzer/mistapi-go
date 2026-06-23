
# Service Policy Ssl Proxy

SRX SSL proxy inspection settings for a service policy

## Structure

`ServicePolicySslProxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CiphersCategory` | [`*models.SslProxyCiphersCategoryEnum`](../../doc/models/ssl-proxy-ciphers-category-enum.md) | Optional | enum: `medium`, `strong`, `weak`<br><br>**Default**: `"strong"` |
| `Enabled` | `*bool` | Optional | Whether SSL proxy inspection is enabled for the service policy<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    servicePolicySslProxy := models.ServicePolicySslProxy{
        CiphersCategory:      models.ToPointer(models.SslProxyCiphersCategoryEnum_STRONG),
        Enabled:              models.ToPointer(false),
    }

}
```

