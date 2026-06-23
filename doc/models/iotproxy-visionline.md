
# Iotproxy Visionline

Visionline integration settings for IoT proxy

## Structure

`IotproxyVisionline`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessId` | `*string` | Optional | Access ID for the Visionline service |
| `Cacerts` | `[]string` | Optional | PEM-encoded CA certificates used to verify the Visionline collector's TLS certificate. Required when the collector uses a self-signed certificate |
| `Enabled` | `*bool` | Optional | Whether the Visionline integration is enabled<br><br>**Default**: `false` |
| `Host` | `*string` | Optional | Collector hostname or IP address for Visionline |
| `Password` | `*string` | Optional | Visionline service password used by the IoT proxy |
| `Port` | `*int` | Optional | TCP port of the Visionline collector<br><br>**Default**: `443` |
| `Username` | `*string` | Optional | Visionline service username used by the IoT proxy |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    iotproxyVisionline := models.IotproxyVisionline{
        AccessId:             models.ToPointer("790e6c1790e6c18541d"),
        Cacerts:              []string{
            "cacerts4",
            "cacerts3",
        },
        Enabled:              models.ToPointer(false),
        Host:                 models.ToPointer("visionline_collector1.local"),
        Password:             models.ToPointer("password0"),
        Port:                 models.ToPointer(443),
        Username:             models.ToPointer("card_administrator"),
    }

}
```

