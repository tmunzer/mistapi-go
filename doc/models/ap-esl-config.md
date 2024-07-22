
# Ap Esl Config

## Structure

`ApEslConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cacert` | `*string` | Optional | Only if `type`==`imagotag` or `type`==`native` |
| `Channel` | `*int` | Optional | Only if `type`==`imagotag` or `type`==`native` |
| `Enabled` | `*bool` | Optional | usb_config is ignored if esl_config enabled<br>**Default**: `false` |
| `Host` | `*string` | Optional | Only if `type`==`imagotag` or `type`==`native` |
| `Port` | `*int` | Optional | Only if `type`==`imagotag` or `type`==`native` |
| `Type` | [`*models.ApEslTypeEnum`](../../doc/models/ap-esl-type-enum.md) | Optional | note: ble_config will be ingored if esl_config is enabled and with native mode. |
| `VerifyCert` | `*bool` | Optional | Only if `type`==`imagotag` or `type`==`native` |
| `VlanId` | `*int` | Optional | Only if `type`==`solum` or `type`==`hanshow`<br>**Default**: `1` |

## Example (as JSON)

```json
{
  "channel": 3,
  "enabled": false,
  "host": "1.1.1.1",
  "port": 0,
  "type": "imagotag",
  "vlan_id": 1,
  "cacert": "cacert0"
}
```

