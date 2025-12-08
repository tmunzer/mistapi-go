
# Webhook Client Info Event

## Structure

`WebhookClientInfoEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hostname` | `*string` | Optional | Hostname of client |
| `Ip` | `*string` | Optional | IP address of client |
| `Mac` | `*string` | Optional | client's MAC Address |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |

## Example (as JSON)

```json
{
  "hostname": "service.company.net",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ip": "ip8",
  "mac": "mac8"
}
```

