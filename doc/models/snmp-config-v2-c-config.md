
# Snmp Config V2 C Config

SNMPv2c community configuration entry

## Structure

`SnmpConfigV2cConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Authorization` | `*string` | Optional | Access level for the SNMPv2c community |
| `ClientListName` | `*string` | Optional | SNMP client list name referenced by this community |
| `CommunityName` | `*string` | Optional | SNMPv2c community string name |
| `View` | `*string` | Optional | SNMP view name that must be defined in the views list |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpConfigV2cConfig := models.SnmpConfigV2cConfig{
        Authorization:        models.ToPointer("read-only"),
        ClientListName:       models.ToPointer("clist-1"),
        CommunityName:        models.ToPointer("abc123"),
        View:                 models.ToPointer("all"),
    }

}
```

