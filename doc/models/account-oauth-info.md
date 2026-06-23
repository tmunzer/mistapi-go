
# Account Oauth Info

OAuth-linked application account status and authorization details

## Structure

`AccountOauthInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accounts` | [`[]models.AccountOauthInfoAccount`](../../doc/models/account-oauth-info-account.md) | Required | List of linked account details |
| `AuthorizationUrl` | `*string` | Optional, Read-only | OAuth authorization URL to open when linking an account |
| `Linked` | `bool` | Required, Read-only | Whether at least one account is linked for this OAuth application |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountOauthInfo := models.AccountOauthInfo{
        Accounts:             []models.AccountOauthInfoAccount{
            models.AccountOauthInfoAccount{
                AccountId:            models.ToPointer("iojzXIJWEuiD73ZvydOfg"),
                AutoProbeSubnet:      models.ToPointer("11.0.0.0/8"),
                ClientId:             models.ToPointer("client_id2"),
                CloudName:            models.ToPointer("Tapi.sase.paloaltonetworks.com"),
                Company:              models.ToPointer("Test Company1 Ltd"),
                EnableProbe:          models.ToPointer(false),
                Error:                models.ToPointer("OAuth token refresh failed, please re-link your account"),
                Errors:               []string{
                    "OAuth token refresh failed, please re-link your account",
                    "API daily rate limit reached for your account",
                },
                KeyId:                models.ToPointer("L72frZcK3JvrZc"),
                LastStatus:           models.ToPointer("failed"),
                LastSync:             models.ToPointer(int64(1665465339000)),
                LinkedBy:             models.ToPointer("Testname1"),
                LinkedTimestamp:      models.ToPointer(float64(1665465339000)),
                MaxDailyApiRequests:  models.ToPointer(5000),
                Name:                 models.ToPointer("Test Compay1 Ltd"),
                Region:               models.ToPointer("americas"),
                ServiceAccountName:   models.ToPointer("Corp SA"),
                SmartgroupName:       models.ToPointer("CompliantGroup1"),
                TsgId:                models.ToPointer("189953456"),
                WebhookAuthType:      models.ToPointer("Basic"),
                WebhookPassword:      models.ToPointer("password_1234"),
                WebhookSecret:        models.ToPointer("secret-value"),
                WebhookToken:         models.ToPointer("token-value"),
                WebhookUrl:           models.ToPointer("https://websync.nac-staging.mistsys.com/v1/S_org-8dcbe9005/ae9dee49-69e7-4710-a114-5b827a777738/crowdstrike/edr"),
                WebhookUsername:      models.ToPointer("username_1234"),
                ZdxOrgId:             models.ToPointer("123456"),
            },
        },
        AuthorizationUrl:     models.ToPointer("authorization_url0"),
        Linked:               false,
    }

}
```

