
# Site Zone Occupancy Alert

Zone Occupancy alert site settings

## Structure

`SiteZoneOccupancyAlert`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EmailNotifiers` | `[]string` | Optional | list of email addresses to send email notifications when the alert threshold is reached |
| `Enabled` | `*bool` | Optional | indicate whether zone occupancy alert is enabled for the site<br>**Default**: `false` |
| `Threshold` | `*int` | Optional | sending zone-occupancy-alert webhook message only if a zone stays non-compliant (i.e. actual occupancy > occupancy_limit) for a minimum duration specified in the threshold, in minutes<br>**Default**: `5`<br>**Constraints**: `>= 0`, `<= 30` |

## Example (as JSON)

```json
{
  "email_notifiers": [
    "foo@juniper.net",
    "bar@juniper.net"
  ],
  "enabled": false,
  "threshold": 5
}
```

