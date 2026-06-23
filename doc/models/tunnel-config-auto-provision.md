
# Tunnel Config Auto Provision

Auto-provisioning configuration for the tunnel. This takes precedence over the `primary` and `secondary` nodes.

## Structure

`TunnelConfigAutoProvision`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Enable auto provisioning for the tunnel. If enabled, the `primary` and `secondary` nodes will be ignored. |
| `Latlng` | [`*models.TunnelConfigAutoProvisionLatLng`](../../doc/models/tunnel-config-auto-provision-lat-lng.md) | Optional | Geographic coordinate override for tunnel POP selection |
| `Primary` | [`*models.TunnelConfigAutoProvisionNode`](../../doc/models/tunnel-config-auto-provision-node.md) | Optional | Auto-provisioned tunnel endpoint settings |
| `Provider` | [`models.TunnelConfigAutoProvisionProviderEnum`](../../doc/models/tunnel-config-auto-provision-provider-enum.md) | Required | Tunnel provider used for automatic endpoint provisioning. enum: `jse-ipsec`, `zscaler-ipsec` |
| `Region` | `*string` | Optional | API override for POP selection in the case user wants to override the auto discovery of remote network location and force the tunnel to use the specified peer location. |
| `Secondary` | [`*models.TunnelConfigAutoProvisionNode`](../../doc/models/tunnel-config-auto-provision-node.md) | Optional | Auto-provisioned tunnel endpoint settings |
| `ServiceConnection` | `*string` | Optional | if `provider`==`prisma-ipsec`. By default, we'll use the location of the site to determine the optimal Remote Network location, optionally, service_connection can be considered, then we'll also consider this along with the site location. Define service_connection if the traffic is to be routed to a specific service connection. This field takes a service connection name that is configured in the Prisma cloud, Prisma Access Setup -> Service Connections. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfigAutoProvision := models.TunnelConfigAutoProvision{
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
        Region:               models.ToPointer("region4"),
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
        ServiceConnection:    models.ToPointer("Juniper-Lab-SC-1"),
    }

}
```

