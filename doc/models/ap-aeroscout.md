
# Ap Aeroscout

Aeroscout AP settings

## Structure

`ApAeroscout`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | whether to enable aeroscout config<br>**Default**: `false` |
| `Host` | `models.Optional[string]` | Optional | required if enabled, aeroscout server host |
| `LocateConnected` | `*bool` | Optional | whether to enable the feature to allow wireless clients data received and sent to AES server for location calculation<br>**Default**: `true` |

## Example (as JSON)

```json
{
  "enabled": false,
  "host": "aero.pvt.net",
  "locate_connected": true
}
```

