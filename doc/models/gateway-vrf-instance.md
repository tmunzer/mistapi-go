
# Gateway Vrf Instance

Gateway VRF instance and its member networks

## Structure

`GatewayVrfInstance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Networks` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayVrfInstance := models.GatewayVrfInstance{
        Networks:             []string{
            "CORP_NET",
            "MGMT_NET",
        },
    }

}
```

