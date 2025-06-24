
# Admin Invite

*This model accepts additional fields of type interface{}.*

## Structure

`AdminInvite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountOnly` | `*bool` | Optional | Skip creating initial setup if true<br><br>**Default**: `false` |
| `AllowMist` | `*bool` | Optional | Whether to allow Mist to look at this org<br><br>**Default**: `false` |
| `City` | `*string` | Optional | City of registering user |
| `Country` | `*string` | Optional | Country/region name or ISO code of registering user |
| `Email` | `string` | Required | **Constraints**: *Maximum Length*: `64` |
| `FirstName` | `string` | Required | - |
| `InviteCode` | `*string` | Optional | Required initially |
| `LastName` | `string` | Required | - |
| `OrgName` | `string` | Required | - |
| `Password` | `string` | Required | - |
| `Recaptcha` | `string` | Required | reCAPTCHA , see https://www.google.com/recaptcha/ |
| `RecaptchaFlavor` | [`*models.RecaptchaFlavorEnum`](../../doc/models/recaptcha-flavor-enum.md) | Optional | flavor of the captcha. enum: `google`, `hcaptcha`<br><br>**Default**: `"google"` |
| `RefererInviteToken` | `*string` | Optional | Invite token to apply after account creation |
| `ReturnTo` | `*string` | Optional | URL the user should be redirected back to |
| `State` | `*string` | Optional | State name or ISO code of registering user, optional (depends on country/region) |
| `StreetAddress` | `*string` | Optional | Street address of registering user |
| `StreetAddress2` | `*string` | Optional | Street address 2 of registering user |
| `Zipcode` | `*string` | Optional | zipcode of registering user |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "account_only": false,
  "allow_mist": false,
  "city": "Cupertino",
  "country": "United States",
  "email": "test@mistsys.com",
  "first_name": "John",
  "invite_code": "MISTROCKS",
  "last_name": "Smith",
  "org_name": "Smith LLC",
  "password": "foryoureyesonly",
  "recaptcha": "recaptcha2",
  "recaptcha_flavor": "hcaptcha",
  "referer_invite_token": "Dm2gtT8dwMeM4Bc2E8FLIaA96VHOjPat",
  "return_to": "https://mist.zendesk.com/hc/quickstart.pdf",
  "state": "CA",
  "street_address": "1601 S De Anza Blvd Ste 248",
  "street_address 2": "1601 S De Anza Blvd Ste 248",
  "zipcode": "95014",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

