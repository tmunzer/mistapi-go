
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

*This model accepts additional fields of type interface{}.*

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
| `PortalAllowedHostnames` | `[]string` | Optional | List of hostnames without http(s):// (matched by substring) |
| `PortalAllowedSubnets` | `[]string` | Optional | List of CIDRs |
| `PortalDeniedHostnames` | `[]string` | Optional | List of hostnames without http(s):// (matched by substring), this takes precedence over portal_allowed_hostnames |
| `Privacy` | `*bool` | Optional | If `auth`==`none` or `auth`==`multi`, whether to show the privacy policy |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "expire": 1440,
  "external_portal_url": "https://yourorg.com/external-guest-portal",
  "forward": true,
  "forward_url": "https://yourorg.com/guest-portal-redirect",
  "portal_allowed_hostnames": [
    "snapchat.com",
    "ibm.com"
  ],
  "portal_allowed_subnets": [
    "63.5.3.0/24"
  ],
  "portal_denied_hostnames": [
    "msg.snapchat.com"
  ],
  "privacy": true,
  "auth": "external",
  "force_reconnect": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

