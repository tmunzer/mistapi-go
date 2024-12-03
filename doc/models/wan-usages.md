
# Wan Usages

*This model accepts additional fields of type interface{}.*

## Structure

`WanUsages`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | - |
| `PathType` | `*string` | Optional | - |
| `PathWeight` | `*int` | Optional | - |
| `PeerMac` | `*string` | Optional | - |
| `PeerPortId` | `*string` | Optional | - |
| `Policy` | `*string` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `Tenant` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "5c5b35000001",
  "path_type": "vpn",
  "path_weight": 10,
  "peer_mac": "0200018c95e1",
  "peer_port_id": "ge-0/0/3",
  "policy": "policy1",
  "port_id": "ge-0/0/0.0",
  "tenant": "tenant1",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

