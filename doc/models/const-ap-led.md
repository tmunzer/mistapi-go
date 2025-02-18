
# Const Ap Led

*This model accepts additional fields of type interface{}.*

## Structure

`ConstApLed`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `string` | Required | - |
| `Description` | `string` | Required | - |
| `Key` | `string` | Required | - |
| `Name` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "code": "01",
  "description": "LED not working",
  "key": "LED_FAILURE",
  "name": "LED Failure",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

