
# Device Ap

AP

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aeroscout` | [`*models.ApAeroscout`](../../doc/models/ap-aeroscout.md) | Optional | Aeroscout AP settings |
| `BleConfig` | [`*models.BleConfig`](../../doc/models/ble-config.md) | Optional | BLE AP settings |
| `Centrak` | [`*models.ApCentrak`](../../doc/models/ap-centrak.md) | Optional | - |
| `ClientBridge` | [`*models.ApClientBridge`](../../doc/models/ap-client-bridge.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional | - |
| `DisableEth1` | `*bool` | Optional | Whether to disable eth1 port<br>**Default**: `false` |
| `DisableEth2` | `*bool` | Optional | Whether to disable eth2 port<br>**Default**: `false` |
| `DisableEth3` | `*bool` | Optional | Whether to disable eth3 port<br>**Default**: `false` |
| `DisableModule` | `*bool` | Optional | Whether to disable module port<br>**Default**: `false` |
| `EslConfig` | [`*models.ApEslConfig`](../../doc/models/ap-esl-config.md) | Optional | - |
| `FlowControl` | `*bool` | Optional | For some AP models, flow_control can be enabled to address some switch compatibility issue<br>**Default**: `false` |
| `ForSite` | `*bool` | Optional | - |
| `Height` | `*float64` | Optional | Height, in meters, optional |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Image1Url` | `models.Optional[string]` | Optional | - |
| `Image2Url` | `models.Optional[string]` | Optional | - |
| `Image3Url` | `models.Optional[string]` | Optional | - |
| `IotConfig` | [`*models.ApIot`](../../doc/models/ap-iot.md) | Optional | IoT AP settings |
| `IpConfig` | [`*models.ApIpConfig`](../../doc/models/ap-ip-config.md) | Optional | IP AP settings |
| `LacpConfig` | [`*models.DeviceApLacpConfig`](../../doc/models/device-ap-lacp-config.md) | Optional | - |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | LED AP settings |
| `Locked` | `*bool` | Optional | Whether this map is considered locked down |
| `Mac` | `*string` | Optional | Device MAC address |
| `MapId` | `*uuid.UUID` | Optional | Map where the device belongs to |
| `Mesh` | [`*models.ApMesh`](../../doc/models/ap-mesh.md) | Optional | Mesh AP settings |
| `Model` | `*string` | Optional | Device Model |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `Notes` | `*string` | Optional | Any notes about this AP |
| `NtpServers` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Orientation` | `*int` | Optional | Orientation, 0-359, in degrees, up is 0, right is 90.<br>**Constraints**: `>= 0`, `<= 359` |
| `PoePassthrough` | `*bool` | Optional | Whether to enable power out through module port (for APH) or eth1 (for APL/BT11)<br>**Default**: `false` |
| `PortConfig` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | eth0 is not allowed here. Property key is the interface(s) name (e.g. `eth1` or `eth1,eth2`) |
| `PwrConfig` | [`*models.ApPwrConfig`](../../doc/models/ap-pwr-config.md) | Optional | Power related configs |
| `RadioConfig` | [`*models.ApRadio`](../../doc/models/ap-radio.md) | Optional | Radio AP settings |
| `Serial` | `*string` | Optional | Device Serial |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Type` | `string` | Required, Constant | Device Type. enum: `ap`<br>**Value**: `"ap"` |
| `UplinkPortConfig` | [`*models.ApUplinkPortConfig`](../../doc/models/ap-uplink-port-config.md) | Optional | AP Uplink port configuration |
| `UsbConfig` | [`*models.ApUsb`](../../doc/models/ap-usb.md) | Optional | USB AP settings<br><br>- Note: if native imagotag is enabled, BLE will be disabled automatically<br>- Note: legacy, new config moved to ESL Config. |
| `Vars` | `map[string]string` | Optional | Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `X` | `*float64` | Optional | X in pixel |
| `Y` | `*float64` | Optional | Y in pixel |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "deviceprofile_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
  "disable_eth1": false,
  "disable_eth2": false,
  "disable_eth3": false,
  "disable_module": false,
  "flow_control": false,
  "height": 2.75,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "name": "conference room",
  "notes": "slightly off center",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "orientation": 45,
  "poe_passthrough": false,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "ap",
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "x": 53.5,
  "y": 173.1,
  "aeroscout": {
    "enabled": false,
    "host": "host6",
    "locate_connected": false,
    "port": 86,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "ble_config": {
    "beacon_enabled": false,
    "beacon_rate": 110,
    "beacon_rate_mode": "custom",
    "beam_disabled": [
      113,
      114,
      115
    ],
    "custom_ble_packet_enabled": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "centrak": {
    "enabled": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "client_bridge": {
    "auth": {
      "psk": "psk4",
      "type": "open",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "enabled": false,
    "ssid": "ssid0",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "created_time": 87.82,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

