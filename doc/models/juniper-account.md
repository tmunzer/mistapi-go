
# Juniper Account

Linked Juniper account available to the organization

## Structure

`JuniperAccount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LinkedBy` | `*string` | Optional, Read-only | User who linked this Juniper account |
| `Name` | `*string` | Optional, Read-only | Display name of the linked Juniper account |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    juniperAccount := models.JuniperAccount{
        LinkedBy:             models.ToPointer("John Smith (john@abccorp.com)"),
        Name:                 models.ToPointer("ABC Corp"),
    }

}
```

