
# Org Setting Jcloud Ra

JCloud Routing Assurance connexion

## Structure

`OrgSettingJcloudRa`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgApitoken` | `*string` | Optional | JCloud Routing Assurance Org Token |
| `OrgApitokenName` | `*string` | Optional | JCloud Routing Assurance Org Token Name |
| `OrgId` | `*string` | Optional | JCloud Routing Assurance Org ID |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingJcloudRa := models.OrgSettingJcloudRa{
        OrgApitoken:          models.ToPointer("org_apitoken4"),
        OrgApitokenName:      models.ToPointer("org_apitoken_name0"),
        OrgId:                models.ToPointer("org_id0"),
    }

}
```

