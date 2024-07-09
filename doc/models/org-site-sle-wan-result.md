
# Org Site Sle Wan Result

## Structure

`OrgSiteSleWanResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApplicationHealth` | `*float64` | Optional | - |
| `GatewayHealth` | `*float64` | Optional | - |
| `NumClients` | `*float64` | Optional | - |
| `NumGateways` | `*float64` | Optional | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `WanLinkHealth` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "application_health": 158.84,
  "gateway-health": 73.18,
  "num_clients": 237.52,
  "num_gateways": 126.16,
  "wan-link-health": 47.38
}
```

