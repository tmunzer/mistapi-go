
# Marvis Auto Operations

*This model accepts additional fields of type interface{}.*

## Structure

`MarvisAutoOperations`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BouncePortForAbnormalPoeClient` | `*bool` | Optional | **Default**: `false` |
| `DisablePortWhenDdosProtocolViolation` | `*bool` | Optional | **Default**: `false` |
| `DisablePortWhenRogueDhcpServerDetected` | `*bool` | Optional | **Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bounce_port_for_abnormal_poe_client": false,
  "disable_port_when_ddos_protocol_violation": false,
  "disable_port_when_rogue_dhcp_server_detected": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

