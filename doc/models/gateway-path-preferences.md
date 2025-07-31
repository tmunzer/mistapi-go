
# Gateway Path Preferences

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayPathPreferences`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Paths` | [`[]models.GatewayPathPreferencesPath`](../../doc/models/gateway-path-preferences-path.md) | Optional | - |
| `Strategy` | [`*models.GatewayPathStrategyEnum`](../../doc/models/gateway-path-strategy-enum.md) | Optional | enum: `ecmp`, `ordered`, `weighted`<br><br>**Default**: `"ordered"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
      "name": "name4",
      "type": "local",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "cost": 118,
      "disabled": false,
      "gateway_ip": "gateway_ip0",
      "internet_access": false,
      "name": "name4",
      "type": "local",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "cost": 118,
      "disabled": false,
      "gateway_ip": "gateway_ip0",
      "internet_access": false,
      "name": "name4",
      "type": "local",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

