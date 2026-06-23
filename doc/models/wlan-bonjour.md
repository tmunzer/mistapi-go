
# Wlan Bonjour

Bonjour gateway wlan settings

## Structure

`WlanBonjour`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalVlanIds` | [`*models.AdditionalVlanIds2`](../../doc/models/containers/additional-vlan-ids-2.md) | Optional | List or Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses |
| `Enabled` | `*bool` | Optional | Whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false<br><br>**Default**: `false` |
| `Services` | [`map[string]models.WlanBonjourServiceProperties`](../../doc/models/wlan-bonjour-service-properties.md) | Optional | What services are allowed.<br>Property key is the service name |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanBonjour := models.WlanBonjour{
        AdditionalVlanIds:    models.ToPointer(models.AdditionalVlanIds2Container.FromString("String1")),
        Enabled:              models.ToPointer(false),
        Services:             map[string]models.WlanBonjourServiceProperties{
            "airplay": models.WlanBonjourServiceProperties{
                DisableLocal:         models.ToPointer(false),
                RadiusGroups:         []string{
                    "teachers",
                },
                Scope:                models.ToPointer(models.WlanBonjourServicePropertiesScopeEnum_SAMEAP),
            },
        },
    }

}
```

