
# Const License Type

*This model accepts additional fields of type interface{}.*

## Structure

`ConstLicenseType`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | - |
| `Includes` | `[]string` | Optional | - |
| `Key` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "description": "Wired Assurance 12",
  "includes": [
    "sub_ex12a",
    "sub_ex12p"
  ],
  "key": "sub_ex12",
  "name": "SUB-EX12",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

