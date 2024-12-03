
# Search Wxtag Apps Item

*This model accepts additional fields of type interface{}.*

## Structure

`SearchWxtagAppsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Group` | `string` | Required | - |
| `Key` | `string` | Required | - |
| `Name` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "group": "Emails",
  "key": "gmail",
  "name": "Gmail - web/app",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

