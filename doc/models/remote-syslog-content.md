
# Remote Syslog Content

## Structure

`RemoteSyslogContent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Facility` | [`*models.RemoteSyslogFacilityEnum`](../../doc/models/remote-syslog-facility-enum.md) | Optional | **Default**: `"any"` |
| `Severity` | [`*models.RemoteSyslogSeverityEnum`](../../doc/models/remote-syslog-severity-enum.md) | Optional | **Default**: `"any"` |

## Example (as JSON)

```json
{
  "facility": "config",
  "severity": "any"
}
```

