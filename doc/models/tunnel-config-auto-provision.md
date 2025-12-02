
# Tunnel Config Auto Provision

Auto Provisioning configuration for the tunne. This takes precedence over the `primary` and `secondary` nodes.

## Structure

`TunnelConfigAutoProvision`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Enable auto provisioning for the tunnel. If enabled, the `primary` and `secondary` nodes will be ignored. |
| `Latlng` | [`*models.TunnelConfigAutoProvisionLatLng`](../../doc/models/tunnel-config-auto-provision-lat-lng.md) | Optional | API override for POP selection |
| `Primary` | [`*models.TunnelConfigAutoProvisionNode`](../../doc/models/tunnel-config-auto-provision-node.md) | Optional | - |
| `Provider` | [`models.TunnelConfigAutoProvisionProviderEnum`](../../doc/models/tunnel-config-auto-provision-provider-enum.md) | Required | enum: `jse-ipsec`, `zscaler-ipsec` |
| `Region` | `*string` | Optional | API override for POP selection in the case user wants to override the auto discovery of remote network location and force the tunnel to use the specified peer location. |
| `Secondary` | [`*models.TunnelConfigAutoProvisionNode`](../../doc/models/tunnel-config-auto-provision-node.md) | Optional | - |
| `ServiceConnection` | `*string` | Optional | if `provider`==`prisma-ipsec`. By default, we'll use the location of the site to determine the optimal Remote Network location, optionally, service_connection can be considered, then we'll also consider this along with the site location. Define service_connection if the traffic is to be routed to a specific service connection. This field takes a service connection name that is configured in the Prisma cloud, Prisma Access Setup -> Service Connections. |

## Example (as JSON)

```json
{
  "provider": "jse-ipsec",
  "service_connection": "Juniper-Lab-SC-1",
  "enabled": false,
  "latlng": {
    "lat": 144.64,
    "lng": 22.82
  },
  "primary": {
    "probe_ips": [
      "probe_ips7",
      "probe_ips8",
      "probe_ips9"
    ],
    "wan_names": [
      "wan_names8"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "region": "region4",
  "secondary": {
    "probe_ips": [
      "probe_ips9",
      "probe_ips0",
      "probe_ips1"
    ],
    "wan_names": [
      "wan_names0"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  }
}
```

