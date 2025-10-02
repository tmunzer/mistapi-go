
# Gateway Port Wan Source Nat

Only if `usage`==`wan`, optional. By default, source-NAT is performed on all WAN Ports using the interface-ip

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayPortWanSourceNat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | Or to disable the source-nat<br><br>**Default**: `false` |
| `Nat6Pool` | `*string` | Optional | If alternative nat_pool is desired |
| `NatPool` | `*string` | Optional | If alternative nat_pool is desired |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disabled": false,
  "nat6_pool": "2601:1700:43c0:dc0:20c:29ff:fea7:93bc/126",
  "nat_pool": "64.2.4.0/30",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

