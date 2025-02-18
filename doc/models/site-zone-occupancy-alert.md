
# Site Zone Occupancy Alert

Zone Occupancy alert site settings

*This model accepts additional fields of type interface{}.*

## Structure

`SiteZoneOccupancyAlert`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EmailNotifiers` | `[]string` | Optional | List of email addresses to send email notifications when the alert threshold is reached |
| `Enabled` | `*bool` | Optional | Indicate whether zone occupancy alert is enabled for the site<br>**Default**: `false` |
| `Threshold` | `*int` | Optional | Sending zone-occupancy-alert webhook message only if a zone stays non-compliant (i.e. actual occupancy > occupancy_limit) for a minimum duration specified in the threshold, in minutes<br>**Default**: `5`<br>**Constraints**: `>= 0`, `<= 30` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "email_notifiers": [
    "foo@juniper.net",
    "bar@juniper.net"
  ],
  "enabled": false,
  "threshold": 5,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

