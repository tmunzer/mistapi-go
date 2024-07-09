
# Org Setting Gateway Mgmt Host Out Policies

optional, for some of the host-out traffic, the path preference can be specified by default, ECMP will be used from all available route/path available services: dns/mist/ntp

## Structure

`OrgSettingGatewayMgmtHostOutPolicies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | [`*models.OrgSettingGatewayMgmtHostOutPoliciesDns`](../../doc/models/org-setting-gateway-mgmt-host-out-policies-dns.md) | Optional | - |
| `Mist` | [`*models.OrgSettingGatewayMgmtHostOutPoliciesMist`](../../doc/models/org-setting-gateway-mgmt-host-out-policies-mist.md) | Optional | - |
| `Ntp` | [`*models.OrgSettingGatewayMgmtHostOutPoliciesNtp`](../../doc/models/org-setting-gateway-mgmt-host-out-policies-ntp.md) | Optional | - |

## Example (as JSON)

```json
{
  "dns": {
    "path_preference": "path_preference8"
  },
  "mist": {
    "path_preference": "path_preference6"
  },
  "ntp": {
    "path_preference": "path_preference4"
  }
}
```

