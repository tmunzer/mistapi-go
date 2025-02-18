
# Utils Zeroise Fips

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsZeroiseFips`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `string` | Required | FIPS zeroize password |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "password": "password2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

