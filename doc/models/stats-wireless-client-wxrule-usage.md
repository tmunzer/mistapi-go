
# Stats Wireless Client Wxrule Usage

WxLAN rule usage counter for one wireless tag

## Structure

`StatsWirelessClientWxruleUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TagId` | `*uuid.UUID` | Optional | WxLAN tag identifier for this rule usage entry |
| `Usage` | `*int` | Optional | Count recorded for this WxLAN rule usage entry |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsWirelessClientWxruleUsage := models.StatsWirelessClientWxruleUsage{
        TagId:                models.ToPointer(uuid.MustParse("00000dea-0000-0000-0000-000000000000")),
        Usage:                models.ToPointer(68),
    }

}
```

