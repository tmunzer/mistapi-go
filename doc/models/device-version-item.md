
# Device Version Item

## Structure

`DeviceVersionItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Model` | `string` | Required | Device model (as seen in the device stats) |
| `Tag` | `*string` | Optional | annotation, stable / beta / alpha. Or it can be empty or nothing which is likely a dev build |
| `Version` | `string` | Required | firmware version |

## Example (as JSON)

```json
{
  "model": "model2",
  "tag": "tag2",
  "version": "version0"
}
```

