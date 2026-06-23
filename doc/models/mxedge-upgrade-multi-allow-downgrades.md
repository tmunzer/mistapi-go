
# Mxedge Upgrade Multi Allow Downgrades

Whether downgrade is allowed when running version is higher than expected version for each service

## Structure

`MxedgeUpgradeMultiAllowDowngrades`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mxagent` | `*bool` | Optional | Whether downgrades are allowed for the mxagent service<br><br>**Default**: `false` |
| `Mxdas` | `*bool` | Optional | Whether downgrades are allowed for the mxdas service<br><br>**Default**: `false` |
| `Mxocproxy` | `*bool` | Optional | Whether downgrades are allowed for the mxocproxy service<br><br>**Default**: `false` |
| `Radsecproxy` | `*bool` | Optional | Whether downgrades are allowed for the radsecproxy service<br><br>**Default**: `false` |
| `Tunterm` | `*bool` | Optional | Whether downgrades are allowed for the tunterm service<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeUpgradeMultiAllowDowngrades := models.MxedgeUpgradeMultiAllowDowngrades{
        Mxagent:              models.ToPointer(false),
        Mxdas:                models.ToPointer(false),
        Mxocproxy:            models.ToPointer(false),
        Radsecproxy:          models.ToPointer(false),
        Tunterm:              models.ToPointer(false),
    }

}
```

