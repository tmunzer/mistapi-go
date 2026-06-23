
# Radius Acct Server

RADIUS accounting server settings

## Structure

`RadiusAcctServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `string` | Required | Address or hostname of the RADIUS accounting server |
| `KeywrapEnabled` | `*bool` | Optional | Whether RADIUS keywrap is enabled for messages sent to this accounting server |
| `KeywrapFormat` | [`*models.RadiusKeywrapFormatEnum`](../../doc/models/radius-keywrap-format-enum.md) | Optional | Encoding format for RADIUS keywrap KEK and MACK values. enum: `ascii`, `hex` |
| `KeywrapKek` | `*string` | Optional | RADIUS keywrap key encryption key (KEK) |
| `KeywrapMack` | `*string` | Optional | RADIUS keywrap message authentication code key (MACK) |
| `Port` | [`*models.RadiusAcctPort`](../../doc/models/containers/radius-acct-port.md) | Optional | RADIUS Auth Port, value from 1 to 65535, default is 1813 |
| `Secret` | `string` | Required | Shared secret used with this RADIUS accounting server |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    radiusAcctServer := models.RadiusAcctServer{
        Host:                 "1.2.3.4",
        KeywrapEnabled:       models.ToPointer(false),
        KeywrapFormat:        models.ToPointer(models.RadiusKeywrapFormatEnum_ASCII),
        KeywrapKek:           models.ToPointer("1122334455"),
        KeywrapMack:          models.ToPointer("1122334455"),
        Port:                 models.ToPointer(models.RadiusAcctPortContainer.FromNumber(36)),
        Secret:               "testing123",
    }

}
```

