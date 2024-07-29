
# Admin Invite

## Structure

`AdminInvite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountOnly` | `*bool` | Optional | skip creating initial setup if true<br>**Default**: `false` |
| `AllowMist` | `*bool` | Optional | whether to allow Mist to look at this org<br>**Default**: `false` |
| `City` | `*string` | Optional | city of registering user |
| `Country` | `*string` | Optional | country/region of registering user |
| `Email` | `string` | Required | **Constraints**: *Maximum Length*: `64` |
| `FirstName` | `string` | Required | - |
| `InviteCode` | `*string` | Optional | required initially |
| `LastName` | `string` | Required | - |
| `OrgName` | `string` | Required | - |
| `Password` | `string` | Required | - |
| `Recaptcha` | `string` | Required | reCAPTCHA , see https://www.google.com/recaptcha/ |
| `RecaptchaFlavor` | [`*models.RecaptchaFlavorEnum`](../../doc/models/recaptcha-flavor-enum.md) | Optional | flavor of the captcha. enum: `google`, `hcaptcha`<br>**Default**: `"google"` |
| `RefererInviteToken` | `*string` | Optional | the invite token to apply after account creation |
| `ReturnTo` | `*string` | Optional | the url the user should be redirected back to |
| `State` | `*string` | Optional | state of registering user, optional (depends on country/region) |
| `StreetAddress` | `*string` | Optional | street address of registering user |
| `StreetAddress2` | `*string` | Optional | street address 2 of registering user |
| `Zipcode` | `*string` | Optional | zipcode of registering user |

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
  "recaptcha": "recaptcha0",
  "recaptcha_flavor": "hcaptcha",
  "referer_invite_token": "Dm2gtT8dwMeM4Bc2E8FLIaA96VHOjPat",
  "return_to": "http://mist.zendesk.com/hc/quickstart.pdf",
  "state": "California",
  "street_address": "1601 S De Anza Blvd Ste 248",
  "street_address 2": "1601 S De Anza Blvd Ste 248",
  "zipcode": "95014"
}
```

