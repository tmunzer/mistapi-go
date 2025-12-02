
# Org Setting Gateway Mgmt Host in Policies

## Structure

`OrgSettingGatewayMgmtHostInPolicies`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Icmp` | [`*models.OrgSettingGatewayMgmtHostInPolicy`](../../doc/models/org-setting-gateway-mgmt-host-in-policy.md) | Optional | - |
| `Snmp` | [`*models.OrgSettingGatewayMgmtHostInPolicy`](../../doc/models/org-setting-gateway-mgmt-host-in-policy.md) | Optional | - |

## Example (as JSON)

```json
{
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
}
```

