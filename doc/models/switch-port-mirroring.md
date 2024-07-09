
# Switch Port Mirroring

Property key is the port mirroring instance name
port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output.

## Structure

`SwitchPortMirroring`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortMirror` | [`*models.GatewayPortMirroringPortMirror`](../../doc/models/gateway-port-mirroring-port-mirror.md) | Optional | - |

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
    "run_length": 214
  }
}
```

