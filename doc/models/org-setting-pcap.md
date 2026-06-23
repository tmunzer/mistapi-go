
# Org Setting Pcap

Packet capture export settings for the organization

## Structure

`OrgSettingPcap`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bucket` | `*string` | Optional | Storage bucket name used for organization packet capture files |
| `MaxPktLen` | `*int` | Optional | Maximum length of non-management packets to capture, in bytes<br><br>**Default**: `128`<br><br>**Constraints**: `<= 128` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingPcap := models.OrgSettingPcap{
        Bucket:               models.ToPointer("myorg_pcap"),
        MaxPktLen:            models.ToPointer(128),
    }

}
```

