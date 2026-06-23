
# Stats Mxedge Inactive Vlan Strs

Inactive wired/L2TP VLANs. Entries can be individual VLANs or ranges.

## Structure

`StatsMxedgeInactiveVlanStrs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `L2tp` | `[]string` | Optional | Inactive L2TP VLANs. Entries can be individual VLANs or ranges. |
| `Wired` | `[]string` | Optional | Inactive wired VLANs. Entries can be individual VLANs or ranges. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgeInactiveVlanStrs := models.StatsMxedgeInactiveVlanStrs{
        L2tp:                 []string{
            "l2tp3",
            "l2tp2",
            "l2tp1",
        },
        Wired:                []string{
            "100",
            "102-106",
        },
    }

}
```

