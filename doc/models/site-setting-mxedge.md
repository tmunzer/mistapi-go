
# Site Setting Mxedge

Service settings for the site Mist Edge cluster

## Structure

`SiteSettingMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MistDas` | [`*models.MxedgeDas`](../../doc/models/mxedge-das.md) | Optional | Cloud-assisted Dynamic Authorization Service settings for a Mist Edge cluster |
| `MistNac` | [`*models.MxclusterNac`](../../doc/models/mxcluster-nac.md) | Optional | Mist NAC RADIUS settings for a Mist Edge cluster. Used when the Mist Edge Cluster is used as a RADIUS Proxy between the local devices and the Mist NAC |
| `MistNacedge` | [`*models.MistNacedge`](../../doc/models/mist-nacedge.md) | Optional | Mist NAC Site Survivability settings for the site |
| `Radsec` | [`*models.MxclusterRadsec`](../../doc/models/mxcluster-radsec.md) | Optional | RadSec proxy configuration for a Mist Edge cluster. Used when the Mist Edge Cluster is used as a RADIUS Proxy between the local devices and external RADIUS Server. |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    siteSettingMxedge := models.SiteSettingMxedge{
        MistDas:              models.ToPointer(models.MxedgeDas{
            CoaServers:           []models.MxedgeDasCoaServer{
                models.MxedgeDasCoaServer{
                    DisableEventTimestampCheck:  models.ToPointer(false),
                    Enabled:                     models.ToPointer(false),
                    Host:                        models.ToPointer("host8"),
                    Port:                        models.ToPointer(28),
                    RequireMessageAuthenticator: models.ToPointer(false),
                },
                models.MxedgeDasCoaServer{
                    DisableEventTimestampCheck:  models.ToPointer(false),
                    Enabled:                     models.ToPointer(false),
                    Host:                        models.ToPointer("host8"),
                    Port:                        models.ToPointer(28),
                    RequireMessageAuthenticator: models.ToPointer(false),
                },
            },
            Enabled:              models.ToPointer(false),
        }),
        MistNac:              models.ToPointer(models.MxclusterNac{
            AcctServerPort:       models.ToPointer(70),
            AuthServerPort:       models.ToPointer(34),
            ClientIps:            map[string]models.MxclusterNacClientIp{
                "key0": models.MxclusterNacClientIp{
                    RequireMessageAuthenticator: models.ToPointer(false),
                    Secret:                      models.ToPointer("secret4"),
                    SiteId:                      models.ToPointer(uuid.MustParse("0000197c-0000-0000-0000-000000000000")),
                    Vendor:                      models.ToPointer(models.MxclusterNacClientVendorEnum_CISCOAIRONET),
                },
                "key1": models.MxclusterNacClientIp{
                    RequireMessageAuthenticator: models.ToPointer(false),
                    Secret:                      models.ToPointer("secret4"),
                    SiteId:                      models.ToPointer(uuid.MustParse("0000197c-0000-0000-0000-000000000000")),
                    Vendor:                      models.ToPointer(models.MxclusterNacClientVendorEnum_CISCOAIRONET),
                },
                "key2": models.MxclusterNacClientIp{
                    RequireMessageAuthenticator: models.ToPointer(false),
                    Secret:                      models.ToPointer("secret4"),
                    SiteId:                      models.ToPointer(uuid.MustParse("0000197c-0000-0000-0000-000000000000")),
                    Vendor:                      models.ToPointer(models.MxclusterNacClientVendorEnum_CISCOAIRONET),
                },
            },
            Enabled:              models.ToPointer(false),
            Secret:               models.ToPointer("secret6"),
        }),
        MistNacedge:          models.ToPointer(models.MistNacedge{
            AuthTtl:              models.ToPointer(110),
            CachingSiteIds:       []uuid.UUID{
                uuid.MustParse("000023ff-0000-0000-0000-000000000000"),
                uuid.MustParse("00002400-0000-0000-0000-000000000000"),
                uuid.MustParse("00002401-0000-0000-0000-000000000000"),
            },
            DefaultDot1xVlan:     models.ToPointer("default_dot1x_vlan4"),
            DefaultVlan:          models.ToPointer("default_vlan6"),
            Enabled:              models.ToPointer(false),
        }),
        Radsec:               models.ToPointer(models.MxclusterRadsec{
            AcctServers:          []models.MxclusterRadsecAcctServer{
                models.MxclusterRadsecAcctServer{
                    Host:                 models.ToPointer("host4"),
                    Port:                 models.ToPointer(254),
                    Secret:               models.ToPointer("secret0"),
                    Ssids:                []string{
                        "ssids5",
                        "ssids6",
                    },
                },
            },
            AuthServers:          []models.MxclusterRadsecAuthServer{
                models.MxclusterRadsecAuthServer{
                    Host:                 models.ToPointer("host0"),
                    InbandStatusCheck:    models.ToPointer(false),
                    InbandStatusInterval: models.ToPointer(160),
                    KeywrapEnabled:       models.ToPointer(false),
                    KeywrapFormat:        models.NewOptional(models.ToPointer(models.MxclusterRadAuthServerKeywrapFormatEnum_ASCII)),
                },
            },
            Enabled:              models.ToPointer(false),
            MatchSsid:            models.ToPointer(false),
            NasIpSource:          models.ToPointer(models.MxclusterRadsecNasIpSourceEnum_TUNNEL),
        }),
    }

}
```

