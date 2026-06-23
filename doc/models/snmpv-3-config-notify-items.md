
# Snmpv 3 Config Notify Items

SNMPv3 notification definition for traps or informs

## Structure

`Snmpv3ConfigNotifyItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | Identifier for this SNMPv3 notification definition |
| `Tag` | `*string` | Optional | Notification tag used to select target addresses |
| `Type` | [`*models.Snmpv3ConfigNotifyTypeEnum`](../../doc/models/snmpv-3-config-notify-type-enum.md) | Optional | Delivery mode for this SNMPv3 notification, such as trap or inform. enum: `inform`, `trap` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpv3ConfigNotifyItems := models.Snmpv3ConfigNotifyItems{
        Name:                 models.ToPointer("name2"),
        Tag:                  models.ToPointer("tag6"),
        Type:                 models.ToPointer(models.Snmpv3ConfigNotifyTypeEnum_INFORM),
    }

}
```

