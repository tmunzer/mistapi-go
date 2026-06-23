
# Edr Provider Enum

EDR provider associated with the NAC client. enum: `crowdstrike`, `sentinelone`

## Enumeration

`EdrProviderEnum`

## Fields

| Name |
|  --- |
| `crowdstrike` |
| `sentinelone` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    edrProvider := models.EdrProviderEnum_CROWDSTRIKE

}
```

