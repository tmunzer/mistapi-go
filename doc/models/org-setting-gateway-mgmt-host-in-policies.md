
# Org Setting Gateway Mgmt Host in Policies

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingGatewayMgmtHostInPolicies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Icmp` | [`*models.OrgSettingGatewayMgmtHostInPolicy`](../../doc/models/org-setting-gateway-mgmt-host-in-policy.md) | Optional | - |
| `Snmp` | [`*models.OrgSettingGatewayMgmtHostInPolicy`](../../doc/models/org-setting-gateway-mgmt-host-in-policy.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "icmp": {
    "name": "name2",
    "tenants": [
      "tenants3"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "snmp": {
    "name": "name4",
    "tenants": [
      "tenants5"
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

