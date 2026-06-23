
# Wlan App Qos

APP qos wlan settings

## Structure

`WlanAppQos`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | [`map[string]models.WlanAppQosAppsProperties`](../../doc/models/wlan-app-qos-apps-properties.md) | Optional | Map of application keys to QoS rewrite settings |
| `Enabled` | `*bool` | Optional | Whether application QoS rewrite rules are enabled for this WLAN<br><br>**Default**: `false` |
| `Others` | [`[]models.WlanAppQosOthersItem`](../../doc/models/wlan-app-qos-others-item.md) | Optional | Custom application QoS rules for traffic matched by subnet, port, and protocol<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanAppQos := models.WlanAppQos{
        Apps:                 map[string]models.WlanAppQosAppsProperties{
            "skype-business-video": models.WlanAppQosAppsProperties{
                Dscp:                 models.ToPointer(models.DscpContainer.FromNumber(32)),
                DstSubnet:            models.ToPointer("10.2.0.0/16"),
                SrcSubnet:            models.ToPointer("10.2.0.0/16"),
            },
        },
        Enabled:              models.ToPointer(false),
        Others:               []models.WlanAppQosOthersItem{
            models.WlanAppQosOthersItem{
                Dscp:                 models.ToPointer(models.DscpContainer.FromString("String7")),
                DstSubnet:            models.ToPointer("dst_subnet2"),
                PortRanges:           models.ToPointer("port_ranges6"),
                Protocol:             models.ToPointer("protocol2"),
                SrcSubnet:            models.ToPointer("src_subnet0"),
            },
            models.WlanAppQosOthersItem{
                Dscp:                 models.ToPointer(models.DscpContainer.FromString("String7")),
                DstSubnet:            models.ToPointer("dst_subnet2"),
                PortRanges:           models.ToPointer("port_ranges6"),
                Protocol:             models.ToPointer("protocol2"),
                SrcSubnet:            models.ToPointer("src_subnet0"),
            },
            models.WlanAppQosOthersItem{
                Dscp:                 models.ToPointer(models.DscpContainer.FromString("String7")),
                DstSubnet:            models.ToPointer("dst_subnet2"),
                PortRanges:           models.ToPointer("port_ranges6"),
                Protocol:             models.ToPointer("protocol2"),
                SrcSubnet:            models.ToPointer("src_subnet0"),
            },
        },
    }

}
```

