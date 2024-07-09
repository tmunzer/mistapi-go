
# Switch Port Mirroring Property

## Structure

`SwitchPortMirroringProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InputNetworksIngress` | `[]string` | Optional | at least one of the `ingress_port_ids`, `egress_port_ids` or `ingress_networks` should be specified |
| `InputPortIdsEgress` | `[]string` | Optional | at least one of the `ingress_port_ids`, `egress_port_ids` or `ingress_networks` should be specified |
| `InputPortIdsIngress` | `[]string` | Optional | at least one of the `ingress_port_ids`, `egress_port_ids` or `ingress_networks` should be specified |
| `OutputNetwork` | `*string` | Optional | - |
| `OutputPortId` | `*string` | Optional | exaclty on of the `output_port_id` or `output_network` should be provided |

## Example (as JSON)

```json
{
  "output_network": "analyze",
  "output_port_id": "ge-0/0/5",
  "input_networks_ingress": [
    "input_networks_ingress8",
    "input_networks_ingress9",
    "input_networks_ingress0"
  ],
  "input_port_ids_egress": [
    "input_port_ids_egress6"
  ],
  "input_port_ids_ingress": [
    "input_port_ids_ingress6",
    "input_port_ids_ingress5"
  ]
}
```

