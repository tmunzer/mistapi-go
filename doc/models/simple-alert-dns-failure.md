
# Simple Alert Dns Failure

## Structure

`SimpleAlertDnsFailure`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientCount` | `*int` | Optional | **Default**: `20` |
| `Duration` | `*int` | Optional | failing within minutes<br>**Default**: `10`<br>**Constraints**: `>= 5`, `<= 60` |
| `IncidentCount` | `*int` | Optional | **Default**: `30` |

## Example (as JSON)

```json
{
  "client_count": 20,
  "duration": 10,
  "incident_count": 30
}
```

