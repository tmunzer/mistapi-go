
# Switch Port Mirroring Property

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchPortMirroringProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InputNetworksIngress` | `[]string` | Optional | at least one of the `input_port_ids_ingress`, `input_port_ids_egress` or `input_networks_ingress` should be specified |
| `InputPortIdsEgress` | `[]string` | Optional | at least one of the `input_port_ids_ingress`, `input_port_ids_egress` or `input_networks_ingress` should be specified |
| `InputPortIdsIngress` | `[]string` | Optional | at least one of the `input_port_ids_ingress`, `input_port_ids_egress` or `input_networks_ingress` should be specified |
| `OutputNetwork` | `*string` | Optional | exaclty one of the `output_port_id` or `output_network` should be provided |
| `OutputPortId` | `*string` | Optional | exaclty one of the `output_port_id` or `output_network` should be provided |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "output_network": "analyze",
  "output_port_id": "ge-0/0/5",
  "input_networks_ingress": [
    "input_networks_ingress2",
    "input_networks_ingress3",
    "input_networks_ingress4"
  ],
  "input_port_ids_egress": [
    "input_port_ids_egress8"
  ],
  "input_port_ids_ingress": [
    "input_port_ids_ingress6",
    "input_port_ids_ingress7",
    "input_port_ids_ingress8"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

