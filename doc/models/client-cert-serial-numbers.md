
# Client Cert Serial Numbers

*This model accepts additional fields of type interface{}.*

## Structure

`ClientCertSerialNumbers`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SerialNumbers` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "serial_numbers": [
    "13 00 13 03 23 EE D5 84 01"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

