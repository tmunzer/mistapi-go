
# Site Setting Critical Url Monitoring Monitor

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingCriticalUrlMonitoringMonitor`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Url` | `*string` | Optional | - |
| `VlanId` | [`*models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "url": "http://50.1.3.5:8080",
  "vlan_id": "String1",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

