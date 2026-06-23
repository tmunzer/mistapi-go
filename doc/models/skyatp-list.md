
# Skyatp List

Sky ATP SecIntel allowlist or blocklist entries

*This model accepts additional fields of type interface{}.*

## Structure

`SkyatpList`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Domains` | [`[]models.SkyatpListDomain`](../../doc/models/skyatp-list-domain.md) | Optional | Domain entries in a Sky ATP SecIntel list |
| `Ip` | [`[]models.SkyatpListIp`](../../doc/models/skyatp-list-ip.md) | Optional | IP address entries in a Sky ATP SecIntel list |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    skyatpList := models.SkyatpList{
        Domains:              []models.SkyatpListDomain{
            models.SkyatpListDomain{
                Comment:              models.ToPointer("comment6"),
                Value:                "value2",
            },
            models.SkyatpListDomain{
                Comment:              models.ToPointer("comment6"),
                Value:                "value2",
            },
            models.SkyatpListDomain{
                Comment:              models.ToPointer("comment6"),
                Value:                "value2",
            },
        },
        Ip:                   []models.SkyatpListIp{
            models.SkyatpListIp{
                Comment:              models.ToPointer("comment8"),
                Value:                "value6",
            },
            models.SkyatpListIp{
                Comment:              models.ToPointer("comment8"),
                Value:                "value6",
            },
            models.SkyatpListIp{
                Comment:              models.ToPointer("comment8"),
                Value:                "value6",
            },
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

