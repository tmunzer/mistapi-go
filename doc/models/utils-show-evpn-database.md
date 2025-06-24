
# Utils Show Evpn Database

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowEvpnDatabase`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 300` |
| `Interval` | `*int` | Optional | Rate at which output will refresh<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 10` |
| `Mac` | `*string` | Optional | Client mac filter |
| `PortId` | `*string` | Optional | Interface name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 0,
  "interval": 0,
  "mac": "f8c1165c6400",
  "port_id": "ge-0/0/0.0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

