
# Protect Re Custom

Custom Protect RE ACL entry

## Structure

`ProtectReCustom`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | Matched dst port, "0" means any<br><br>**Default**: `"0"` |
| `Protocol` | [`*models.ProtectReCustomProtocolEnum`](../../doc/models/protect-re-custom-protocol-enum.md) | Optional | enum: `any`, `icmp`, `tcp`, `udp`<br><br>**Default**: `"any"` |
| `Subnets` | `[]string` | Optional | Source subnets matched by a custom Protect RE ACL |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    protectReCustom := models.ProtectReCustom{
        PortRange:            models.ToPointer("80,1035-1040"),
        Protocol:             models.ToPointer(models.ProtectReCustomProtocolEnum_ANY),
        Subnets:              []string{
            "subnets1",
            "subnets0",
            "subnets9",
        },
    }

}
```

