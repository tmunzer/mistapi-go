
# Org Device

Organization device identifier returned by the devices list

## Structure

`OrgDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Organization device MAC address |
| `Name` | `string` | Required | Display name of the organization device |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgDevice := models.OrgDevice{
        Mac:                  "mac8",
        Name:                 "name4",
    }

}
```

