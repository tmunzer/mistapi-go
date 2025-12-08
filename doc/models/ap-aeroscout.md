
# Ap Aeroscout

Aeroscout AP settings

## Structure

`ApAeroscout`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable aeroscout config<br><br>**Default**: `false` |
| `Host` | `models.Optional[string]` | Optional | Required if enabled, aeroscout server host |
| `LocateConnected` | `*bool` | Optional | Whether to enable the feature to allow wireless clients data received and sent to AES server for location calculation<br><br>**Default**: `false` |
| `Port` | `models.Optional[int]` | Optional | **Default**: `1144` |

## Example (as JSON)

```json
{
  "enabled": false,
  "host": "aero.pvt.net",
  "locate_connected": false,
  "port": 1144
}
```

