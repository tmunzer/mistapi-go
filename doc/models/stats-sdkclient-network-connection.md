
# Stats Sdkclient Network Connection

Current network connection details reported for an SDK client

## Structure

`StatsSdkclientNetworkConnection`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Current network connection MAC address reported for the SDK client |
| `Rssi` | `float64` | Required | Received signal strength indicator for the SDK client's current network connection, in dBm |
| `SignalLevel` | `float64` | Required | Numeric signal quality level reported with the SDK client's current network connection |
| `Type` | `string` | Required | Network connection type reported for the SDK client, such as WiFi |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsSdkclientNetworkConnection := models.StatsSdkclientNetworkConnection{
        Mac:                  "mac0",
        Rssi:                 float64(231.48),
        SignalLevel:          float64(43.5),
        Type:                 "type4",
    }

}
```

