
# Ap Usb

USB AP settings
Note: if native imagotag is enabled, BLE will be disabled automatically
Note: legacy, new config moved to ESL Config.

*This model accepts additional fields of type interface{}.*

## Structure

`ApUsb`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cacert` | `models.Optional[string]` | Optional | only if `type`==`imagotag` |
| `Channel` | `*int` | Optional | only if `type`==`imagotag`, channel selection, not needed by default, required for manual channel override only |
| `Enabled` | `*bool` | Optional | whether to enable any usb config |
| `Host` | `*string` | Optional | only if `type`==`imagotag` |
| `Port` | `*int` | Optional | only if `type`==`imagotag`<br>**Default**: `0` |
| `Type` | [`*models.ApUsbTypeEnum`](../../doc/models/ap-usb-type-enum.md) | Optional | usb config type. enum: `hanshow`, `imagotag`, `solum` |
| `VerifyCert` | `*bool` | Optional | only if `type`==`imagotag`, whether to turn on SSL verification |
| `VlanId` | `*int` | Optional | only if `type`==`solum` or `type`==`hanshow`<br>**Default**: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "channel": 3,
  "host": "1.1.1.1",
  "port": 0,
  "type": "imagotag",
  "vlan_id": 1,
  "cacert": "cacert4",
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

