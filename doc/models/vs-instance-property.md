
# Vs Instance Property

*This model accepts additional fields of type interface{}.*

## Structure

`VsInstanceProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Networks` | `[]string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "networks": [
    "networks0",
    "networks1",
    "networks2"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

