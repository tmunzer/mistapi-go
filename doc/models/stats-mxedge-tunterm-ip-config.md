
# Stats Mxedge Tunterm Ip Config

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxedgeTuntermIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `*string` | Optional | - |
| `Ip` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "gateway": "192.168.11.1",
  "ip": "192.168.11.91",
  "netmask": "255.255.255.0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

