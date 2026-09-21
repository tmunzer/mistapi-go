
# Switch Port Local Usage Mac Auth Protocol Enum

Only if `enable_mac_auth` ==`true`. When Mist NAC is enabled, this is forced to `pap`, unless the Org `mist_nac.enable_eap_md5_for_mab` setting is enabled: in that case `eap-md5` is kept and the port still performs MAB (mac-radius) but sends the request as EAP-MD5. enum: `eap-md5`, `eap-peap`, `pap`

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

