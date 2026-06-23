
# Network Tenant

Tenant address entry for a network

## Structure

`NetworkTenant`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Addresses` | `[]string` | Optional | IP addresses or subnets assigned to a network tenant |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    networkTenant := models.NetworkTenant{
        Addresses:            []string{
            "addresses0",
            "addresses1",
        },
    }

}
```

