
# Remote Syslog Content

## Structure

`RemoteSyslogContent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Facility` | [`*models.RemoteSyslogFacilityEnum`](../../doc/models/remote-syslog-facility-enum.md) | Optional | enum: `any`, `authorization`, `change-log`, `config`, `conflict-log`, `daemon`, `dfc`, `external`, `firewall`, `ftp`, `interactive-commands`, `kernel`, `ntp`, `pfe`, `security`, `user`<br>**Default**: `"any"` |
| `Severity` | [`*models.RemoteSyslogSeverityEnum`](../../doc/models/remote-syslog-severity-enum.md) | Optional | enum: `alert`, `any`, `critical`, `emergency`, `error`, `info`, `notice`, `warning`<br>**Default**: `"any"` |

## Example (as JSON)

```json
{
  "facility": "config",
  "severity": "any"
}
```

