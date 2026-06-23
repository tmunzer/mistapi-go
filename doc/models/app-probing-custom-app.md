
# App Probing Custom App

User-defined application probe definition

## Structure

`AppProbingCustomApp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `*string` | Optional | Required if `protocol`==`icmp`. IP address probed by the ICMP custom app |
| `AppType` | `*string` | Optional | Category label used for this custom application probe |
| `Hostnames` | `[]string` | Optional | If `protocol`==`http`. HTTP hostnames or URLs probed by the custom app |
| `Key` | `*string` | Optional | Stable key used to identify this custom application probe |
| `Name` | `*string` | Optional | Display name for this custom application probe |
| `Network` | `*string` | Optional | Gateway network used as the source context for this probe |
| `PacketSize` | `*int` | Optional | If `protocol`==`icmp`. ICMP packet size used by this custom app probe<br><br>**Constraints**: `>= 0`, `<= 65400` |
| `Protocol` | [`*models.AppProbingCustomAppProtocolEnum`](../../doc/models/app-probing-custom-app-protocol-enum.md) | Optional | Probe protocol for a custom application. enum: `http`, `icmp`<br><br>**Default**: `"http"` |
| `Url` | `*string` | Optional | If `protocol`==`http`. HTTP URL or hostname probed by this custom app |
| `Vrf` | `*string` | Optional | Gateway VRF used as the source context for this probe |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    appProbingCustomApp := models.AppProbingCustomApp{
        Address:              models.ToPointer("192.168.1.1"),
        AppType:              models.ToPointer("app_type4"),
        Hostnames:            []string{
            "https://www.abc.com",
        },
        Key:                  models.ToPointer("key0"),
        Name:                 models.ToPointer("pos_app"),
        Network:              models.ToPointer("lan"),
        Protocol:             models.ToPointer(models.AppProbingCustomAppProtocolEnum_HTTP),
        Url:                  models.ToPointer("www.abc.com"),
        Vrf:                  models.ToPointer("lan"),
    }

}
```

