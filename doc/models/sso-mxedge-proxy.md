
# Sso Mxedge Proxy

Mist Edge proxy settings for NAC SSO. If `idp_type`==`mxedge_proxy`, this requires `mist_nac` to be enabled on the mxcluster

## Structure

`SsoMxedgeProxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServers` | [`[]models.SsoMxedgeProxyAcctServer`](../../doc/models/sso-mxedge-proxy-acct-server.md) | Optional | RADIUS accounting servers used by the Mist Edge SSO proxy |
| `AuthServers` | [`[]models.SsoMxedgeProxyAuthServer`](../../doc/models/sso-mxedge-proxy-auth-server.md) | Optional | RADIUS authentication servers used by the Mist Edge SSO proxy |
| `MxclusterId` | `*uuid.UUID` | Optional | Mist Edge cluster identifier that provides the SSO proxy |
| `OperatorName` | `*string` | Optional | Operator name as RADIUS attribute while proxying |
| `ProxyHosts` | `[]string` | Optional | Public hostnames or IP addresses for the Mist Edge SSO proxy |
| `Ssids` | `[]string` | Optional | Eduroam SSID names handled by the Mist Edge SSO proxy |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    ssoMxedgeProxy := models.SsoMxedgeProxy{
        AcctServers:          []models.SsoMxedgeProxyAcctServer{
            models.SsoMxedgeProxyAcctServer{
                Host:                 models.ToPointer("host4"),
                Port:                 models.ToPointer(254),
                Secret:               models.ToPointer("secret0"),
            },
            models.SsoMxedgeProxyAcctServer{
                Host:                 models.ToPointer("host4"),
                Port:                 models.ToPointer(254),
                Secret:               models.ToPointer("secret0"),
            },
        },
        AuthServers:          []models.SsoMxedgeProxyAuthServer{
            models.SsoMxedgeProxyAuthServer{
                Host:                        models.ToPointer("host0"),
                Port:                        models.ToPointer(114),
                RequireMessageAuthenticator: models.ToPointer(false),
                Retry:                       models.ToPointer(126),
                Secret:                      models.ToPointer("secret4"),
            },
            models.SsoMxedgeProxyAuthServer{
                Host:                        models.ToPointer("host0"),
                Port:                        models.ToPointer(114),
                RequireMessageAuthenticator: models.ToPointer(false),
                Retry:                       models.ToPointer(126),
                Secret:                      models.ToPointer("secret4"),
            },
        },
        MxclusterId:          models.ToPointer(uuid.MustParse("572586b7-f97b-a22b-526c-8b97a3f609c4")),
        OperatorName:         models.ToPointer("operator_name8"),
        ProxyHosts:           []string{
            "mxedge1.corp.com",
            "63.1.3.5",
        },
        Ssids:                []string{
            "eduroam_test, eduroam_main",
        },
    }

}
```

