
# Zone Vertex

*This model accepts additional fields of type interface{}.*

## Structure

`ZoneVertex`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `X` | `float64` | Required | x in pixel |
| `Y` | `float64` | Required | y in pixel |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "x": 202.28,
  "y": 185.0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

