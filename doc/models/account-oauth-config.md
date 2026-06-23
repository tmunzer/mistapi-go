
# Account Oauth Config

OAuth linked apps (zoom/teams/intune) account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountOauthConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountId` | `string` | Required | Linked app(zoom/teams/intune) account id |
| `DiscardGuestInfo` | `*bool` | Optional | Optional, for Zoom/Teams. Whether to redact identifying information for call participants that are not part of the Zoom/Teams account identified by `account_id` |
| `MaxDailyApiRequests` | `*int` | Optional | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountOauthConfig := models.AccountOauthConfig{
        AccountId:            "iojzXIJWEuiD73ZvydOfg",
        DiscardGuestInfo:     models.ToPointer(false),
        MaxDailyApiRequests:  models.ToPointer(5000),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

