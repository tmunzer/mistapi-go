
# Org Site Sle Wifi Result

## Structure

`OrgSiteSleWifiResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApAvailability` | `*float64` | Optional | - |
| `ApHealth` | `*float64` | Optional | - |
| `Capacity` | `*float64` | Optional | - |
| `Coverage` | `*float64` | Optional | - |
| `NumAps` | `*float64` | Optional | - |
| `NumClients` | `*float64` | Optional | - |
| `Roaming` | `*float64` | Optional | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `SuccessfulConnect` | `*float64` | Optional | - |
| `Throughput` | `*float64` | Optional | - |
| `TimeToConnect` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ap-availability": 129.46,
  "ap-health": 161.7,
  "capacity": 181.66,
  "coverage": 53.32,
  "num_aps": 66.1
}
```

