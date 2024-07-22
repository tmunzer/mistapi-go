
# Network Source Nat

if `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub

## Structure

`NetworkSourceNat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExternalIp` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "external_ip": "172.16.0.8/30"
}
```

