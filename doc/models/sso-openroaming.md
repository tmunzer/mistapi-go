
# Sso Openroaming

Deprecated. OpenRoaming configuration is now expressed as top-level fields on the SSO object: `openroaming_ssids`, `openroaming_wba_client_cert`, and `openroaming_wba_client_key`.

## Structure

`SsoOpenroaming`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ssids` | `[]string` | Optional | Network SSID names enabled for OpenRoaming SSO |
| `WbaCert` | `*string` | Optional | Deprecated. Use `openroaming_wba_client_cert` instead. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ssoOpenroaming := models.SsoOpenroaming{
        Ssids:                []string{
            "ssid_name1",
            "ssid_name2",
        },
        WbaCert:              models.ToPointer("wba_cert4"),
    }

}
```

