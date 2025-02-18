
# Simple Alert Arp Failure

*This model accepts additional fields of type interface{}.*

## Structure

`SimpleAlertArpFailure`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientCount` | `*int` | Optional | **Default**: `10` |
| `Duration` | `*int` | Optional | failing within minutes<br>**Default**: `20`<br>**Constraints**: `>= 5`, `<= 60` |
| `IncidentCount` | `*int` | Optional | **Default**: `10` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "client_count": 10,
  "duration": 20,
  "incident_count": 10,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

