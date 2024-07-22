
# Site Setting Critical Url Monitoring

you can define some URLs that's critical to site operaitons the latency will be captured and considered for site health

## Structure

`SiteSettingCriticalUrlMonitoring`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `true` |
| `Monitors` | [`[]models.SiteSettingCriticalUrlMonitoringMonitor`](../../doc/models/site-setting-critical-url-monitoring-monitor.md) | Optional | - |

## Example (as JSON)

```json
{
  "enabled": true,
  "monitors": [
    {
      "url": "url0",
      "vlan_id": "String5"
    },
    {
      "url": "url0",
      "vlan_id": "String5"
    },
    {
      "url": "url0",
      "vlan_id": "String5"
    }
  ]
}
```

