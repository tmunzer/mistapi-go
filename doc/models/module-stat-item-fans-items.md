
# Module Stat Item Fans Items

*This model accepts additional fields of type interface{}.*

## Structure

`ModuleStatItemFansItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Airflow` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Status` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "airflow": "out",
  "name": "Fan 0",
  "status": "ok",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

