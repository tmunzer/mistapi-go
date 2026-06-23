
# Admin Invite

Administrator invitation and initial registration payload

*This model accepts additional fields of type interface{}.*

## Structure

`AdminInvite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountOnly` | `*bool` | Optional | Skip creating initial setup if true<br><br>**Default**: `false` |
| `AllowMist` | `*bool` | Optional | Whether to allow Mist to look at this org<br><br>**Default**: `false` |
| `City` | `*string` | Optional | Registration city for the admin user |
| `Country` | `*string` | Optional | Registration country or region for the admin user, as a name or ISO code |
| `Email` | `string` | Required | Registration email address for the admin user<br><br>**Constraints**: *Maximum Length*: `64` |
| `FirstName` | `string` | Required | Given name for the registering admin user |
| `InviteCode` | `*string` | Optional | Invite code used to authorize new admin registration |
| `LastName` | `string` | Required | Family name for the registering admin user |
| `OrgName` | `string` | Required | Organization name supplied during initial admin registration |
| `Password` | `string` | Required | Credential password for the registering admin account |
| `Recaptcha` | `string` | Required | CAPTCHA verification token submitted during admin registration |
| `RecaptchaFlavor` | [`*models.RecaptchaFlavorEnum`](../../doc/models/recaptcha-flavor-enum.md) | Optional | CAPTCHA provider flavor. enum: `google`, `hcaptcha`<br><br>**Default**: `"google"` |
| `RefererInviteToken` | `*string` | Optional | Invite token to apply after account creation |
| `ReturnTo` | `*string` | Optional | URL the user should be redirected back to |
| `State` | `*string` | Optional | Registration state or province for the admin user, optional depending on country or region |
| `StreetAddress` | `*string` | Optional | Street address of registering user |
| `StreetAddress2` | `*string` | Optional | Street address 2 of registering user |
| `Zipcode` | `*string` | Optional | Postal code for the registering admin user |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    adminInvite := models.AdminInvite{
        AccountOnly:          models.ToPointer(false),
        AllowMist:            models.ToPointer(false),
        City:                 models.ToPointer("Cupertino"),
        Country:              models.ToPointer("United States"),
        Email:                "test@mistsys.com",
        FirstName:            "John",
        InviteCode:           models.ToPointer("MISTROCKS"),
        LastName:             "Smith",
        OrgName:              "Smith LLC",
        Password:             "foryoureyesonly",
        Recaptcha:            "recaptcha8",
        RecaptchaFlavor:      models.ToPointer(models.RecaptchaFlavorEnum_HCAPTCHA),
        RefererInviteToken:   models.ToPointer("Dm2gtT8dwMeM4Bc2E8FLIaA96VHOjPat"),
        ReturnTo:             models.ToPointer("https://mist.zendesk.com/hc/quickstart.pdf"),
        State:                models.ToPointer("CA"),
        StreetAddress:        models.ToPointer("1601 S De Anza Blvd Ste 248"),
        StreetAddress2:       models.ToPointer("1601 S De Anza Blvd Ste 248"),
        Zipcode:              models.ToPointer("95014"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

