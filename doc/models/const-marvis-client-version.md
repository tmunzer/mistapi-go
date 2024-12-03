
# Const Marvis Client Version

*This model accepts additional fields of type interface{}.*

## Structure

`ConstMarvisClientVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Label` | `*string` | Optional | - |
| `Notes` | `*string` | Optional | - |
| `Os` | `*string` | Optional | Client OS |
| `Url` | `*string` | Optional | Client download url |
| `Version` | `*string` | Optional | Client version |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "label": "default",
  "os": "windows",
  "url": "https://mobile.mist.com/installers/marvisclient/...",
  "version": "0.100.29",
  "notes": "notes8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

