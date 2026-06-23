
# Synthetictest Config Lan Network

configure minis probes to be tested on lan networks of gateways

## Structure

`SynthetictestConfigLanNetwork`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Networks` | `[]string` | Optional | List of networks to be used for synthetic tests |
| `Probes` | `[]string` | Optional | app name comes from `custom_probes` above or /const/synthetic_test_probes |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    synthetictestConfigLanNetwork := models.SynthetictestConfigLanNetwork{
        Networks:             []string{
            "pos-stations",
            "pos-machines",
        },
        Probes:               []string{
            "probes9",
        },
    }

}
```

