
# Site Wids Repeated Auth Failures

*This model accepts additional fields of type interface{}.*

## Structure

`SiteWidsRepeatedAuthFailures`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | window where a trigger will be detected and action to be taken (in seconds) |
| `Threshold` | `*int` | Optional | count of events to trigger |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 60,
  "threshold": 80,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

