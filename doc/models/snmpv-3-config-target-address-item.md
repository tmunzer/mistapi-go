
# Snmpv 3 Config Target Address Item

SNMPv3 notification target address entry

## Structure

`Snmpv3ConfigTargetAddressItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `*string` | Optional | IP address or hostname of the SNMP target |
| `AddressMask` | `*string` | Optional | Mask applied to the SNMP target address |
| `Port` | `models.Optional[string]` | Optional | UDP port used by the SNMP target<br><br>**Default**: `"161"` |
| `TagList` | `*string` | Optional | Set of notification tags for this target address; use spaces between multiple tags |
| `TargetAddressName` | `*string` | Optional | Name of the SNMP target address entry |
| `TargetParameters` | `*string` | Optional | Target parameter profile referenced by this target address |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpv3ConfigTargetAddressItem := models.Snmpv3ConfigTargetAddressItem{
        Address:              models.ToPointer("10.11.0.2"),
        AddressMask:          models.ToPointer("255.255.255.0"),
        Port:                 models.NewOptional(models.ToPointer("161")),
        TagList:              models.ToPointer("tag_list6"),
        TargetAddressName:    models.ToPointer("target_address_name"),
    }

}
```

