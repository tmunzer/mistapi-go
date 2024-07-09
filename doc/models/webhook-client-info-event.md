
# Webhook Client Info Event

## Structure

`WebhookClientInfoEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hostname` | `*string` | Optional | Hostname of client |
| `Ip` | `*string` | Optional | IP address of client |
| `Mac` | `*string` | Optional | the client’s mac |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | time at which IP address was assigned E.g. 1703003956 |

## Example (as JSON)

```json
{
  "hostname": "ervice.company.net",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ip": "ip0",
  "mac": "mac0"
}
```

