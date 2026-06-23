
# Org Setting Gateway Mgmt Host in Policy

Host-in access policy for a gateway management service

## Structure

`OrgSettingGatewayMgmtHostInPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Tenants` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingGatewayMgmtHostInPolicy := models.OrgSettingGatewayMgmtHostInPolicy{
        Tenants:              []string{
            "tenants1",
            "tenants2",
        },
    }

}
```

