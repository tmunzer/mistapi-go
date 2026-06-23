
# Root Password String

Temporary root password response for device ZTP recovery

## Structure

`RootPasswordString`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RootPassword` | `string` | Required | Temporary root password returned for ZTP recovery<br><br>**Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rootPasswordString := models.RootPasswordString{
        RootPassword:         "root_password0",
    }

}
```

