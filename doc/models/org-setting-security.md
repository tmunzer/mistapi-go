
# Org Setting Security

Organization security controls for local SSH and FIPS zeroize access

## Structure

`OrgSettingSecurity`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableLocalSsh` | `*bool` | Optional | Whether to disable local SSH (by default, local SSH is enabled with allow_mist in Org is enabled |
| `FipsZeroizePassword` | `*string` | Optional | password required to zeroize devices (FIPS) on site level |
| `LimitSshAccess` | `*bool` | Optional | Whether to allow certain SSH keys to SSH into the AP (see Site:Setting)<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingSecurity := models.OrgSettingSecurity{
        DisableLocalSsh:      models.ToPointer(false),
        FipsZeroizePassword:  models.ToPointer("NUKETHESITE"),
        LimitSshAccess:       models.ToPointer(false),
    }

}
```

