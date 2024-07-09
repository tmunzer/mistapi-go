
# Org Setting Gateway Mgmt

## Structure

`OrgSettingGatewayMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppProbing` | [`*models.OrgSettingGatewayMgmtAppProbing`](../../doc/models/org-setting-gateway-mgmt-app-probing.md) | Optional | - |
| `AppUsage` | `*bool` | Optional | - |
| `HostOutPolicies` | [`*models.OrgSettingGatewayMgmtHostOutPolicies`](../../doc/models/org-setting-gateway-mgmt-host-out-policies.md) | Optional | optional, for some of the host-out traffic, the path preference can be specified by default, ECMP will be used from all available route/path available services: dns/mist/ntp |
| `OverlayIp` | [`*models.OrgSettingGatewayMgmtOverlayIp`](../../doc/models/org-setting-gateway-mgmt-overlay-ip.md) | Optional | - |

## Example (as JSON)

```json
{
  "app_probing": {
    "apps": [
      "apps8",
      "apps9",
      "apps0"
    ]
  },
  "app_usage": false,
  "host_out_policies": {
    "dns": {
      "path_preference": "path_preference8"
    },
    "mist": {
      "path_preference": "path_preference6"
    },
    "ntp": {
      "path_preference": "path_preference4"
    }
  },
  "overlay_ip": {
    "ip": "ip4",
    "node1_ip": "node1_ip2"
  }
}
```

