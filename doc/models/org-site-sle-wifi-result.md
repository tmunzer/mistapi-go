
# Org Site Sle Wifi Result

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ap-availability": 120.98,
  "ap-health": 153.22,
  "capacity": 65.86,
  "coverage": 44.84,
  "num_aps": 74.58,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

