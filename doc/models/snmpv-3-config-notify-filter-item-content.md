
# Snmpv 3 Config Notify Filter Item Content

OID filter rule for an SNMPv3 notification profile

## Structure

`Snmpv3ConfigNotifyFilterItemContent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Include` | `*bool` | Optional | Whether the matching OID subtree is included |
| `Oid` | `*string` | Optional | Matched OID subtree for this notification filter rule |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpv3ConfigNotifyFilterItemContent := models.Snmpv3ConfigNotifyFilterItemContent{
        Include:              models.ToPointer(false),
        Oid:                  models.ToPointer("1.3.6.1.4.1"),
    }

}
```

