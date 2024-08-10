
# Webhook Device Updowns Event

## Structure

`WebhookDeviceUpdownsEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | - |
| `ApName` | `string` | Required | - |
| `ForSite` | `*bool` | Optional | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `SiteName` | `string` | Required | - |
| `Timestamp` | `float64` | Required | - |
| `Type` | `string` | Required | - |

## Example (as JSON)

```json
{
  "ap": "ap6",
  "ap_name": "ap_name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "site_name": "site_name0",
  "timestamp": 16.04,
  "type": "type2",
  "for_site": false
}
```

