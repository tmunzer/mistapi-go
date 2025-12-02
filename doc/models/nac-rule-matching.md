
# Nac Rule Matching

## Structure

`NacRuleMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthType` | [`*models.NacAuthTypeEnum`](../../doc/models/nac-auth-type-enum.md) | Optional | enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `eap-peap` |
| `Family` | `[]string` | Optional | List of client device families to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed family values |
| `Mfg` | `[]string` | Optional | List of client device models to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed model values |
| `Model` | `[]string` | Optional | List of client device manufacturers to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed mfg values |
| `Nactags` | `[]string` | Optional | - |
| `OsType` | `[]string` | Optional | List of client device os types to match. Refer to [List Fingerprint Types]]($e/Constants%20Definitions/listFingerprintTypes) for allowed os_type values |
| `PortTypes` | [`[]models.NacRuleMatchingPortTypeEnum`](../../doc/models/nac-rule-matching-port-type-enum.md) | Optional | - |
| `SiteIds` | `[]uuid.UUID` | Optional | List of site ids to match |
| `SitegroupIds` | `[]uuid.UUID` | Optional | List of sitegroup ids to match |
| `Vendor` | `[]string` | Optional | List of vendors to match |

## Example (as JSON)

```json
{
  "auth_type": "eap-tls",
  "nactags": [
    "041d5d36-716c-4cfb-4988-3857c6aa14a2",
    "a809a97f-d599-f812-eb8c-c3f84aabf6ba"
  ],
  "port_types": [
    "wired"
  ],
  "site_ids": [
    "bb19fc3e-4124-4b57-80d9-c3f6edce47c4",
    "bb19fc3e-6564-4b57-80d9-c3f6edce47c1"
  ],
  "sitegroup_ids": [
    "bb19fc3e-4124-4b57-80d9-c3f6edce47c4",
    "bb19fc3e-6564-4b57-80d9-c3f6edce47c1"
  ],
  "family": [
    "family7",
    "family8",
    "family9"
  ],
  "mfg": [
    "mfg4",
    "mfg3"
  ],
  "model": [
    "model6",
    "model7",
    "model8"
  ]
}
```

