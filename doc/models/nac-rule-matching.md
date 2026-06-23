
# Nac Rule Matching

Criteria used to include or exclude a NAC authentication request from a rule

## Structure

`NacRuleMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthType` | [`*models.NacAuthTypeEnum`](../../doc/models/nac-auth-type-enum.md) | Optional | enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `eap-peap` |
| `Family` | `[]string` | Optional | List of client device families to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed family values |
| `Mfg` | `[]string` | Optional | List of client device models to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed model values |
| `Model` | `[]string` | Optional | List of client device manufacturers to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed mfg values |
| `Nactags` | `[]string` | Optional | NAC tag identifiers used as rule-matching criteria |
| `OsType` | `[]string` | Optional | List of client device os types to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed os_type values |
| `PortTypes` | [`[]models.NacRuleMatchingPortTypeEnum`](../../doc/models/nac-rule-matching-port-type-enum.md) | Optional | Wired or wireless access types used as NAC rule-matching criteria |
| `SiteIds` | `[]uuid.UUID` | Optional | List of site ids to match |
| `SitegroupIds` | `[]uuid.UUID` | Optional | List of sitegroup ids to match |
| `Vendor` | `[]string` | Optional | List of vendors to match |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    nacRuleMatching := models.NacRuleMatching{
        AuthType:             models.ToPointer(models.NacAuthTypeEnum_EAPTLS),
        Family:               []string{
            "family9",
            "family0",
        },
        Mfg:                  []string{
            "mfg2",
        },
        Model:                []string{
            "model8",
            "model9",
        },
        Nactags:              []string{
            "041d5d36-716c-4cfb-4988-3857c6aa14a2",
            "a809a97f-d599-f812-eb8c-c3f84aabf6ba",
        },
        PortTypes:            []models.NacRuleMatchingPortTypeEnum{
            models.NacRuleMatchingPortTypeEnum_WIRED,
        },
        SiteIds:              []uuid.UUID{
            uuid.MustParse("bb19fc3e-4124-4b57-80d9-c3f6edce47c4"),
            uuid.MustParse("bb19fc3e-6564-4b57-80d9-c3f6edce47c1"),
        },
        SitegroupIds:         []uuid.UUID{
            uuid.MustParse("bb19fc3e-4124-4b57-80d9-c3f6edce47c4"),
            uuid.MustParse("bb19fc3e-6564-4b57-80d9-c3f6edce47c1"),
        },
    }

}
```

