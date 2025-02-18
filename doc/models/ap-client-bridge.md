
# Ap Client Bridge

*This model accepts additional fields of type interface{}.*

## Structure

`ApClientBridge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Auth` | [`*models.ApClientBridgeAuth`](../../doc/models/ap-client-bridge-auth.md) | Optional | - |
| `Enabled` | `*bool` | Optional | When acted as client bridge:<br><br>* only 5G radio can be used<br>* will not serve as AP on any radios<br>**Default**: `false` |
| `Ssid` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "ssid": "Uplink-SSID",
  "auth": {
    "psk": "psk4",
    "type": "open",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

