
# Gateway Compliance Major Version Properties

Version-compliance details for one gateway model

## Structure

`GatewayComplianceMajorVersionProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MajorCount` | `*int` | Optional | Number of gateways counted in this major-version compliance entry |
| `MajorVersion` | `*string` | Optional | Gateway software major version represented by this compliance entry |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayComplianceMajorVersionProperties := models.GatewayComplianceMajorVersionProperties{
        MajorCount:           models.ToPointer(46),
        MajorVersion:         models.ToPointer("19.4R2-S1.2"),
    }

}
```

