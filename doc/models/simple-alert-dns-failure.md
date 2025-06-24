
# Simple Alert Dns Failure

*This model accepts additional fields of type interface{}.*

## Structure

`SimpleAlertDnsFailure`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientCount` | `*int` | Optional | **Default**: `20` |
| `Duration` | `*int` | Optional | failing within minutes<br><br>**Default**: `10`<br><br>**Constraints**: `>= 5`, `<= 60` |
| `IncidentCount` | `*int` | Optional | **Default**: `30` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "client_count": 20,
  "duration": 10,
  "incident_count": 30,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

