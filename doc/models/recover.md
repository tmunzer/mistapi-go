
# Recover

## Structure

`Recover`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Email` | `string` | Required | **Constraints**: *Maximum Length*: `64` |
| `Recaptcha` | `*string` | Optional | see https://www.google.com/recaptcha/ |
| `RecaptchaFlavor` | [`*models.RecaptchaFlavorEnum`](../../doc/models/recaptcha-flavor-enum.md) | Optional | flavor of the captcha. enum: `google`, `hcaptcha`<br>**Default**: `"google"` |

## Example (as JSON)

```json
{
  "email": "test@mistsys.com",
  "recaptcha_flavor": "hcaptcha",
  "recaptcha": "recaptcha8"
}
```

