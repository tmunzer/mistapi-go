
# Org Setting Gateway Mgmt Overlay Ip

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingGatewayMgmtOverlayIp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | When it's going overlay, a routable IP to overlay will be required |
| `Node1Ip` | `*string` | Optional | For SSR HA cluster, another IP for node1 will be required, too |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ip": "ip2",
  "node1_ip": "node1_ip0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

