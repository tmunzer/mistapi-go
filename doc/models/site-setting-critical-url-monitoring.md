
# Site Setting Critical Url Monitoring

you can define some URLs that's critical to site operaitons the latency will be captured and considered for site health

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingCriticalUrlMonitoring`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `true` |
| `Monitors` | [`[]models.SiteSettingCriticalUrlMonitoringMonitor`](../../doc/models/site-setting-critical-url-monitoring-monitor.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": true,
  "monitors": [
    {
      "url": "url0",
      "vlan_id": "String5",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "url": "url0",
      "vlan_id": "String5",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "url": "url0",
      "vlan_id": "String5",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

