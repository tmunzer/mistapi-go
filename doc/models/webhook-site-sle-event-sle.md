
# Webhook Site Sle Event Sle

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookSiteSleEventSle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApAvailability` | `*float64` | Optional | - |
| `SuccessfulConnect` | `*float64` | Optional | - |
| `TimeToConnect` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap-availability": 0.6,
  "successful-connect": 0.7,
  "time-to-connect": 0.9,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

