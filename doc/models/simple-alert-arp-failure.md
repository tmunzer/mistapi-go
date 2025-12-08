
# Simple Alert Arp Failure

## Structure

`SimpleAlertArpFailure`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientCount` | `*int` | Optional | **Default**: `10` |
| `Duration` | `*int` | Optional | failing within minutes<br><br>**Default**: `20`<br><br>**Constraints**: `>= 5`, `<= 60` |
| `IncidentCount` | `*int` | Optional | **Default**: `10` |

## Example (as JSON)

```json
{
  "client_count": 10,
  "duration": 20,
  "incident_count": 10
}
```

