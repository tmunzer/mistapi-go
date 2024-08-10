
# Delivery

Delivery object to configure the alarm delivery

## Structure

`Delivery`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalEmails` | `[]string` | Optional | List of additional email string to deliver the alarms via emails |
| `Enabled` | `bool` | Required | Whether to enable the alarm delivery via emails or not |
| `ToOrgAdmins` | `*bool` | Optional | Whether to deliver the alarms via emails to Org admins or not |
| `ToSiteAdmins` | `*bool` | Optional | Whether to deliver the alarms via emails to Site admins or not |

## Example (as JSON)

```json
{
  "enabled": true,
  "to_org_admins": true,
  "to_site_admins": false,
  "additional_emails": [
    "additional_emails9",
    "additional_emails0",
    "additional_emails1"
  ]
}
```

