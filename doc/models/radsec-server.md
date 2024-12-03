
# Radsec Server

*This model accepts additional fields of type interface{}.*

## Structure

`RadsecServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | - |
| `Port` | `*int` | Optional | **Constraints**: `>= 1`, `<= 65535` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "host": "1.1.1.1",
  "port": 1812,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

