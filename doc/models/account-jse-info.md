
# Account Jse Info

Linked Juniper Security Exchange account information

## Structure

`AccountJseInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | `*string` | Optional | JSE cloud hostname configured for the integration |
| `OrgNames` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Username` | `*string` | Optional | JSE integration username configured for the linked account |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountJseInfo := models.AccountJseInfo{
        CloudName:            models.ToPointer("devcentral.juniperclouds.net"),
        OrgNames:             []string{
            "org_names2",
        },
        Username:             models.ToPointer("john@abc.com"),
    }

}
```

