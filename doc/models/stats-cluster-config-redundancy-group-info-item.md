
# Stats Cluster Config Redundancy Group Info Item

*This model accepts additional fields of type interface{}.*

## Structure

`StatsClusterConfigRedundancyGroupInfoItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `MonitoringFailure` | `*string` | Optional | - |
| `Threshold` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "Id": 212,
  "MonitoringFailure": "MonitoringFailure2",
  "Threshold": 150,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

