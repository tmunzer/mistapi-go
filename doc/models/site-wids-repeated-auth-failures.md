
# Site Wids Repeated Auth Failures

## Structure

`SiteWidsRepeatedAuthFailures`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | window where a trigger will be detected and action to be taken (in seconds) |
| `Threshold` | `*int` | Optional | count of events to trigger |

## Example (as JSON)

```json
{
  "duration": 60,
  "threshold": 80
}
```

