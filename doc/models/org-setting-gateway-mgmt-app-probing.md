
# Org Setting Gateway Mgmt App Probing

Application probing settings for organization gateway management

## Structure

`OrgSettingGatewayMgmtAppProbing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | `[]string` | Optional | Application keys from [List Applications](../../doc/controllers/constants-definitions.md#list-applications) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingGatewayMgmtAppProbing := models.OrgSettingGatewayMgmtAppProbing{
        Apps:                 []string{
            "facebook",
        },
    }

}
```

