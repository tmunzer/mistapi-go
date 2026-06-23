
# Ssr Version

SSR firmware version available for upgrade

## Structure

`SsrVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Default` | `*bool` | Optional, Read-only | Whether this is the default SSR firmware version for its channel |
| `Package` | `string` | Required, Read-only | Firmware package name for this SSR version |
| `Tags` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Version` | `string` | Required, Read-only | Firmware version string available for SSR upgrades |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ssrVersion := models.SsrVersion{
        Default:              models.ToPointer(false),
        Package:              "package4",
        Tags:                 []string{
            "tags7",
        },
        Version:              "version8",
    }

}
```

