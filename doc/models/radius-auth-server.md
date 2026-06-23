
# Radius Auth Server

RADIUS authentication server settings

## Structure

`RadiusAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `string` | Required | Address or hostname of the RADIUS authentication server |
| `KeywrapEnabled` | `*bool` | Optional | Whether RADIUS keywrap is enabled for messages sent to this authentication server |
| `KeywrapFormat` | [`*models.RadiusKeywrapFormatEnum`](../../doc/models/radius-keywrap-format-enum.md) | Optional | Encoding format for RADIUS keywrap KEK and MACK values. enum: `ascii`, `hex` |
| `KeywrapKek` | `*string` | Optional | RADIUS keywrap key encryption key (KEK) |
| `KeywrapMack` | `*string` | Optional | RADIUS keywrap message authentication code key (MACK) |
| `Port` | [`*models.RadiusAuthPort`](../../doc/models/containers/radius-auth-port.md) | Optional | RADIUS Auth Port, value from 1 to 65535, default is 1812 |
| `RequireMessageAuthenticator` | `*bool` | Optional | Whether to require Message-Authenticator in requests<br><br>**Default**: `false` |
| `Secret` | `string` | Required | Shared secret used with this RADIUS authentication server |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    radiusAuthServer := models.RadiusAuthServer{
        Host:                        "1.2.3.4",
        KeywrapEnabled:              models.ToPointer(false),
        KeywrapFormat:               models.ToPointer(models.RadiusKeywrapFormatEnum_ASCII),
        KeywrapKek:                  models.ToPointer("1122334455"),
        KeywrapMack:                 models.ToPointer("1122334455"),
        Port:                        models.ToPointer(models.RadiusAuthPortContainer.FromNumber(182)),
        RequireMessageAuthenticator: models.ToPointer(false),
        Secret:                      "testing123",
    }

}
```

