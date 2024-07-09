
# Gateway Port Wan Source Nat

optional, by default, source-NAT is performed on all WAN Ports using the interface-ip

## Structure

`GatewayPortWanSourceNat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | or to disable the source-nat<br>**Default**: `false` |
| `NatPool` | `*string` | Optional | if alternative nat_pool is desired |

## Example (as JSON)

```json
{
  "disabled": false,
  "nat_pool": "64.2.4.0/30"
}
```

