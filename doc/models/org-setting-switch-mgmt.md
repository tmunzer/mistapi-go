
# Org Setting Switch Mgmt

Organization-level switch management settings

## Structure

`OrgSettingSwitchMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApAffinityThreshold` | `*int` | Optional | If the field is set in both site/setting and org/setting, the value from site/setting will be used.<br><br>**Default**: `12` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingSwitchMgmt := models.OrgSettingSwitchMgmt{
        ApAffinityThreshold:  models.ToPointer(10),
    }

}
```

