
# Org Setting Gateway Mgmt Host Out Policies

optional, for some of the host-out traffic, the path preference can be specified by default, ECMP will be used from all available route/path available services: dns/mist/ntp/pim

## Structure

`OrgSettingGatewayMgmtHostOutPolicies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | [`*models.GatewayMgmtHostOutPolicy`](../../doc/models/gateway-mgmt-host-out-policy.md) | Optional | - |
| `Ntp` | [`*models.GatewayMgmtHostOutPolicy`](../../doc/models/gateway-mgmt-host-out-policy.md) | Optional | - |
| `Syslog` | [`*models.GatewayMgmtHostOutPolicySyslog`](../../doc/models/gateway-mgmt-host-out-policy-syslog.md) | Optional | - |

## Example (as JSON)

```json
{
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
```

