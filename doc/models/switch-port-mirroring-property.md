
# Switch Port Mirroring Property

Input and output settings for one switch port mirroring session

## Structure

`SwitchPortMirroringProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InputNetworksIngress` | `[]string` | Optional | At least one of the `input_port_ids_ingress`, `input_port_ids_egress` or `input_networks_ingress ` should be specified |
| `InputPortIdsEgress` | `[]string` | Optional | At least one of the `input_port_ids_ingress`, `input_port_ids_egress` or `input_networks_ingress ` should be specified |
| `InputPortIdsIngress` | `[]string` | Optional | At least one of the `input_port_ids_ingress`, `input_port_ids_egress` or `input_networks_ingress ` should be specified |
| `OutputIpAddress` | `*string` | Optional | Exactly one of the `output_ip_address`, `output_port_id` or `output_network` should be provided |
| `OutputNetwork` | `*string` | Optional | Exactly one of the `output_ip_address`, `output_port_id` or `output_network` should be provided |
| `OutputPortId` | `*string` | Optional | Exactly one of the `output_ip_address`, `output_port_id` or `output_network` should be provided |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchPortMirroringProperty := models.SwitchPortMirroringProperty{
        InputNetworksIngress: []string{
            "input_networks_ingress2",
            "input_networks_ingress3",
        },
        InputPortIdsEgress:   []string{
            "input_port_ids_egress8",
            "input_port_ids_egress9",
            "input_port_ids_egress0",
        },
        InputPortIdsIngress:  []string{
            "input_port_ids_ingress6",
            "input_port_ids_ingress7",
        },
        OutputIpAddress:      models.ToPointer("1.2.3.4"),
        OutputNetwork:        models.ToPointer("analyze"),
        OutputPortId:         models.ToPointer("ge-0/0/5"),
    }

}
```

