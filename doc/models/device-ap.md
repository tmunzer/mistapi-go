
# Device Ap

Access point configuration and placement data

## Structure

`DeviceAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aeroscout` | [`*models.ApAeroscout`](../../doc/models/ap-aeroscout.md) | Optional | AeroScout location integration settings applied to an AP or AP profile |
| `Airista` | [`*models.ApAirista`](../../doc/models/ap-airista.md) | Optional | Airista RTLS integration settings for an AP |
| `BleConfig` | [`*models.BleConfig`](../../doc/models/ble-config.md) | Optional | Bluetooth Low Energy beacon and asset advertising settings for an AP |
| `Centrak` | [`*models.ApCentrak`](../../doc/models/ap-centrak.md) | Optional | CenTrak integration settings for an AP |
| `ClientBridge` | [`*models.ApClientBridge`](../../doc/models/ap-client-bridge.md) | Optional | AP client bridge mode configuration |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional | Device profile assigned to this access point |
| `DisableEth1` | `*bool` | Optional | Whether to disable eth1 port<br><br>**Default**: `false` |
| `DisableEth2` | `*bool` | Optional | Whether to disable eth2 port<br><br>**Default**: `false` |
| `DisableEth3` | `*bool` | Optional | Whether to disable eth3 port<br><br>**Default**: `false` |
| `DisableModule` | `*bool` | Optional | Whether to disable module port<br><br>**Default**: `false` |
| `EslConfig` | [`*models.ApEslConfig`](../../doc/models/ap-esl-config.md) | Optional | Electronic shelf label integration settings for an AP |
| `FlowControl` | `*bool` | Optional | For some AP models, flow_control can be enabled to address some switch compatibility issue<br><br>**Default**: `false` |
| `ForSite` | `*bool` | Optional, Read-only | Whether the access point configuration is scoped directly to a site |
| `Height` | `*float64` | Optional | Installation height of the AP, in meters |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Image1Url` | `models.Optional[string]` | Optional | First custom image URL associated with the access point |
| `Image2Url` | `models.Optional[string]` | Optional | Second custom image URL associated with the access point |
| `Image3Url` | `models.Optional[string]` | Optional | Third custom image URL associated with the access point |
| `IotConfig` | [`*models.ApIot`](../../doc/models/ap-iot.md) | Optional | Digital and analog IoT port settings applied to an AP or AP profile |
| `IpConfig` | [`*models.ApIpConfig`](../../doc/models/ap-ip-config.md) | Optional | Management IP addressing settings for an access point |
| `LacpConfig` | [`*models.DeviceApLacpConfig`](../../doc/models/device-ap-lacp-config.md) | Optional | LACP settings for supported AP Ethernet uplinks |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | Indicator light settings for an access point |
| `Locked` | `*bool` | Optional | Whether this map is considered locked down |
| `Mac` | `*string` | Optional, Read-only | Access point MAC address used to identify the device |
| `MapId` | `*uuid.UUID` | Optional | Map where the device belongs to |
| `Mesh` | [`*models.ApMesh`](../../doc/models/ap-mesh.md) | Optional | Wireless mesh settings for an access point |
| `Model` | `*string` | Optional, Read-only | Hardware model reported for the access point |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MqttConfig` | [`*models.ApMqtt`](../../doc/models/ap-mqtt.md) | Optional | MQTT broker publishing settings for an AP; use `mqtt_topic` on individual AssetFilter entries to specify which MQTT topic each matching BLE advertisement is forwarded to |
| `Name` | `*string` | Optional | Configured hostname assigned to the access point |
| `Notes` | `*string` | Optional | Any notes about this AP |
| `NtpServers` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Orientation` | `*int` | Optional | AP orientation in degrees from 0 to 359, where 0 is up and 90 is right<br><br>**Constraints**: `>= 0`, `<= 359` |
| `PoePassthrough` | `*bool` | Optional | Whether to enable power out through module port (for APH) or eth1 (for APL/BT11)<br><br>**Default**: `false` |
| `PortConfig` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | eth0 is not allowed here. Property key is the interface(s) name (e.g. `eth1` or `eth1,eth2`). If spcified, this takes predecence over switch_config (switch_config requires user to configure all vlans manually, which is error-prone. thus deprecated) |
| `PwrConfig` | [`*models.ApPwrConfig`](../../doc/models/ap-pwr-config.md) | Optional | Power negotiation and peripheral power settings for an AP or AP profile |
| `RadioConfig` | [`*models.ApRadio`](../../doc/models/ap-radio.md) | Optional | Radio configuration settings for an access point |
| `Serial` | `*string` | Optional, Read-only | Manufacturer serial number for the access point |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Type` | `string` | Required, Constant, Read-only | Device Type. enum: `ap`<br><br>**Value**: `"ap"` |
| `UplinkPortConfig` | [`*models.ApUplinkPortConfig`](../../doc/models/ap-uplink-port-config.md) | Optional | AP Uplink port configuration |
| `UsbConfig` | [`*models.ApUsb`](../../doc/models/ap-usb.md) | Optional | Legacy USB integration settings for an access point<br><br>- Note: if native imagotag is enabled, BLE will be disabled automatically<br>- Note: legacy, new config moved to ESL Config. |
| `Vars` | `map[string]string` | Optional | Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `X` | `*float64` | Optional | Horizontal map position of the AP, in pixels |
| `Y` | `*float64` | Optional | Vertical map position of the AP, in pixels |
| `ZigbeeConfig` | [`*models.ApZigbee`](../../doc/models/ap-zigbee.md) | Optional | Zigbee radio and network settings applied to an AP or AP profile |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    deviceAp := models.DeviceAp{
        Aeroscout:            models.ToPointer(models.ApAeroscout{
            Enabled:              models.ToPointer(false),
            Host:                 models.NewOptional(models.ToPointer("host6")),
            LocateConnected:      models.ToPointer(false),
            Port:                 models.NewOptional(models.ToPointer(86)),
        }),
        Airista:              models.ToPointer(models.ApAirista{
            Enabled:              models.ToPointer(false),
            Host:                 models.NewOptional(models.ToPointer("host8")),
            Port:                 models.NewOptional(models.ToPointer(218)),
        }),
        BleConfig:            models.ToPointer(models.BleConfig{
            BeaconEnabled:           models.ToPointer(false),
            BeaconRate:              models.ToPointer(110),
            BeaconRateMode:          models.ToPointer(models.BleConfigBeaconRateModeEnum_CUSTOM),
            BeamDisabled:            []int{
                113,
                114,
                115,
            },
            CustomBlePacketEnabled:  models.ToPointer(false),
        }),
        Centrak:              models.ToPointer(models.ApCentrak{
            Enabled:              models.ToPointer(false),
        }),
        ClientBridge:         models.ToPointer(models.ApClientBridge{
            Auth:                 models.ToPointer(models.ApClientBridgeAuth{
                Psk:                  models.ToPointer("psk4"),
                Type:                 models.ToPointer(models.ApClientBridgeAuthTypeEnum_OPEN),
            }),
            Enabled:              models.ToPointer(false),
            Ssid:                 models.ToPointer("ssid0"),
        }),
        DeviceprofileId:      models.NewOptional(models.ToPointer(uuid.MustParse("6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"))),
        DisableEth1:          models.ToPointer(false),
        DisableEth2:          models.ToPointer(false),
        DisableEth3:          models.ToPointer(false),
        DisableModule:        models.ToPointer(false),
        FlowControl:          models.ToPointer(false),
        Height:               models.ToPointer(float64(2.75)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        MapId:                models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc")),
        Name:                 models.ToPointer("conference room"),
        Notes:                models.ToPointer("slightly off center"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Orientation:          models.ToPointer(45),
        PoePassthrough:       models.ToPointer(false),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                 "ap",
        Vars:                 map[string]string{
            "RADIUS_IP1": "172.31.2.5",
            "RADIUS_SECRET": "11s64632d",
        },
        X:                    models.ToPointer(float64(53.5)),
        Y:                    models.ToPointer(float64(173.1)),
    }

}
```

