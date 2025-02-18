
# Recover

*This model accepts additional fields of type interface{}.*

## Structure

`Recover`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Email` | `string` | Required | **Constraints**: *Maximum Length*: `64` |
| `Recaptcha` | `*string` | Optional | See  https://www.google.com/recaptcha/ |
| `RecaptchaFlavor` | [`*models.RecaptchaFlavorEnum`](../../doc/models/recaptcha-flavor-enum.md) | Optional | flavor of the captcha. enum: `google`, `hcaptcha`<br>**Default**: `"google"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "email": "test@mistsys.com",
  "recaptcha_flavor": "hcaptcha",
  "recaptcha": "recaptcha4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

