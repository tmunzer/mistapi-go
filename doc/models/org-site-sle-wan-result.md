
# Org Site Sle Wan Result

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSiteSleWanResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApplicationHealth` | `*float64` | Optional | - |
| `GatewayHealth` | `float64` | Required | - |
| `NumClients` | `*float64` | Optional | - |
| `NumGateways` | `*float64` | Optional | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `WanLinkHealth` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "gateway-health": 21.24,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "application_health": 253.26,
  "num_clients": 75.94,
  "num_gateways": 220.58,
  "wan-link-health": 141.8,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

