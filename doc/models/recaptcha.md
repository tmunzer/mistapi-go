
# Recaptcha

## Structure

`Recaptcha`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Flavor` | [`*models.RecaptchaFlavorEnum`](../../doc/models/recaptcha-flavor-enum.md) | Optional | flavor of the captcha. enum: `google`, `hcaptcha`<br>**Default**: `"google"` |
| `Required` | `*bool` | Optional | - |
| `Sitekey` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "flavor": "hcaptcha",
  "required": false,
  "sitekey": "sitekey4"
}
```

