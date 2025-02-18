
# Ha Cluster Config

*This model accepts additional fields of type interface{}.*

## Structure

`HaClusterConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableAutoConfig` | `*bool` | Optional | If the device is claimed |
| `Managed` | `*bool` | Optional | If the device is adopted |
| `Nodes` | [`[]models.HaClusterConfigNode`](../../doc/models/ha-cluster-config-node.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "43e9c864-a7e4-4310-8031-d9817d2c5a43",
  "disable_auto_config": false,
  "managed": false,
  "nodes": [
    {
      "mac": "mac0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "mac": "mac0",
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

