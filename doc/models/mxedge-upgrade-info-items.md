
# Mxedge Upgrade Info Items

Available upgrade version for a Mist Edge package

## Structure

`MxedgeUpgradeInfoItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Default` | `*bool` | Optional | Whether this version is the default upgrade target for the package |
| `Distro` | `*string` | Optional | Linux distribution codename for the package version |
| `Package` | `string` | Required | Mist Edge service package name |
| `Version` | `string` | Required | Available version for this package |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeUpgradeInfoItems := models.MxedgeUpgradeInfoItems{
        Default:              models.ToPointer(false),
        Distro:               models.ToPointer("distro6"),
        Package:              "package0",
        Version:              "version4",
    }

}
```

