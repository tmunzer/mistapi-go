
# Gateway Path Preferences

## Structure

`GatewayPathPreferences`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Paths` | [`[]models.GatewayPathPreferencesPath`](../../doc/models/gateway-path-preferences-path.md) | Optional | - |
| `Strategy` | [`*models.GatewayPathStrategyEnum`](../../doc/models/gateway-path-strategy-enum.md) | Optional | enum: `ecmp`, `ordered`, `weighted`<br>**Default**: `"ordered"` |

## Example (as JSON)

```json
{
  "strategy": "ordered",
  "paths": [
    {
      "cost": 118,
      "disabled": false,
      "gateway_ip": "gateway_ip0",
      "internet_access": false,
      "name": "name4"
    }
  ]
}
```

