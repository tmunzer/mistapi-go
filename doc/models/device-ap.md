
# Device Ap

AP

## Structure

`DeviceAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aeroscout` | [`*models.ApAeroscout`](../../doc/models/ap-aeroscout.md) | Optional | Aeroscout AP settings |
| `BleConfig` | [`*models.BleConfig`](../../doc/models/ble-config.md) | Optional | BLE AP settings |
| `Centrak` | [`*models.ApCentrak`](../../doc/models/ap-centrak.md) | Optional | - |
| `ClientBridge` | [`*models.ApClientBridge`](../../doc/models/ap-client-bridge.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | - |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional | - |
| `DisableEth1` | `*bool` | Optional | whether to disable eth1 port<br>**Default**: `false` |
| `DisableEth2` | `*bool` | Optional | whether to disable eth2 port<br>**Default**: `false` |
| `DisableEth3` | `*bool` | Optional | whether to disable eth3 port<br>**Default**: `false` |
| `DisableModule` | `*bool` | Optional | whether to disable module port<br>**Default**: `false` |
| `EslConfig` | [`*models.ApEslConfig`](../../doc/models/ap-esl-config.md) | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Height` | `*float64` | Optional | height, in meters, optional |
| `Id` | `*uuid.UUID` | Optional | - |
| `Image1Url` | `models.Optional[string]` | Optional | - |
| `Image2Url` | `models.Optional[string]` | Optional | - |
| `Image3Url` | `models.Optional[string]` | Optional | - |
| `IotConfig` | [`*models.ApIot`](../../doc/models/ap-iot.md) | Optional | IoT AP settings |
| `IpConfig` | [`*models.ApIpConfig`](../../doc/models/ap-ip-config.md) | Optional | IP AP settings |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | LED AP settings |
| `Locked` | `*bool` | Optional | whether this map is considered locked down |
| `Mac` | `*string` | Optional | device MAC address |
| `MapId` | `*uuid.UUID` | Optional | map where the device belongs to |
| `Mesh` | [`*models.ApMesh`](../../doc/models/ap-mesh.md) | Optional | Mesh AP settings |
| `Model` | `*string` | Optional | device Model |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Notes` | `*string` | Optional | any notes about this AP |
| `NtpServers` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Orientation` | `*int` | Optional | orientation, 0-359, in degrees, up is 0, right is 90. |
| `PoePassthrough` | `*bool` | Optional | whether to enable power out through module port (for APH) or eth1 (for APL/BT11)<br>**Default**: `false` |
| `PortConfig` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | eth0 is not allowed here.<br>Property key is the interface(s) name (e.g. "eth1" or"eth1,eth2") |
| `PwrConfig` | [`*models.ApPwrConfig`](../../doc/models/ap-pwr-config.md) | Optional | power related configs |
| `RadioConfig` | [`*models.ApRadio`](../../doc/models/ap-radio.md) | Optional | Radio AP settings |
| `Serial` | `*string` | Optional | device Serial |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Type` | [`*models.DeviceTypeApEnum`](../../doc/models/device-type-ap-enum.md) | Optional | Device Type |
| `UplinkPortConfig` | [`*models.ApUplinkPortConfig`](../../doc/models/ap-uplink-port-config.md) | Optional | - |
| `UsbConfig` | [`*models.ApUsb`](../../doc/models/ap-usb.md) | Optional | USB AP settings<br>Note: if native imagotag is enabled, BLE will be disabled automatically<br>Note: legacy, new config moved to ESL Config. |
| `Vars` | `map[string]string` | Optional | a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `X` | `*float64` | Optional | x in pixel |
| `Y` | `*float64` | Optional | y in pixel |

## Example (as JSON)

```json
{
  "deviceprofile_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
  "disable_eth1": false,
  "disable_eth2": false,
  "disable_eth3": false,
  "disable_module": false,
  "height": 2.75,
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "name": "conference room",
  "notes": "slightly off center",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "orientation": 45,
  "poe_passthrough": false,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "x": 53.5,
  "y": 173.1,
  "aeroscout": {
    "enabled": false,
    "host": "host6",
    "locate_connected": false
  },
  "ble_config": {
    "beacon_enabled": false,
    "beacon_rate": 110,
    "beacon_rate_mode": "default",
    "beam_disabled": [
      113,
      114,
      115
    ],
    "custom_ble_packet_enabled": false
  },
  "centrak": {
    "enabled": false
  },
  "client_bridge": {
    "auth": {
      "psk": "psk4",
      "type": "open"
    },
    "enabled": false,
    "ssid": "ssid0"
  },
  "created_time": 157.04
}
```
