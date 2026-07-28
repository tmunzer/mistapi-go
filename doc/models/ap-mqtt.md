
# Ap Mqtt

MQTT publishing configuration for an AP. Use `mqtt_topic` on individual AssetFilter entries to specify which MQTT topic each matching BLE advertisement is forwarded to. Only AssetFilters with `mqtt_topic` set are used; disabled filters and filters without `mqtt_topic` are skipped. Set `default_topic` to publish advertisements that match no AssetFilter to a catch-all topic, allowing MQTT to be used without configuring any AssetFilter.

## Structure

`ApMqtt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BrokerHost` | `*string` | Optional | MQTT broker hostname or IP address; required when `enabled` is `true` |
| `BrokerPort` | `*int` | Optional | MQTT broker port; defaults to `1883` for `tcp` and `8883` for `ssl` |
| `BrokerProto` | [`*models.ApMqttBrokerProtoEnum`](../../doc/models/ap-mqtt-broker-proto-enum.md) | Optional | MQTT broker transport protocol. enum: `ssl`, `tcp`<br><br>**Default**: `"tcp"` |
| `DefaultTopic` | `*string` | Optional | Optional catch-all MQTT topic; BLE advertisements matching no AssetFilter are published here |
| `Enabled` | `*bool` | Optional | Whether to enable MQTT publishing<br><br>**Default**: `false` |
| `Format` | [`*models.ApMqttFormatEnum`](../../doc/models/ap-mqtt-format-enum.md) | Optional | Payload format for MQTT published messages. enum: `json`, `raw`<br><br>**Default**: `"raw"` |
| `Password` | `*string` | Optional | Optional MQTT password; masked in GET responses |
| `Username` | `*string` | Optional | Optional MQTT username |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apMqtt := models.ApMqtt{
        BrokerHost:           models.ToPointer("broker_host4"),
        BrokerPort:           models.ToPointer(188),
        BrokerProto:          models.ToPointer(models.ApMqttBrokerProtoEnum_TCP),
        DefaultTopic:         models.ToPointer("mist/asset/raw"),
        Enabled:              models.ToPointer(false),
        Format:               models.ToPointer(models.ApMqttFormatEnum_RAW),
    }

}
```

