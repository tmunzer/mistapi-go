
# Nac Rule Matching

*This model accepts additional fields of type interface{}.*

## Structure

`NacRuleMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthType` | [`*models.NacRuleMatchingAuthTypeEnum`](../../doc/models/nac-rule-matching-auth-type-enum.md) | Optional | enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `peap-tls`, `psk` |
| `Nactags` | `[]string` | Optional | - |
| `PortTypes` | [`[]models.NacRuleMatchingPortTypeEnum`](../../doc/models/nac-rule-matching-port-type-enum.md) | Optional | - |
| `SiteIds` | `[]uuid.UUID` | Optional | List of site ids to match |
| `SitegroupIds` | `[]uuid.UUID` | Optional | List of sitegroup ids to match |
| `Vendor` | `[]string` | Optional | List of vendors to match |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
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
  "auth_type": "idp",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

