
# Wxtag Matching

*This model accepts additional fields of type interface{}.*

## Structure

`WxtagMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | - |
| `Since` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "5684dae9ac8b",
  "since": 1428939600,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

