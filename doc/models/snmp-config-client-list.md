
# Snmp Config Client List

SNMP client allowlist definition

## Structure

`SnmpConfigClientList`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientListName` | `*string` | Optional | Name of the SNMP client list |
| `Clients` | `[]string` | Optional | SNMP client IP addresses or CIDR ranges allowed by a client list |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpConfigClientList := models.SnmpConfigClientList{
        ClientListName:       models.ToPointer("clist-1"),
        Clients:              []string{
            "clients2",
        },
    }

}
```

