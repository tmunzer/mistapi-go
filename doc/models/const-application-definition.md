
# Const Application Definition

*This model accepts additional fields of type interface{}.*

## Structure

`ConstApplicationDefinition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppId` | `*bool` | Optional | - |
| `AppImageUrl` | `*string` | Optional | - |
| `AppProbe` | `*bool` | Optional | - |
| `Category` | `*string` | Optional | - |
| `Group` | `*string` | Optional | - |
| `Key` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `SignatureBased` | `*bool` | Optional | - |
| `SsrAppId` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "app_image_url": "\"\"",
  "category": "FileSharing",
  "group": "File Sharing",
  "key": "dropbox",
  "name": "Dropbox",
  "app_id": false,
  "app_probe": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

