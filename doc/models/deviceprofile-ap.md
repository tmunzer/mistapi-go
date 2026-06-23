
# Deviceprofile Ap

AP device profile configuration applied to APs at a site or organization

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceprofileAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aeroscout` | [`*models.ApAeroscout`](../../doc/models/ap-aeroscout.md) | Optional | AeroScout location integration settings applied to an AP or AP profile |
| `Airista` | [`*models.ApAirista`](../../doc/models/ap-airista.md) | Optional | Airista RTLS integration settings for an AP |
| `BleConfig` | [`*models.BleConfig`](../../doc/models/ble-config.md) | Optional | Bluetooth Low Energy beacon and asset advertising settings for an AP |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DisableEth1` | `*bool` | Optional | Whether to disable eth1 port<br><br>**Default**: `false` |
| `DisableEth2` | `*bool` | Optional | Whether to disable eth2 port<br><br>**Default**: `false` |
| `DisableEth3` | `*bool` | Optional | Whether to disable eth3 port<br><br>**Default**: `false` |
| `DisableModule` | `*bool` | Optional | Whether to disable module port<br><br>**Default**: `false` |
| `EslConfig` | [`*models.ApEslConfig`](../../doc/models/ap-esl-config.md) | Optional | Electronic shelf label integration settings for an AP |
| `ForSite` | `*bool` | Optional, Read-only | Whether this AP profile is scoped directly to a site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `IotConfig` | [`*models.ApIot`](../../doc/models/ap-iot.md) | Optional | Digital and analog IoT port settings applied to an AP or AP profile |
| `IpConfig` | [`*models.ApIpConfig`](../../doc/models/ap-ip-config.md) | Optional | Management IP addressing settings for an access point |
| `LacpConfig` | [`*models.DeviceApLacpConfig`](../../doc/models/device-ap-lacp-config.md) | Optional | LACP settings for supported AP Ethernet uplinks |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | Indicator light settings for an access point |
| `Mesh` | [`*models.ApMesh`](../../doc/models/ap-mesh.md) | Optional | Wireless mesh settings for an access point |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MqttConfig` | [`*models.ApMqtt`](../../doc/models/ap-mqtt.md) | Optional | MQTT broker publishing settings for an AP; use `mqtt_topic` on individual AssetFilter entries to specify which MQTT topic each matching BLE advertisement is forwarded to |
| `Name` | `models.Optional[string]` | Optional | Display name of the AP device profile |
| `NtpServers` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PoePassthrough` | `*bool` | Optional | Whether to enable power out through module port (for APH) or eth1 (for APL/BT11)<br><br>**Default**: `false` |
| `PortConfig` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | eth0 is not allowed here. Property key is the interface(s) name (e.g. `eth1` or `eth1,eth2`). If specified, this takes precedence over switch_config (deprecated) |
| `PwrConfig` | [`*models.ApPwrConfig`](../../doc/models/ap-pwr-config.md) | Optional | Power negotiation and peripheral power settings for an AP or AP profile |
| `RadioConfig` | [`*models.ApRadio`](../../doc/models/ap-radio.md) | Optional | Radio configuration settings for an access point |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SwitchConfig` | [`*models.ApSwitch`](../../doc/models/ap-switch.md) | Optional | Deprecated AP switch VLAN control settings for advanced per-port configuration |
| `Type` | `string` | Required, Constant | Device Type. enum: `ap`<br><br>**Value**: `"ap"` |
| `UplinkPortConfig` | [`*models.ApUplinkPortConfig`](../../doc/models/ap-uplink-port-config.md) | Optional | AP Uplink port configuration |
| `UsbConfig` | [`*models.ApUsb`](../../doc/models/ap-usb.md) | Optional | Legacy USB integration settings for an access point<br><br>- Note: if native imagotag is enabled, BLE will be disabled automatically<br>- Note: legacy, new config moved to ESL Config. |
| `Vars` | `map[string]string` | Optional | Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `ZigbeeConfig` | [`*models.ApZigbee`](../../doc/models/ap-zigbee.md) | Optional | Zigbee radio and network settings applied to an AP or AP profile |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    deviceprofileAp := models.DeviceprofileAp{
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
        CreatedTime:          models.ToPointer(float64(45.42)),
        DisableEth1:          models.ToPointer(false),
        DisableEth2:          models.ToPointer(false),
        DisableEth3:          models.ToPointer(false),
        DisableModule:        models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PoePassthrough:       models.ToPointer(false),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                 "ap",
        Vars:                 map[string]string{
            "RADIUS_IP1": "172.31.2.5",
            "RADIUS_SECRET": "11s64632d",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

