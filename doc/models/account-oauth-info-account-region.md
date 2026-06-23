
# Account Oauth Info Account Region

Prisma Access region bandwidth allocation for a linked OAuth account

## Structure

`AccountOauthInfoAccountRegion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AggregateRegion` | `*string` | Optional | Bandwidth Aggregate region for this region |
| `AllocatedBandwidth` | `*int` | Optional, Read-only | Allocated bandwidth for the region, in Mbps |
| `Name` | `*string` | Optional | Display name for this region |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountOauthInfoAccountRegion := models.AccountOauthInfoAccountRegion{
        AggregateRegion:      models.ToPointer("us-southwest"),
        AllocatedBandwidth:   models.ToPointer(1000),
        Name:                 models.ToPointer("US West"),
    }

}
```

