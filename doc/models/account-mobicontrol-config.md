
# Account Mobicontrol Config

MobiControl account credentials used for OAuth application linking

## Structure

`AccountMobicontrolConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `string` | Required | Customer account Client ID |
| `ClientSecret` | `string` | Required | Customer account Client Secret |
| `InstanceUrl` | `string` | Required | Customer account MobiControl instance URL |
| `Password` | `string` | Required | Customer account password instance URL |
| `Username` | `string` | Required | Login username used to authenticate to the MobiControl account |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountMobicontrolConfig := models.AccountMobicontrolConfig{
        ClientId:             "client_id4",
        ClientSecret:         "client_secret0",
        InstanceUrl:          "instance_url8",
        Password:             "password6",
        Username:             "username2",
    }

}
```

