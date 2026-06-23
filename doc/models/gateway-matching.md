
# Gateway Matching

Gateway matching configuration used to apply gateway-specific settings

## Structure

`GatewayMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | Whether gateway matching is enabled |
| `Rules` | [`[]models.GatewayMatchingRule`](../../doc/models/gateway-matching-rule.md) | Optional | Gateway matching rules used by a gateway matching configuration<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayMatching := models.GatewayMatching{
        Enable:               models.ToPointer(false),
        Rules:                []models.GatewayMatchingRule{
            models.GatewayMatchingRule{
                AdditionalConfigCmds: []string{
                    "additional_config_cmds8",
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
                    "key1": models.GatewayPortConfig{
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
            },
        },
    }

}
```

