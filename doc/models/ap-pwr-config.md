
# Ap Pwr Config

Power negotiation and peripheral power settings for an AP or AP profile

## Structure

`ApPwrConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Base` | `*int` | Optional | Additional power to request during negotiating with PSE over PoE, in mW<br><br>**Default**: `0` |
| `PreferUsbOverWifi` | `*bool` | Optional | Whether to enable power out to peripheral, meanwhile will reduce power to Wi-Fi (only for AP45 at power mode)<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apPwrConfig := models.ApPwrConfig{
        Base:                 models.ToPointer(2000),
        PreferUsbOverWifi:    models.ToPointer(false),
    }

}
```

