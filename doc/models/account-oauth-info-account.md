
# Account Oauth Info Account

OAuth linked apps account info

## Structure

`AccountOauthInfoAccount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountId` | `*string` | Optional, Read-only | Linked app account id |
| `AutoProbeSubnet` | `*string` | Optional, Read-only | For Prisma accounts only, tunnel auto probe subnet |
| `ClientId` | `*string` | Optional, Read-only | Customer account Client ID |
| `CloudName` | `*string` | Optional, Read-only | Name of the company whose account mist has subscribed to |
| `Company` | `*string` | Optional, Read-only | Name of the company whose account mist has subscribed to |
| `EnableProbe` | `*bool` | Optional, Read-only | For Prisma accounts only, tunnel probe enable/disable |
| `Error` | `*string` | Optional, Read-only | This error is provided when the account fails to fetch token/data |
| `Errors` | `[]string` | Optional, Read-only | Read-only OAuth account error messages |
| `InstanceUrl` | `*string` | Optional, Read-only | Customer account instance URL |
| `KeyId` | `*string` | Optional | For ZDX Account only, Customer account API key ID |
| `LastStatus` | `*string` | Optional, Read-only | Is the last data pull for account is successful or not |
| `LastSync` | `*int64` | Optional, Read-only | Last data pull timestamp, background jobs that pull account data |
| `LinkedBy` | `*string` | Optional, Read-only | First name of the user who linked the account |
| `LinkedTimestamp` | `*float64` | Optional, Read-only | Timestamp when this third-party account was linked |
| `MaxDailyApiRequests` | `*int` | Optional, Read-only | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ |
| `Name` | `*string` | Optional, Read-only | Display name of the linked third-party account or company |
| `Password` | `*string` | Optional, Read-only | Customer account password instance URL |
| `Region` | `*string` | Optional, Read-only | For Prisma accounts only |
| `Regions` | [`map[string]models.AccountOauthInfoAccountRegion`](../../doc/models/account-oauth-info-account-region.md) | Optional | For Prisma accounts only, property key is the region name. Regions with allocated bandwidth |
| `ServiceAccountName` | `*string` | Optional, Read-only | For Prisma accounts only |
| `ServiceConnections` | [`map[string]models.AccountOauthInfoAccountServiceConnection`](../../doc/models/account-oauth-info-account-service-connection.md) | Optional | For Prisma accounts only, property key is the service connection name |
| `SmartgroupName` | `*string` | Optional, Read-only | Smart group membership for determining compliance status |
| `TsgId` | `*string` | Optional, Read-only | For Prisma accounts only, Prisma Tenant Service Group id |
| `Username` | `*string` | Optional, Read-only | Login name configured for the linked third-party account |
| `WebhookAuthType` | `*string` | Optional | For Crowdstrike, JAMF, SentinelOne and VMWare accounts only |
| `WebhookEnabled` | `*bool` | Optional | For Crowdstrike, JAMF, SentinelOne and VMWare accounts only |
| `WebhookPassword` | `*string` | Optional | For VMWare accounts only |
| `WebhookSecret` | `*string` | Optional | For Crowdstrike accounts only |
| `WebhookToken` | `*string` | Optional | For JAMF and SentinelOne accounts only |
| `WebhookUrl` | `*string` | Optional | For Crowdstrike, JAMF, SentinelOne and VMWare accounts only |
| `WebhookUsername` | `*string` | Optional | For VMWare accounts only |
| `ZdxOrgId` | `*string` | Optional | For ZDX Account only, ZDX organization id |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountOauthInfoAccount := models.AccountOauthInfoAccount{
        AccountId:            models.ToPointer("iojzXIJWEuiD73ZvydOfg"),
        AutoProbeSubnet:      models.ToPointer("11.0.0.0/8"),
        ClientId:             models.ToPointer("client_id8"),
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
    }

}
```

