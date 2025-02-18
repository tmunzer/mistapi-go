
# Ap Aeroscout

Aeroscout AP settings

*This model accepts additional fields of type interface{}.*

## Structure

`ApAeroscout`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable aeroscout config<br>**Default**: `false` |
| `Host` | `models.Optional[string]` | Optional | Required if enabled, aeroscout server host |
| `LocateConnected` | `*bool` | Optional | Whether to enable the feature to allow wireless clients data received and sent to AES server for location calculation<br>**Default**: `true` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "host": "aero.pvt.net",
  "locate_connected": true,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

