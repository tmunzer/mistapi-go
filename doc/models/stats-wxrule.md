
# Stats Wxrule

WxLAN rule usage statistics for a site

## Structure

`StatsWxrule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`models.StatsWxruleActionEnum`](../../doc/models/stats-wxrule-action-enum.md) | Required | Allow or block behavior enforced by this WxLAN rule. enum: `allow`, `block` |
| `ClientMac` | `[]string` | Required | Wireless client MAC addresses matching a WxLAN rule |
| `DstAllowWxtags` | `[]uuid.UUID` | Required | Destination WxLAN tag identifiers explicitly allowed by a WxLAN rule |
| `DstDenyWxtags` | `[]uuid.UUID` | Required | Destination WxLAN tag identifiers denied by a WxLAN rule |
| `DstWxtags` | `[]uuid.UUID` | Required | Destination WxLAN tag identifiers matched by a WxLAN rule |
| `Name` | `string` | Required | Display name of the WxLAN rule |
| `Order` | `int` | Required | Rule evaluation order for WxLAN matching |
| `SrcWxtags` | `[]uuid.UUID` | Required | Source WxLAN tag identifiers matched by a WxLAN rule |
| `Usage` | [`map[string]models.StatsWxruleUsageProperties`](../../doc/models/stats-wxrule-usage-properties.md) | Required | Flow counts keyed by destination WxLAN tag identifier |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsWxrule := models.StatsWxrule{
        Action:               models.StatsWxruleActionEnum_ALLOW,
        ClientMac:            []string{
            "3bbbf819bb6f",
            "bd96cbc4910f",
        },
        DstAllowWxtags:       []uuid.UUID{
            uuid.MustParse("fff34466-eec0-3756-6765-381c728a6037"),
            uuid.MustParse("eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"),
        },
        DstDenyWxtags:        []uuid.UUID{
            uuid.MustParse("aaa34466-eec0-3756-6765-381c728a6037"),
            uuid.MustParse("bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"),
        },
        DstWxtags:            []uuid.UUID{
            uuid.MustParse("d4134466-eec0-3756-6765-381c728a6037"),
            uuid.MustParse("1a42c7b0-d1d0-5a30-f349-e35fa43dc3b3"),
        },
        Name:                 "Guest",
        Order:                1,
        SrcWxtags:            []uuid.UUID{
            uuid.MustParse("8bfc2490-d726-3587-038d-cb2e71bd2330"),
            uuid.MustParse("3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"),
        },
        Usage:                map[string]models.StatsWxruleUsageProperties{
            "1a42c7b0-d1d0-5a30-f349-e35fa43dc3b3": models.StatsWxruleUsageProperties{
                NumFlows:             models.ToPointer(60),
            },
            "d4134466-eec0-3756-6765-381c728a6037": models.StatsWxruleUsageProperties{
                NumFlows:             models.ToPointer(60),
            },
        },
    }

}
```

