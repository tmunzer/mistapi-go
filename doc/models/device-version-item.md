
# Device Version Item

## Structure

`DeviceVersionItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Model` | `string` | Required | Device model (as seen in the device stats) |
| `Tag` | `*string` | Optional | Annotation, stable / beta / alpha. Or it can be empty or nothing which is likely a dev build |
| `Version` | `string` | Required | Firmware version |

## Example (as JSON)

```json
{
  "model": "model8",
  "tag": "tag6",
  "version": "version6"
}
```

