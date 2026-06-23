
# Synthetictest Config

Synthetic test configuration for Marvis Minis

## Structure

`SynthetictestConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aggressiveness` | [`*models.SynthetictestConfigAggressivenessEnum`](../../doc/models/synthetictest-config-aggressiveness-enum.md) | Optional | Aggressiveness level for a synthetic test. enum: `auto`, `high`, `med`, `low`<br><br>**Default**: `"auto"` |
| `CustomProbes` | [`map[string]models.SynthetictestConfigCustomProbe`](../../doc/models/synthetictest-config-custom-probe.md) | Optional | Custom probes to be used for synthetic tests |
| `Disabled` | `*bool` | Optional | Whether synthetic tests are disabled<br><br>**Default**: `false` |
| `LanNetworks` | [`[]models.SynthetictestConfigLanNetwork`](../../doc/models/synthetictest-config-lan-network.md) | Optional | List of networks to be used for synthetic tests |
| `Vlans` | [`[]models.SynthetictestConfigVlan`](../../doc/models/synthetictest-config-vlan.md) | Optional | Deprecated VLAN-based synthetic test configurations |
| `WanSpeedtest` | [`*models.SynthetictestConfigWanSpeedtest`](../../doc/models/synthetictest-config-wan-speedtest.md) | Optional | WAN speedtest scheduling settings for synthetic tests |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    synthetictestConfig := models.SynthetictestConfig{
        Aggressiveness:       models.ToPointer(models.SynthetictestConfigAggressivenessEnum_AUTO),
        CustomProbes:         map[string]models.SynthetictestConfigCustomProbe{
            "key0": models.SynthetictestConfigCustomProbe{
                Aggressiveness:       models.ToPointer(models.SynthetictestConfigAggressivenessEnum_MED),
                Target:               models.ToPointer("target6"),
                Threshold:            models.ToPointer(178),
                Type:                 models.ToPointer(models.SynthetictestConfigCustomProbeTypeEnum_REACHABILITY),
            },
        },
        Disabled:             models.ToPointer(false),
        LanNetworks:          []models.SynthetictestConfigLanNetwork{
            models.SynthetictestConfigLanNetwork{
                Networks:             []string{
                    "networks4",
                },
                Probes:               []string{
                    "probes5",
                    "probes4",
                },
            },
            models.SynthetictestConfigLanNetwork{
                Networks:             []string{
                    "networks4",
                },
                Probes:               []string{
                    "probes5",
                    "probes4",
                },
            },
            models.SynthetictestConfigLanNetwork{
                Networks:             []string{
                    "networks4",
                },
                Probes:               []string{
                    "probes5",
                    "probes4",
                },
            },
        },
        Vlans:                []models.SynthetictestConfigVlan{
            models.SynthetictestConfigVlan{
                CustomTestUrls:       []string{
                    "custom_test_urls9",
                },
                Disabled:             models.ToPointer(false),
                Probes:               []string{
                    "probes7",
                    "probes8",
                    "probes9",
                },
                VlanIds:              []models.VlanIdWithVariable{
                    models.VlanIdWithVariableContainer.FromString("String7"),
                    models.VlanIdWithVariableContainer.FromString("String8"),
                },
            },
        },
    }

}
```

