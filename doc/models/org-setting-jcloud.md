
# Org Setting Jcloud

JCloud integration settings for this Mist organization

## Structure

`OrgSettingJcloud`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgApitoken` | `*string` | Optional | JCloud organization API token used by this Mist organization |
| `OrgApitokenName` | `*string` | Optional | Display name for the JCloud organization API token |
| `OrgId` | `*string` | Optional | JCloud organization identifier linked to this Mist organization |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingJcloud := models.OrgSettingJcloud{
        OrgApitoken:          models.ToPointer("org_apitoken0"),
        OrgApitokenName:      models.ToPointer("org_apitoken_name6"),
        OrgId:                models.ToPointer("org_id4"),
    }

}
```

