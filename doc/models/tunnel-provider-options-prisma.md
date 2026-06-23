
# Tunnel Provider Options Prisma

Prisma Access provider options for tunnel auto provisioning

## Structure

`TunnelProviderOptionsPrisma`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ServiceAccountName` | `*string` | Optional | For prisma-ipsec, service account name to used for tunnel auto provisioning |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelProviderOptionsPrisma := models.TunnelProviderOptionsPrisma{
        ServiceAccountName:   models.ToPointer("sa1@1823425211"),
    }

}
```

