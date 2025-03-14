
# Gateway Port Mirroring

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayPortMirroring`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortMirror` | [`*models.GatewayPortMirroringPortMirror`](../../doc/models/gateway-port-mirroring-port-mirror.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "port_mirror": {
    "family_type": "family_type0",
    "ingress_port_ids": [
      "ingress_port_ids8",
      "ingress_port_ids7",
      "ingress_port_ids6"
    ],
    "output_port_id": "output_port_id2",
    "rate": 150,
    "run_length": 214,
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

