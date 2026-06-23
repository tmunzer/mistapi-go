
# Mxcluster Radsec Auth Server

RadSec authentication server settings for a Mist Edge cluster

## Structure

`MxclusterRadsecAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | IP / hostname of RADIUS server |
| `InbandStatusCheck` | `*bool` | Optional | Whether to enable inband status check<br><br>**Default**: `false` |
| `InbandStatusInterval` | `*int` | Optional | Inband status interval, in seconds<br><br>**Default**: `300`<br><br>**Constraints**: `>= 0` |
| `KeywrapEnabled` | `*bool` | Optional | If used for Mist APs, enable keywrap algorithm. Default is false |
| `KeywrapFormat` | [`models.Optional[models.MxclusterRadAuthServerKeywrapFormatEnum]`](../../doc/models/mxcluster-rad-auth-server-keywrap-format-enum.md) | Optional | if used for Mist APs. enum: `ascii`, `hex`<br><br>**Default**: `"ascii"` |
| `KeywrapKek` | `*string` | Optional | If used for Mist APs, encryption key |
| `KeywrapMack` | `*string` | Optional | If used for Mist APs, Message Authentication Code Key |
| `Port` | `*int` | Optional | Auth port of RADIUS server<br><br>**Default**: `1812` |
| `Retry` | `*int` | Optional | Number of authentication request retries before failing over<br><br>**Default**: `2` |
| `Secret` | `*string` | Optional | Shared secret used with this RADIUS authentication server |
| `Ssids` | `[]string` | Optional | List of ssids that will use this server if match_ssid is true and match is found |
| `Timeout` | `*int` | Optional | Authentication request timeout, in seconds<br><br>**Default**: `5` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxclusterRadsecAuthServer := models.MxclusterRadsecAuthServer{
        Host:                 models.ToPointer("host2"),
        InbandStatusCheck:    models.ToPointer(false),
        InbandStatusInterval: models.ToPointer(300),
        KeywrapEnabled:       models.ToPointer(false),
        KeywrapFormat:        models.NewOptional(models.ToPointer(models.MxclusterRadAuthServerKeywrapFormatEnum_ASCII)),
        Port:                 models.ToPointer(1812),
        Retry:                models.ToPointer(2),
        Timeout:              models.ToPointer(5),
    }

}
```

