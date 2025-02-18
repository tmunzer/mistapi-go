
# Alarm Template Rule

*This model accepts additional fields of type interface{}.*

## Structure

`AlarmTemplateRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Delivery` | [`*models.Delivery`](../../doc/models/delivery.md) | Optional | Delivery object to configure the alarm delivery |
| `Enabled` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "delivery": {
    "additional_emails": [
      "additional_emails9",
      "additional_emails0",
      "additional_emails1"
    ],
    "enabled": false,
    "to_org_admins": false,
    "to_site_admins": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

