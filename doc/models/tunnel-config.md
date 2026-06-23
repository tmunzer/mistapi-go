
# Tunnel Config

Gateway tunnel configuration for provider-managed or custom tunnels

## Structure

`TunnelConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoProvision` | [`*models.TunnelConfigAutoProvision`](../../doc/models/tunnel-config-auto-provision.md) | Optional | Auto-provisioning configuration for the tunnel. This takes precedence over the `primary` and `secondary` nodes. |
| `IkeLifetime` | `*int` | Optional | Only if `provider`==`custom-ipsec`. IKE lifetime configured for the custom IPsec tunnel |
| `IkeMode` | [`*models.TunnelConfigIkeModeEnum`](../../doc/models/tunnel-config-ike-mode-enum.md) | Optional | Only if `provider`==`custom-ipsec`. enum: `aggressive`, `main`<br><br>**Default**: `"main"` |
| `IkeProposals` | [`[]models.TunnelConfigIkeProposal`](../../doc/models/tunnel-config-ike-proposal.md) | Optional | If `provider`==`custom-ipsec`, IKE proposals used for custom IPsec negotiation |
| `IpsecLifetime` | `*int` | Optional | If `provider`==`custom-ipsec`, IPsec lifetime configured for the custom tunnel |
| `IpsecProposals` | [`[]models.TunnelConfigIpsecProposal`](../../doc/models/tunnel-config-ipsec-proposal.md) | Optional | Only if `provider`==`custom-ipsec`. IPsec proposals used for custom IPsec negotiation |
| `LocalId` | `*string` | Optional | Required if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec` |
| `LocalSubnets` | `[]string` | Optional | List of Local protected subnet for policy-based IPSec negotiation |
| `Mode` | [`*models.TunnelConfigTunnelModeEnum`](../../doc/models/tunnel-config-tunnel-mode-enum.md) | Optional | Required if `provider`==`zscaler-gre`, `provider`==`jse-ipsec`. enum: `active-active`, `active-standby`<br><br>**Default**: `"active-standby"` |
| `Networks` | `[]string` | Optional | If `provider`==`custom-ipsec` or `provider`==`prisma-ipsec`, networks reachable via this tunnel |
| `Primary` | [`*models.TunnelConfigNode`](../../doc/models/tunnel-config-node.md) | Optional | Only if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec` |
| `Probe` | [`*models.TunnelConfigProbe`](../../doc/models/tunnel-config-probe.md) | Optional | Tunnel health probe settings |
| `Protocol` | [`*models.TunnelConfigProtocolEnum`](../../doc/models/tunnel-config-protocol-enum.md) | Optional | Only if `provider`==`custom-ipsec`. enum: `gre`, `ipsec` |
| `Provider` | [`*models.TunnelConfigProviderEnum`](../../doc/models/tunnel-config-provider-enum.md) | Optional | Only if `auto_provision.enabled`==`false`. enum: `custom-ipsec`, `custom-gre`, `jse-ipsec`, `prisma-ipsec`, `zscaler-gre`, `zscaler-ipsec` |
| `Psk` | `*string` | Optional | Required if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec` |
| `RemoteSubnets` | `[]string` | Optional | List of Remote protected subnet for policy-based IPSec negotiation |
| `Secondary` | [`*models.TunnelConfigNode`](../../doc/models/tunnel-config-node.md) | Optional | Only if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec` |
| `Version` | [`*models.TunnelConfigVersionEnum`](../../doc/models/tunnel-config-version-enum.md) | Optional | Only if `provider`==`custom-gre` or `provider`==`custom-ipsec`. enum: `1`, `2`<br><br>**Default**: `"2"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfig := models.TunnelConfig{
        AutoProvision:        models.ToPointer(models.TunnelConfigAutoProvision{
            Enabled:              models.ToPointer(false),
            Latlng:               models.ToPointer(models.TunnelConfigAutoProvisionLatLng{
                Lat:                  float64(144.64),
                Lng:                  float64(22.82),
            }),
            Primary:              models.ToPointer(models.TunnelConfigAutoProvisionNode{
                ProbeIps:             []string{
                    "probe_ips7",
                    "probe_ips8",
                    "probe_ips9",
                },
                WanNames:             []string{
                    "wan_names8",
                },
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            }),
            Provider:             models.TunnelConfigAutoProvisionProviderEnum_JSEIPSEC,
            Region:               models.ToPointer("region6"),
            Secondary:            models.ToPointer(models.TunnelConfigAutoProvisionNode{
                ProbeIps:             []string{
                    "probe_ips9",
                    "probe_ips0",
                    "probe_ips1",
                },
                WanNames:             []string{
                    "wan_names0",
                },
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            }),
        }),
        IkeLifetime:          models.ToPointer(28),
        IkeMode:              models.ToPointer(models.TunnelConfigIkeModeEnum_MAIN),
        IkeProposals:         []models.TunnelConfigIkeProposal{
            models.TunnelConfigIkeProposal{
                AuthAlgo:             models.ToPointer(models.TunnelConfigAuthAlgoEnum_SHA1),
                DhGroup:              models.ToPointer(models.TunnelConfigIkeDhGroupEnum_ENUM19),
                EncAlgo:              models.NewOptional(models.ToPointer(models.TunnelConfigEncAlgoEnum_AESGCM256)),
            },
            models.TunnelConfigIkeProposal{
                AuthAlgo:             models.ToPointer(models.TunnelConfigAuthAlgoEnum_SHA1),
                DhGroup:              models.ToPointer(models.TunnelConfigIkeDhGroupEnum_ENUM19),
                EncAlgo:              models.NewOptional(models.ToPointer(models.TunnelConfigEncAlgoEnum_AESGCM256)),
            },
            models.TunnelConfigIkeProposal{
                AuthAlgo:             models.ToPointer(models.TunnelConfigAuthAlgoEnum_SHA1),
                DhGroup:              models.ToPointer(models.TunnelConfigIkeDhGroupEnum_ENUM19),
                EncAlgo:              models.NewOptional(models.ToPointer(models.TunnelConfigEncAlgoEnum_AESGCM256)),
            },
        },
        IpsecLifetime:        models.ToPointer(88),
        Mode:                 models.ToPointer(models.TunnelConfigTunnelModeEnum_ACTIVESTANDBY),
        Version:              models.ToPointer(models.TunnelConfigVersionEnum_ENUM2),
    }

}
```

