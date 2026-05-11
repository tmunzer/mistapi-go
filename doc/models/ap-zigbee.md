
# Ap Zigbee

Zigbee AP settings

## Structure

`ApZigbee`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowJoin` | [`*models.ApZigbeeAllowJoinEnum`](../../doc/models/ap-zigbee-allow-join-enum.md) | Optional | Controls whether new Zigbee devices are allowed to join the network. enum: `always`, `manual`<br><br>**Default**: `"manual"` |
| `Channel` | `*int` | Optional | Zigbee channel (2.4 GHz). `0` means auto; valid fixed values are 11–26<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 26` |
| `Enabled` | `*bool` | Optional | Whether to enable Zigbee on this AP<br><br>**Default**: `false` |
| `ExtendedPanId` | `models.Optional[string]` | Optional | Extended PAN ID in hex string format; only applicable when `pan_id` is also specified |
| `PanId` | `models.Optional[string]` | Optional | PAN ID in hex string format; if not specified, assigned automatically |

## Example (as JSON)

```json
{
  "allow_join": "manual",
  "channel": 0,
  "enabled": false,
  "extended_pan_id": "1311768467294899695",
  "pan_id": "0x1234"
}
```

