
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
  "num_clients": 245.6,
  "num_switches": 244.5,
  "switch_health": 245.0,
  "switch_stc": 83.18,
  "switch_throughput": 99.56
}
```

