
# Switch Matching

Switch template

## Structure

`SwitchMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | - |
| `Rules` | [`[]models.SwitchMatchingRule`](../../doc/models/switch-matching-rule.md) | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "enable": false,
  "rules": [
    {
      "additional_config_cmds": [
        "additional_config_cmds8"
      ],
      "match_role": "match_role8",
      "name": "name8",
      "port_config": {
        "key0": {
          "ae_disable_lacp": false,
          "ae_idx": 230,
          "ae_lacp_slow": false,
          "aggregated": false,
          "critical": false,
          "usage": "usage6"
        },
        "key1": {
          "ae_disable_lacp": false,
          "ae_idx": 230,
          "ae_lacp_slow": false,
          "aggregated": false,
          "critical": false,
          "usage": "usage6"
        }
      },
      "port_mirroring": {
        "key0": {
          "input_networks_ingress": [
            "input_networks_ingress8"
          ],
          "input_port_ids_egress": [
            "input_port_ids_egress4",
            "input_port_ids_egress5"
          ],
          "input_port_ids_ingress": [
            "input_port_ids_ingress2"
          ],
          "output_network": "output_network4",
          "output_port_id": "output_port_id2"
        },
        "key1": {
          "input_networks_ingress": [
            "input_networks_ingress8"
          ],
          "input_port_ids_egress": [
            "input_port_ids_egress4",
            "input_port_ids_egress5"
          ],
          "input_port_ids_ingress": [
            "input_port_ids_ingress2"
          ],
          "output_network": "output_network4",
          "output_port_id": "output_port_id2"
        },
        "key2": {
          "input_networks_ingress": [
            "input_networks_ingress8"
          ],
          "input_port_ids_egress": [
            "input_port_ids_egress4",
            "input_port_ids_egress5"
          ],
          "input_port_ids_ingress": [
            "input_port_ids_ingress2"
          ],
          "output_network": "output_network4",
          "output_port_id": "output_port_id2"
        }
      }
    },
    {
      "additional_config_cmds": [
        "additional_config_cmds8"
      ],
      "match_role": "match_role8",
      "name": "name8",
      "port_config": {
        "key0": {
          "ae_disable_lacp": false,
          "ae_idx": 230,
          "ae_lacp_slow": false,
          "aggregated": false,
          "critical": false,
          "usage": "usage6"
        },
        "key1": {
          "ae_disable_lacp": false,
          "ae_idx": 230,
          "ae_lacp_slow": false,
          "aggregated": false,
          "critical": false,
          "usage": "usage6"
        }
      },
      "port_mirroring": {
        "key0": {
          "input_networks_ingress": [
            "input_networks_ingress8"
          ],
          "input_port_ids_egress": [
            "input_port_ids_egress4",
            "input_port_ids_egress5"
          ],
          "input_port_ids_ingress": [
            "input_port_ids_ingress2"
          ],
          "output_network": "output_network4",
          "output_port_id": "output_port_id2"
        },
        "key1": {
          "input_networks_ingress": [
            "input_networks_ingress8"
          ],
          "input_port_ids_egress": [
            "input_port_ids_egress4",
            "input_port_ids_egress5"
          ],
          "input_port_ids_ingress": [
            "input_port_ids_ingress2"
          ],
          "output_network": "output_network4",
          "output_port_id": "output_port_id2"
        },
        "key2": {
          "input_networks_ingress": [
            "input_networks_ingress8"
          ],
          "input_port_ids_egress": [
            "input_port_ids_egress4",
            "input_port_ids_egress5"
          ],
          "input_port_ids_ingress": [
            "input_port_ids_ingress2"
          ],
          "output_network": "output_network4",
          "output_port_id": "output_port_id2"
        }
      }
    }
  ]
}
```

