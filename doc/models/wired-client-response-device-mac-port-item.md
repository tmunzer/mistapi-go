
# Wired Client Response Device Mac Port Item

*This model accepts additional fields of type interface{}.*

## Structure

`WiredClientResponseDeviceMacPortItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Ip` | `*string` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `PortParent` | `*string` | Optional | - |
| `Start` | `*string` | Optional | - |
| `Vlan` | `*int` | Optional | - |
| `When` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "device_mac": "device_mac6",
  "ip": "ip6",
  "port_id": "port_id2",
  "port_parent": "port_parent6",
  "start": "start6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

