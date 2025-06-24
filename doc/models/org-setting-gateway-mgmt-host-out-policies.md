
# Org Setting Gateway Mgmt Host Out Policies

optional, for some of the host-out traffic, the path preference can be specified by default, ECMP will be used from all available route/path available services: dns/mist/ntp/pim

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingGatewayMgmtHostOutPolicies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | [`*models.GatewayMgmtHostOutPolicy`](../../doc/models/gateway-mgmt-host-out-policy.md) | Optional | - |
| `Ntp` | [`*models.GatewayMgmtHostOutPolicy`](../../doc/models/gateway-mgmt-host-out-policy.md) | Optional | - |
| `Syslog` | [`*models.GatewayMgmtHostOutPolicySyslog`](../../doc/models/gateway-mgmt-host-out-policy-syslog.md) | Optional | - |
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
  "ntp": {
    "path_preference": "path_preference4",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "syslog": {
    "path_preference": "path_preference2",
    "servers": [
      {
        "host": "host4",
        "name": "name2",
        "path_preference": "path_preference8",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    ],
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

