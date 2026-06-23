
# Ap Template Wifi

Wi-Fi behavior settings applied by an AP template

## Structure

`ApTemplateWifi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CiscoEnabled` | `*bool` | Optional | Whether Cisco-specific Wi-Fi compatibility behavior is enabled |
| `Disable11k` | `*bool` | Optional | Whether 802.11k neighbor reports are disabled<br><br>**Default**: `false` |
| `DisableRadiosWhenPowerConstrained` | `*bool` | Optional | Whether AP radios are disabled when power is constrained |
| `EnableArpSpoof` | `*bool` | Optional | Whether ARP spoofing checks are enabled when proxy ARP is used |
| `EnableSharedRadioScanning` | `*bool` | Optional | Whether shared radio scanning is enabled for AP radio scanning<br><br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | Whether Wi-Fi settings in this AP template are enabled<br><br>**Default**: `true` |
| `LocateConnected` | `*bool` | Optional | Whether location tracking is enabled for connected clients<br><br>**Default**: `false` |
| `LocateUnconnected` | `*bool` | Optional | Whether location tracking is enabled for unconnected clients<br><br>**Default**: `false` |
| `MeshAllowDfs` | `*bool` | Optional | Whether mesh links may use DFS channels, which can add CAC delays during scanning<br><br>**Default**: `false` |
| `MeshEnableCrm` | `*bool` | Optional | Whether CRM is enabled for mesh networking in this AP template |
| `MeshEnabled` | `*bool` | Optional | Whether mesh networking is enabled by this AP template |
| `ProxyArp` | `*bool` | Optional | Whether proxy ARP is enabled for Wi-Fi clients<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apTemplateWifi := models.ApTemplateWifi{
        CiscoEnabled:                      models.ToPointer(false),
        Disable11k:                        models.ToPointer(false),
        DisableRadiosWhenPowerConstrained: models.ToPointer(false),
        EnableArpSpoof:                    models.ToPointer(false),
        EnableSharedRadioScanning:         models.ToPointer(false),
        Enabled:                           models.ToPointer(true),
        LocateConnected:                   models.ToPointer(false),
        LocateUnconnected:                 models.ToPointer(false),
        MeshAllowDfs:                      models.ToPointer(false),
        ProxyArp:                          models.ToPointer(false),
    }

}
```

