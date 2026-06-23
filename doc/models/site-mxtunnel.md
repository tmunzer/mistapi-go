
# Site Mxtunnel

Site Mist Tunnel configuration for tunneling AP user VLANs to Mist Edge tunnel peers

## Structure

`SiteMxtunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalMxtunnels` | [`map[string]models.SiteMxtunnelAdditionalMxtunnel`](../../doc/models/site-mxtunnel-additional-mxtunnel.md) | Optional | Additional named Mist Tunnel definitions keyed by tunnel name |
| `ApSubnets` | `[]string` | Optional | List of subnets where we allow AP to establish Mist Tunnels from |
| `AutoPreemption` | [`*models.AutoPreemption`](../../doc/models/auto-preemption.md) | Optional | Schedule to preempt AP tunnels that are not connected to their preferred peer |
| `Clusters` | [`[]models.SiteMxtunnelCluster`](../../doc/models/site-mxtunnel-cluster.md) | Optional | Tunnel peer cluster definitions APs use for the site Mist Tunnel |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Enabled` | `*bool` | Optional | Whether site Mist Tunnel tunneling is enabled |
| `ForSite` | `*bool` | Optional, Read-only | Whether this Mist Tunnel configuration is scoped to a site |
| `HelloInterval` | `*int` | Optional | In seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries<br><br>**Default**: `60`<br><br>**Constraints**: `>= 1`, `<= 300` |
| `HelloRetries` | `*int` | Optional | Number of missed hello heartbeats before an AP tries another tunnel peer<br><br>**Default**: `7`<br><br>**Constraints**: `>= 2`, `<= 30` |
| `Hosts` | `[]string` | Optional | Tunnel peer hostnames or IP addresses reachable from APs |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Mtu` | `*int` | Optional | 0 to enable MTU, 552-1500 to start MTU with a lower MTU<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 1500` |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Protocol` | [`*models.MxtunnelProtocolEnum`](../../doc/models/mxtunnel-protocol-enum.md) | Optional | Encapsulation protocol used for Mist Tunnel traffic. enum: `ip`, `udp`<br><br>**Default**: `"udp"` |
| `Radsec` | [`*models.SiteMxtunnelRadsec`](../../doc/models/site-mxtunnel-radsec.md) | Optional | RadSec proxy settings for a site Mist Tunnel |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `VlanIds` | `[]int` | Optional | List of VLAN IDs carried by this Mist Tunnel |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    siteMxtunnel := models.SiteMxtunnel{
        AdditionalMxtunnels:  map[string]models.SiteMxtunnelAdditionalMxtunnel{
            "key0": models.SiteMxtunnelAdditionalMxtunnel{
                Clusters:             []models.SiteMxtunnelCluster{
                    models.SiteMxtunnelCluster{
                        Name:                 models.ToPointer("name6"),
                        TuntermHosts:         []string{
                            "tunterm_hosts6",
                        },
                    },
                    models.SiteMxtunnelCluster{
                        Name:                 models.ToPointer("name6"),
                        TuntermHosts:         []string{
                            "tunterm_hosts6",
                        },
                    },
                },
                HelloInterval:        models.ToPointer(38),
                HelloRetries:         models.ToPointer(30),
                Protocol:             models.ToPointer(models.SiteMxtunnelProtocolEnum_IP),
                VlanIds:              []int{
                    142,
                    143,
                },
            },
            "key1": models.SiteMxtunnelAdditionalMxtunnel{
                Clusters:             []models.SiteMxtunnelCluster{
                    models.SiteMxtunnelCluster{
                        Name:                 models.ToPointer("name6"),
                        TuntermHosts:         []string{
                            "tunterm_hosts6",
                        },
                    },
                    models.SiteMxtunnelCluster{
                        Name:                 models.ToPointer("name6"),
                        TuntermHosts:         []string{
                            "tunterm_hosts6",
                        },
                    },
                },
                HelloInterval:        models.ToPointer(38),
                HelloRetries:         models.ToPointer(30),
                Protocol:             models.ToPointer(models.SiteMxtunnelProtocolEnum_IP),
                VlanIds:              []int{
                    142,
                    143,
                },
            },
            "key2": models.SiteMxtunnelAdditionalMxtunnel{
                Clusters:             []models.SiteMxtunnelCluster{
                    models.SiteMxtunnelCluster{
                        Name:                 models.ToPointer("name6"),
                        TuntermHosts:         []string{
                            "tunterm_hosts6",
                        },
                    },
                    models.SiteMxtunnelCluster{
                        Name:                 models.ToPointer("name6"),
                        TuntermHosts:         []string{
                            "tunterm_hosts6",
                        },
                    },
                },
                HelloInterval:        models.ToPointer(38),
                HelloRetries:         models.ToPointer(30),
                Protocol:             models.ToPointer(models.SiteMxtunnelProtocolEnum_IP),
                VlanIds:              []int{
                    142,
                    143,
                },
            },
        },
        ApSubnets:            []string{
            "ap_subnets2",
            "ap_subnets3",
            "ap_subnets4",
        },
        AutoPreemption:       models.ToPointer(models.AutoPreemption{
            DayOfWeek:            models.ToPointer(models.DayOfWeekEnum_MON),
            Enabled:              models.ToPointer(false),
            TimeOfDay:            models.ToPointer("time_of_day4"),
        }),
        Clusters:             []models.SiteMxtunnelCluster{
            models.SiteMxtunnelCluster{
                Name:                 models.ToPointer("name6"),
                TuntermHosts:         []string{
                    "tunterm_hosts6",
                },
            },
        },
        CreatedTime:          models.ToPointer(float64(199.14)),
        HelloInterval:        models.ToPointer(60),
        HelloRetries:         models.ToPointer(3),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Mtu:                  models.ToPointer(1100),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Protocol:             models.ToPointer(models.MxtunnelProtocolEnum_UDP),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

