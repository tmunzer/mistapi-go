
# Alarm Template Rule

## Structure

`AlarmTemplateRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Delivery` | [`*models.Delivery`](../../doc/models/delivery.md) | Optional | Delivery object to configure the alarm delivery |
| `Enabled` | `*bool` | Optional | - |

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
    "to_site_admins": false
  },
  "enabled": false
}
```

