
# Alarm Template

Alarm Template

## Structure

`AlarmTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `Delivery` | [`models.Delivery`](../../doc/models/delivery.md) | Required | Delivery object to configure the alarm delivery |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | Some string to name the alarm template |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Rules` | [`map[string]models.AlarmTemplateRule`](../../doc/models/alarm-template-rule.md) | Required | Alarm Rules object to configure the individual alarm keys/types. Property key is the alarm name. |

## Example (as JSON)

```json
{
  "created_time": 1711001253.0,
  "delivery": {
    "enabled": true,
    "to_org_admins": true,
    "to_site_admins": false,
    "additional_emails": [
      "additional_emails9",
      "additional_emails0",
      "additional_emails1"
    ]
  },
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 1711038102.0,
  "name": "default",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rules": {
    "ap_offline": {
      "delivery": {
        "additional_emails": [
          "string"
        ],
        "enabled": true,
        "to_org_admins": true,
        "to_site_admins": true
      },
      "enabled": true
    },
    "bad_cable": {
      "delivery": {
        "additional_emails": [
          "string"
        ],
        "enabled": true,
        "to_org_admins": true,
        "to_site_admins": true
      },
      "enabled": true
    }
  }
}
```

