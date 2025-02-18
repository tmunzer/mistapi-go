
# Module Stat Item Temperatures Item

*This model accepts additional fields of type interface{}.*

## Structure

`ModuleStatItemTemperaturesItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Celsius` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Status` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "celsius": 45.0,
  "name": "CPU",
  "status": "ok",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

