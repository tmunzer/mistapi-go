
# Marvis

*This model accepts additional fields of type interface{}.*

## Structure

`Marvis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoOperations` | [`*models.MarvisAutoOperations`](../../doc/models/marvis-auto-operations.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "auto_operations": {
    "ap_insufficient_capacity": false,
    "ap_loop": false,
    "ap_non_compliant": false,
    "bounce_port_for_abnormal_poe_client": false,
    "disable_port_when_ddos_protocol_violation": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

