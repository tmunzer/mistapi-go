
# Account Jamf Config

OAuth linked Jamf apps account details

## Structure

`AccountJamfConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `string` | Required | Customer account api client id. Required if `app_name`==`crowdstrike` |
| `ClientSecret` | `string` | Required | Customer account api client secret |
| `InstanceUrl` | `string` | Required | Customer account Jamf instance URL |
| `SmartgroupName` | `string` | Required | Smart group membership for determining compliance status |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountJamfConfig := models.AccountJamfConfig{
        ClientId:             "client_id4",
        ClientSecret:         "client_secret0",
        InstanceUrl:          "junipertest.jamfcloud.com",
        SmartgroupName:       "CompliantGroup1",
    }

}
```

