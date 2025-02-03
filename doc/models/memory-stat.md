
# Memory Stat

Memory usage stat (for virtual chassis, memory usage of master RE)

*This model accepts additional fields of type interface{}.*

## Structure

`MemoryStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Usage` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "usage": 57.54,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

