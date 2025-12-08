
# Ap Airista

## Structure

`ApAirista`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable Airista config<br><br>**Default**: `false` |
| `Host` | `models.Optional[string]` | Optional | Required if enabled, Airista server host |
| `Port` | `models.Optional[int]` | Optional | **Default**: `1144` |

## Example (as JSON)

```json
{
  "enabled": false,
  "host": "airista.pvt.net",
  "port": 1144
}
```

