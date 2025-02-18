
# Search Wan Usage

*This model accepts additional fields of type interface{}.*

## Structure

`SearchWanUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.WanUsages`](../../doc/models/wan-usages.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "results": [
    {
      "mac": "mac0",
      "path_type": "path_type8",
      "path_weight": 242,
      "peer_mac": "peer_mac6",
      "peer_port_id": "peer_port_id4",
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

