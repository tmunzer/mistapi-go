
# Snmp Vacm Security to Group

## Structure

`SnmpVacmSecurityToGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Content` | [`[]models.SnmpVacmSecurityToGroupContentItem`](../../doc/models/snmp-vacm-security-to-group-content-item.md) | Optional | - |
| `SecurityModel` | [`*models.SnmpVacmSecurityModelEnum`](../../doc/models/snmp-vacm-security-model-enum.md) | Optional | enum: `usm`, `v1`, `v2c` |

## Example (as JSON)

```json
{
  "content": [
    {
      "group": "group2",
      "security_name": "security_name6"
    },
    {
      "group": "group2",
      "security_name": "security_name6"
    }
  ],
  "security_model": "v1"
}
```

