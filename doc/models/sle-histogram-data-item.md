
# Sle Histogram Data Item

*This model accepts additional fields of type interface{}.*

## Structure

`SleHistogramDataItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Range` | `[]float64` | Optional | - |
| `Value` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "range": [
    63.23,
    63.22
  ],
  "value": 25.36,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

