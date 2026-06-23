
# Account Crowdstrike Config

OAuth linked CrowdStrike apps account details

## Structure

`AccountCrowdstrikeConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `string` | Required | Customer account api client ID |
| `ClientSecret` | `string` | Required | Customer account api client Secret |
| `CustomerId` | `string` | Required | Customer id of an admin |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountCrowdstrikeConfig := models.AccountCrowdstrikeConfig{
        ClientId:             "client_id0",
        ClientSecret:         "client_secret6",
        CustomerId:           "customer_id6",
    }

}
```

