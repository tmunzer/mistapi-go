
# Jsi Pbn Item

PBN (Problem Bug Notification) advisory item

## Structure

`JsiPbnItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BugType` | [`*models.JsiPbnBugTypeEnum`](../../doc/models/jsi-pbn-bug-type-enum.md) | Optional | Type of the bug. enum: `Day-1`, `Regression` |
| `CustomerRisk` | `*string` | Optional | Customer impact risk level for the PBN advisory |
| `FixedIn` | `*string` | Optional | Release in which the issue was fixed |
| `Id` | `*string` | Optional | Unique PBN advisory identifier from Juniper Support Insights |
| `IntroducedIn` | `*string` | Optional | Release where the PBN issue was introduced |
| `Models` | `[]string` | Optional | Device models affected by the PBN issue |
| `ProductFamily` | `[]string` | Optional | Product families affected by the PBN issue |
| `ReleaseNotes` | `*string` | Optional | Release notes for this PBN |
| `Restoration` | `*string` | Optional | Steps recommended to restore service or recover from the PBN issue |
| `Title` | `*string` | Optional | Summary title for the PBN issue |
| `UpdatedDate` | `*int` | Optional | Time when the PBN advisory was last updated |
| `Versions` | `[]string` | Optional | Software versions affected by the PBN issue |
| `Workaround` | `*string` | Optional | Mitigation or workaround guidance for the PBN issue |
| `WorkaroundProvided` | `*string` | Optional | Indicator of whether workaround guidance is available |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    jsiPbnItem := models.JsiPbnItem{
        BugType:              models.ToPointer(models.JsiPbnBugTypeEnum_DAY1),
        CustomerRisk:         models.ToPointer("customer_risk4"),
        FixedIn:              models.ToPointer("fixed_in8"),
        Id:                   models.ToPointer("1403338"),
        IntroducedIn:         models.ToPointer("introduced_in2"),
    }

}
```

