
# Org Setting Gateway Mgmt

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingGatewayMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppProbing` | [`*models.OrgSettingGatewayMgmtAppProbing`](../../doc/models/org-setting-gateway-mgmt-app-probing.md) | Optional | - |
| `AppUsage` | `*bool` | Optional | consumes uplink bandwidth, requires WA license |
| `HostOutPolicies` | [`*models.OrgSettingGatewayMgmtHostOutPolicies`](../../doc/models/org-setting-gateway-mgmt-host-out-policies.md) | Optional | optional, for some of the host-out traffic, the path preference can be specified by default, ECMP will be used from all available route/path available services: dns/mist/ntp/pim |
| `OverlayIp` | [`*models.OrgSettingGatewayMgmtOverlayIp`](../../doc/models/org-setting-gateway-mgmt-overlay-ip.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "app_probing": {
    "apps": [
      "apps8",
      "apps9",
      "apps0"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "app_usage": false,
  "host_out_policies": {
    "dns": {
      "path_preference": "path_preference8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "mist": {
      "path_preference": "path_preference6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "ntp": {
      "path_preference": "path_preference4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "pim": {
      "path_preference": "path_preference0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "overlay_ip": {
    "ip": "ip4",
    "node1_ip": "node1_ip2",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

