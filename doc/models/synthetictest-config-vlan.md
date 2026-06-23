
# Synthetictest Config Vlan

Deprecated VLAN-based synthetic test settings

## Structure

`SynthetictestConfigVlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomTestUrls` | `[]string` | Optional | Deprecated custom URLs tested by VLAN-based synthetic probes |
| `Disabled` | `*bool` | Optional | For some vlans where we don't want this to run<br><br>**Default**: `false` |
| `Probes` | `[]string` | Optional | app name comes from `custom_probes` above or /const/synthetic_test_probes |
| `VlanIds` | [`[]models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | VLAN ID, either numeric or expressed as a template variable string |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    synthetictestConfigVlan := models.SynthetictestConfigVlan{
        CustomTestUrls:       []string{
            "https://www.abc.com/",
            "https://10.3.5.1:8080/about",
        },
        Disabled:             models.ToPointer(false),
        Probes:               []string{
            "probes9",
        },
        VlanIds:              []models.VlanIdWithVariable{
            models.VlanIdWithVariableContainer.FromNumber(10),
            models.VlanIdWithVariableContainer.FromNumber(20),
            models.VlanIdWithVariableContainer.FromString("{{vlan}}"),
        },
    }

}
```

