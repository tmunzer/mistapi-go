
# Nac Portal Guest Portal

Guest portal configuration when `type`==`guest_portal`. If

* `auth`==`none`, the user is presented with a terms of service and can click and continue.
* `auth`==`external`, the user is redirected to an external URL for authentication.
* `auth`==`multi`, the user is presented with a choice of authentication methods:
  - social logins: facebook / google / amazon / microsoft / azure
  - sponsor
  - sms: supported provider: twillio
  - email
  - sso
  - userpass: pre created guest list

## Structure

`NacPortalGuestPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Auth` | [`*models.NacPortalGuestPortalAuthEnum`](../../doc/models/nac-portal-guest-portal-auth-enum.md) | Optional | Guest portal authentication type. enum: `external`, `multi`, `none` |
| `Expire` | `*int` | Optional | If `auth`==`none` or `auth`==`multi`, whether to expire the guest after a certain time |
| `ExternalPortalUrl` | `*string` | Optional | If `auth`==`external`, the URL to redirect the user to for authentication |
| `ForceReconnect` | `*bool` | Optional | Disconnect client (workaround for reauth issues) |
| `Forward` | `*bool` | Optional | If `auth`==`none` or `auth`==`multi`, whether to forward the user to the guest portal after authentication |
| `ForwardUrl` | `*string` | Optional | If `auth`==`none` or `auth`==`multi`, URL to forward the user to after authentication |
| `MaxNumDevices` | `*int` | Optional | Maximum number of clients allowed per guest. 0 (default, unlimited), 1-100 range<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 100` |
| `Privacy` | `*bool` | Optional | If `auth`==`none` or `auth`==`multi`, whether to show the privacy policy |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    nacPortalGuestPortal := models.NacPortalGuestPortal{
        Auth:                 models.ToPointer(models.NacPortalGuestPortalAuthEnum_MULTI),
        Expire:               models.ToPointer(1440),
        ExternalPortalUrl:    models.ToPointer("https://yourorg.com/external-guest-portal"),
        ForceReconnect:       models.ToPointer(false),
        Forward:              models.ToPointer(true),
        ForwardUrl:           models.ToPointer("https://yourorg.com/guest-portal-redirect"),
        MaxNumDevices:        models.ToPointer(10),
        Privacy:              models.ToPointer(true),
    }

}
```

