
# Org Site Sle Wired Result

## Structure

`OrgSiteSleWiredResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumClients` | `*float64` | Optional | - |
| `NumSwitches` | `*float64` | Optional | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `SwitchHealth` | `*float64` | Optional | - |
| `SwitchStc` | `*float64` | Optional | - |
| `SwitchThroughput` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "num_clients": 64.3,
  "num_switches": 63.2,
  "switch_health": 63.7,
  "switch_stc": 157.88,
  "switch_throughput": 231.14
}
```

