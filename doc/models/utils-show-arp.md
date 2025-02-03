
# Utils Show Arp

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowArp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 300` |
| `Interval` | `*int` | Optional | Rate at which output will refresh<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 10` |
| `Ip` | `*string` | Optional | IP Address |
| `PortId` | `*string` | Optional | Device Port ID |
| `Vrf` | `*string` | Optional | VRF Name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 0,
  "interval": 0,
  "ip": "192.168.30.7",
  "port_id": "ge-0/0/0.0",
  "vrf": "guest",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

