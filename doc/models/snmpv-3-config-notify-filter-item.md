
# Snmpv 3 Config Notify Filter Item

SNMPv3 notification filter profile

## Structure

`Snmpv3ConfigNotifyFilterItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Contents` | [`[]models.Snmpv3ConfigNotifyFilterItemContent`](../../doc/models/snmpv-3-config-notify-filter-item-content.md) | Optional | OID filter rules in an SNMPv3 notification filter profile |
| `ProfileName` | `*string` | Optional | Notification filter profile name |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpv3ConfigNotifyFilterItem := models.Snmpv3ConfigNotifyFilterItem{
        Contents:             []models.Snmpv3ConfigNotifyFilterItemContent{
            models.Snmpv3ConfigNotifyFilterItemContent{
                Include:              models.ToPointer(false),
                Oid:                  models.ToPointer("oid4"),
            },
            models.Snmpv3ConfigNotifyFilterItemContent{
                Include:              models.ToPointer(false),
                Oid:                  models.ToPointer("oid4"),
            },
        },
        ProfileName:          models.ToPointer("profile_name8"),
    }

}
```

