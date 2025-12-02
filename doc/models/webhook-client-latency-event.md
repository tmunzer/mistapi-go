
# Webhook Client Latency Event

## Structure

`WebhookClientLatencyEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AvgAuth` | `*float64` | Optional | - |
| `AvgDhcp` | `*float64` | Optional | - |
| `AvgDns` | `*float64` | Optional | - |
| `MaxAuth` | `*float64` | Optional | - |
| `MaxDhcp` | `*float64` | Optional | - |
| `MaxDns` | `*float64` | Optional | - |
| `MinAuth` | `*float64` | Optional | - |
| `MinDhcp` | `*float64` | Optional | - |
| `MinDns` | `*float64` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |

## Example (as JSON)

```json
{
  "avg_auth": 0.17170219,
  "avg_dhcp": 0.017828934,
  "avg_dns": 0.024532124,
  "max_auth": 0.18170219,
  "max_dhcp": 0.027828934,
  "max_dns": 0.022532124,
  "min_auth": 0.16050219,
  "min_dhcp": 0.015828934,
  "min_dns": 0.029532124,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6"
}
```

