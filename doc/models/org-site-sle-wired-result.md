
# Org Site Sle Wired Result

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSiteSleWiredResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumClients` | `*float64` | Optional | - |
| `NumSwitches` | `*float64` | Optional | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `SwitchBandwidth` | `*float64` | Optional | - |
| `SwitchHealth` | `float64` | Required | - |
| `SwitchThroughput` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "switch-health": 214.08,
  "num_clients": 64.3,
  "num_switches": 63.2,
  "switch-bandwidth": 66.5,
  "switch-throughput": 152.98,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

