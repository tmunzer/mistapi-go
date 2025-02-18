
# Org Setting Gateway Mgmt Host Out Policies

optional, for some of the host-out traffic, the path preference can be specified by default, ECMP will be used from all available route/path available services: dns/mist/ntp/pim

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingGatewayMgmtHostOutPolicies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | [`*models.OrgSettingGatewayMgmtHostOutPoliciesDns`](../../doc/models/org-setting-gateway-mgmt-host-out-policies-dns.md) | Optional | - |
| `Mist` | [`*models.OrgSettingGatewayMgmtHostOutPoliciesMist`](../../doc/models/org-setting-gateway-mgmt-host-out-policies-mist.md) | Optional | - |
| `Ntp` | [`*models.OrgSettingGatewayMgmtHostOutPoliciesNtp`](../../doc/models/org-setting-gateway-mgmt-host-out-policies-ntp.md) | Optional | - |
| `Pim` | [`*models.OrgSettingGatewayMgmtHostOutPoliciesNtp`](../../doc/models/org-setting-gateway-mgmt-host-out-policies-ntp.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
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
}
```

