
# Mxedge Upgrade Version

Version to upgrade for each service, `current` / `latest` / `default` / specific version (e.g. `2.5.100`).\nIgnored if distro upgrade, `tunterm`, `radsecproxy`, `mxagent`, `mxocproxy`, `mxdas` or `mxnacedge`

## Structure

`MxedgeUpgradeVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mxagent` | `string` | Required | Target version for the mxagent service<br><br>**Default**: `"current"` |
| `Mxdas` | `*string` | Optional | Target version for the mxdas service<br><br>**Default**: `"current"` |
| `Mxocproxy` | `*string` | Optional | Target version for the mxocproxy service<br><br>**Default**: `"current"` |
| `Radsecproxy` | `*string` | Optional | Target version for the radsecproxy service<br><br>**Default**: `"current"` |
| `Tunterm` | `string` | Required | Target version for the tunterm service<br><br>**Default**: `"current"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeUpgradeVersion := models.MxedgeUpgradeVersion{
        Mxagent:              "current",
        Mxdas:                models.ToPointer("current"),
        Mxocproxy:            models.ToPointer("current"),
        Radsecproxy:          models.ToPointer("current"),
        Tunterm:              "current",
    }

}
```

