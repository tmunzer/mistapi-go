
# Response Nac Crl Files

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseNacCrlFiles`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.NacCrlFile`](../../doc/models/nac-crl-file.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "results": [
    {
      "created_time": 73.76,
      "id": "id6",
      "modified_time": 5.2,
      "name": "name6",
      "url": "url0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "created_time": 73.76,
      "id": "id6",
      "modified_time": 5.2,
      "name": "name6",
      "url": "url0",
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

