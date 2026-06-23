
# Gateway Matching Rule

Gateway matching rule that applies settings when its match keys select a gateway

*This model accepts additional fields of type string.*

## Structure

`GatewayMatchingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `Name` | `*string` | Optional | Display name for the gateway matching rule |
| `PortConfig` | [`map[string]models.GatewayPortConfig`](../../doc/models/gateway-port-config.md) | Optional | Property key is the port(s) name or range (e.g. "ge-0/0/0-10"). |
| `AdditionalProperties` | `map[string]string` | Optional | Property key defines the type of matching. e.g: `match_name[0:3]`, `match_model[0-6]` or `match_role` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayMatchingRule := models.GatewayMatchingRule{
        AdditionalConfigCmds: []string{
            "additional_config_cmds2",
            "additional_config_cmds1",
            "additional_config_cmds0",
        },
        Name:                 models.ToPointer("name8"),
        PortConfig:           map[string]models.GatewayPortConfig{
            "key0": models.GatewayPortConfig{
                AeDisableLacp:          models.ToPointer(false),
                AeIdx:                  models.NewOptional(models.ToPointer("ae_idx8")),
                AeLacpForceUp:          models.ToPointer(false),
                Aggregated:             models.ToPointer(false),
                Critical:               models.ToPointer(false),
                Usage:                  models.GatewayPortUsageEnum_LAN,
            },
        },
        AdditionalProperties: map[string]string{
            "exampleAdditionalProperty": "gateway_matching_rule_additionalProperties8",
        },
    }

}
```

