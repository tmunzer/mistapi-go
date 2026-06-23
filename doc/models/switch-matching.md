
# Switch Matching

Defines custom switch configuration based on different criteria

## Structure

`SwitchMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | Whether custom switch matching rules are enabled |
| `Rules` | [`[]models.SwitchMatchingRule`](../../doc/models/switch-matching-rule.md) | Optional | Ordered switch matching rules for conditional configuration<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchMatching := models.SwitchMatching{
        Enable:               models.ToPointer(false),
        Rules:                []models.SwitchMatchingRule{
            models.SwitchMatchingRule{
                AdditionalConfigCmds: []string{
                    "additional_config_cmds8",
                },
                DefaultPortUsage:     models.ToPointer("default_port_usage4"),
                IpConfig:             models.ToPointer(models.SwitchMatchingRuleIpConfig{
                    Network:              models.ToPointer("network6"),
                    Type:                 models.ToPointer(models.IpTypeEnum_DHCP),
                }),
                Name:                 models.ToPointer("name8"),
                OobIpConfig:          models.ToPointer(models.SwitchMatchingRuleOobIpConfig{
                    Type:                 models.ToPointer(models.IpTypeEnum_DHCP),
                    UseMgmtVrf:           models.ToPointer(false),
                    UseMgmtVrfForHostOut: models.ToPointer(false),
                }),
                AdditionalProperties: map[string]string{
                    "exampleAdditionalProperty": "switch_matching_rule_additionalProperties2",
                },
            },
        },
    }

}
```

