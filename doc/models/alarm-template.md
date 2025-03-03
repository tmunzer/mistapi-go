
# Alarm Template

Alarm Template

*This model accepts additional fields of type interface{}.*

## Structure

`AlarmTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Delivery` | [`models.Delivery`](../../doc/models/delivery.md) | Required | Delivery object to configure the alarm delivery |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Some string to name the alarm template |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Rules` | [`map[string]models.AlarmTemplateRule`](../../doc/models/alarm-template-rule.md) | Required | Alarm Rules object to configure the individual alarm keys/types. Property key is the alarm name. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "delivery": {
    "enabled": true,
    "to_org_admins": true,
    "to_site_admins": false,
    "additional_emails": [
      "additional_emails9",
      "additional_emails0",
      "additional_emails1"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
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
        "to_site_admins": true,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "enabled": true,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "bad_cable": {
      "delivery": {
        "additional_emails": [
          "string"
        ],
        "enabled": true,
        "to_org_admins": true,
        "to_site_admins": true,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "enabled": true,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "created_time": 2.88,
  "modified_time": 76.08,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

