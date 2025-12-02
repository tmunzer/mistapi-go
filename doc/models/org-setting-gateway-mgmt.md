
# Org Setting Gateway Mgmt

## Structure

`OrgSettingGatewayMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppProbing` | [`*models.OrgSettingGatewayMgmtAppProbing`](../../doc/models/org-setting-gateway-mgmt-app-probing.md) | Optional | - |
| `AppUsage` | `*bool` | Optional | consumes uplink bandwidth, requires WA license |
| `FipsEnabled` | `*bool` | Optional | **Default**: `false` |
| `HostInPolicies` | [`*models.OrgSettingGatewayMgmtHostInPolicies`](../../doc/models/org-setting-gateway-mgmt-host-in-policies.md) | Optional | - |
| `HostOutPolicies` | [`*models.OrgSettingGatewayMgmtHostOutPolicies`](../../doc/models/org-setting-gateway-mgmt-host-out-policies.md) | Optional | optional, for some of the host-out traffic, the path preference can be specified by default, ECMP will be used from all available route/path available services: dns/mist/ntp/pim |
| `OverlayIp` | [`*models.OrgSettingGatewayMgmtOverlayIp`](../../doc/models/org-setting-gateway-mgmt-overlay-ip.md) | Optional | - |

## Example (as JSON)

```json
{
  "fips_enabled": false,
  "app_probing": {
    "apps": [
      "apps8",
      "apps9",
      "apps0"
    ]
  },
  "app_usage": false,
  "host_in_policies": {
    "icmp": {
      "tenants": [
        "tenants3"
      ]
    },
    "snmp": {
      "tenants": [
        "tenants5"
      ]
    }
  },
  "host_out_policies": {
    "dns": {
      "path_preference": "path_preference8"
    },
    "ntp": {
      "path_preference": "path_preference4"
    },
    "syslog": {
      "path_preference": "path_preference2",
      "servers": [
        {
          "host": "host4",
          "path_preference": "path_preference8",
          "server_name": "server_name8"
        }
      ]
    }
  }
}
```

