
# Switch Port Local Usage Mac Auth Protocol Enum

Only if `enable_mac_auth` ==`true`. This type is ignored if mist_nac is enabled. enum: `eap-md5`, `eap-peap`, `pap`

## Enumeration

`SwitchPortLocalUsageMacAuthProtocolEnum`

## Fields

| Name |
|  --- |
| `eap-md5` |
| `eap-peap` |
| `pap` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchPortLocalUsageMacAuthProtocol := models.SwitchPortLocalUsageMacAuthProtocolEnum_EAPPEAP

}
```

