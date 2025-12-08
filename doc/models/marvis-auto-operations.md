
# Marvis Auto Operations

## Structure

`MarvisAutoOperations`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApInsufficientCapacity` | `*bool` | Optional | **Default**: `false` |
| `ApLoop` | `*bool` | Optional | **Default**: `false` |
| `ApNonCompliant` | `*bool` | Optional | **Default**: `false` |
| `BouncePortForAbnormalPoeClient` | `*bool` | Optional | **Default**: `false` |
| `DisablePortWhenDdosProtocolViolation` | `*bool` | Optional | **Default**: `false` |
| `DisablePortWhenRogueDhcpServerDetected` | `*bool` | Optional | **Default**: `false` |
| `GatewayNonCompliant` | `*bool` | Optional | **Default**: `false` |
| `SwitchMisconfiguredPort` | `*bool` | Optional | **Default**: `false` |
| `SwitchPortStuck` | `*bool` | Optional | **Default**: `false` |

## Example (as JSON)

```json
{
  "ap_insufficient_capacity": false,
  "ap_loop": false,
  "ap_non_compliant": false,
  "bounce_port_for_abnormal_poe_client": false,
  "disable_port_when_ddos_protocol_violation": false,
  "disable_port_when_rogue_dhcp_server_detected": false,
  "gateway_non_compliant": false,
  "switch_misconfigured_port": false,
  "switch_port_stuck": false
}
```

