
# Stats Cluster Config Fabric Link Info

*This model accepts additional fields of type interface{}.*

## Structure

`StatsClusterConfigFabricLinkInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DataPlaneNotifiedStatus` | `*string` | Optional | - |
| `Interface` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `InternalStatus` | `*string` | Optional | - |
| `State` | `*string` | Optional | - |
| `Status` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "DataPlaneNotifiedStatus": "DataPlaneNotifiedStatus8",
  "Interface": [
    "Interface0"
  ],
  "InternalStatus": "InternalStatus2",
  "State": "State4",
  "Status": "Status0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

