
# Wxlan Tag Spec

Traffic match specification used by a WxLAN tag

## Structure

`WxlanTagSpec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | Matched destination port, "0" means any<br><br>**Default**: `"0"` |
| `Protocol` | `*string` | Optional | tcp / udp / icmp / gre / any / ":protocol_number", `protocol_number` is between 1-254<br><br>**Default**: `"any"` |
| `Subnets` | `[]string` | Optional | Matched destination subnets and/or IP addresses |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wxlanTagSpec := models.WxlanTagSpec{
        PortRange:            models.ToPointer("0"),
        Protocol:             models.ToPointer("any"),
        Subnets:              []string{
            "0.0.0.0/0",
        },
    }

}
```

