
# Discovered Switch Metric Ap

*This model accepts additional fields of type interface{}.*

## Structure

`DiscoveredSwitchMetricAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hostname` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `PoeStatus` | `*bool` | Optional | - |
| `Port` | `*string` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `PowerDraw` | `*int` | Optional | - |
| `When` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "hostname": "hostname8",
  "mac": "mac0",
  "poe_status": false,
  "port": "port6",
  "port_id": "port_id4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

