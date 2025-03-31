
# Deviceprofile Ap

Device Profile

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceprofileAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aeroscout` | [`*models.ApAeroscout`](../../doc/models/ap-aeroscout.md) | Optional | Aeroscout AP settings |
| `BleConfig` | [`*models.BleConfig`](../../doc/models/ble-config.md) | Optional | BLE AP settings |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DisableEth1` | `*bool` | Optional | Whether to disable eth1 port<br>**Default**: `false` |
| `DisableEth2` | `*bool` | Optional | Whether to disable eth2 port<br>**Default**: `false` |
| `DisableEth3` | `*bool` | Optional | Whether to disable eth3 port<br>**Default**: `false` |
| `DisableModule` | `*bool` | Optional | Whether to disable module port<br>**Default**: `false` |
| `EslConfig` | [`*models.ApEslConfig`](../../doc/models/ap-esl-config.md) | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `IotConfig` | [`*models.ApIot`](../../doc/models/ap-iot.md) | Optional | IoT AP settings |
| `IpConfig` | [`*models.ApIpConfig`](../../doc/models/ap-ip-config.md) | Optional | IP AP settings |
| `LacpConfig` | [`*models.DeviceApLacpConfig`](../../doc/models/device-ap-lacp-config.md) | Optional | - |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | LED AP settings |
| `Mesh` | [`*models.ApMesh`](../../doc/models/ap-mesh.md) | Optional | Mesh AP settings |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Required | - |
| `NtpServers` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PoePassthrough` | `*bool` | Optional | Whether to enable power out through module port (for APH) or eth1 (for APL/BT11)<br>**Default**: `false` |
| `PortConfig` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | Property key is the interface(s) name (e.g. "eth1,eth2") |
| `PwrConfig` | [`*models.ApPwrConfig`](../../doc/models/ap-pwr-config.md) | Optional | Power related configs |
| `RadioConfig` | [`*models.ApRadio`](../../doc/models/ap-radio.md) | Optional | Radio AP settings |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SwitchConfig` | [`*models.ApSwitch`](../../doc/models/ap-switch.md) | Optional | For people who want to fully control the vlans (advanced) |
| `Type` | `string` | Required, Constant | Device Type. enum: `ap`<br>**Value**: `"ap"` |
| `UplinkPortConfig` | [`*models.ApUplinkPortConfig`](../../doc/models/ap-uplink-port-config.md) | Optional | AP Uplink port configuration |
| `UsbConfig` | [`*models.ApUsb`](../../doc/models/ap-usb.md) | Optional | USB AP settings<br><br>- Note: if native imagotag is enabled, BLE will be disabled automatically<br>- Note: legacy, new config moved to ESL Config. |
| `Vars` | `map[string]string` | Optional | Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disable_eth1": false,
  "disable_eth2": false,
  "disable_eth3": false,
  "disable_module": false,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "poe_passthrough": false,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "ap",
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
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
  "created_time": 42.6,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

