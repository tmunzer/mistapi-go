
# Device Version Item

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceVersionItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Model` | `string` | Required | Device model (as seen in the device stats) |
| `Tag` | `*string` | Optional | Annotation, stable / beta / alpha. Or it can be empty or nothing which is likely a dev build |
| `Version` | `string` | Required | Firmware version |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "model": "model8",
  "tag": "tag6",
  "version": "version6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

