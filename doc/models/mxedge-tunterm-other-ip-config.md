
# Mxedge Tunterm Other Ip Config

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeTuntermOtherIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `string` | Required | - |
| `Netmask` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ip": "ip2",
  "netmask": "netmask8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

