
# Wlan Airwatch

AirWatch integration settings for the WLAN

## Structure

`WlanAirwatch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApiKey` | `*string` | Optional | API key used to authenticate to the AirWatch service |
| `ConsoleUrl` | `*string` | Optional | Base console URL of the AirWatch deployment |
| `Enabled` | `*bool` | Optional | Whether AirWatch integration is enabled for the WLAN<br><br>**Default**: `false` |
| `Password` | `*string` | Optional | AirWatch integration account password for this WLAN |
| `Username` | `*string` | Optional | AirWatch integration account username for this WLAN |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanAirwatch := models.WlanAirwatch{
        ApiKey:               models.ToPointer("aHhlbGxvYXNkZmFzZGZhc2Rmc2RmCg==\""),
        ConsoleUrl:           models.ToPointer("https://hs1.airwatchportals.com"),
        Enabled:              models.ToPointer(false),
        Password:             models.ToPointer("user1"),
        Username:             models.ToPointer("test123"),
    }

}
```

