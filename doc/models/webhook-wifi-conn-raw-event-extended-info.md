
# Webhook Wifi Conn Raw Event Extended Info

Extended telemetry details decoded from a connected Wi-Fi packet

## Structure

`WebhookWifiConnRawEventExtendedInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FrameCtrl` | `*int` | Optional | IEEE 802.11 frame control value from the telemetry packet |
| `Payload` | `*string` | Optional | Raw telemetry payload carried by the connected Wi-Fi packet |
| `SequenceCtrl` | `*int` | Optional | IEEE 802.11 sequence control value from the telemetry packet |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    webhookWifiConnRawEventExtendedInfo := models.WebhookWifiConnRawEventExtendedInfo{
        FrameCtrl:            models.ToPointer(198),
        Payload:              models.ToPointer("payload2"),
        SequenceCtrl:         models.ToPointer(248),
    }

}
```

