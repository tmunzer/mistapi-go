
# Recaptcha

*This model accepts additional fields of type interface{}.*

## Structure

`Recaptcha`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Flavor` | [`*models.RecaptchaFlavorEnum`](../../doc/models/recaptcha-flavor-enum.md) | Optional | flavor of the captcha. enum: `google`, `hcaptcha`<br>**Default**: `"google"` |
| `Required` | `*bool` | Optional | - |
| `Sitekey` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "flavor": "hcaptcha",
  "required": false,
  "sitekey": "sitekey0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

