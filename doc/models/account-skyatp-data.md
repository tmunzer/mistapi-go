
# Account Skyatp Data

*This model accepts additional fields of type interface{}.*

## Structure

`AccountSkyatpData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Secintel` | [`*models.AccountSkyatpDataSecintel`](../../doc/models/account-skyatp-data-secintel.md) | Optional | juniper secintel_feeds are enabled depending on your license tier: infected_host, geo_ip, attacker_ip, command_and_control.<br>third party:<br><br>* ip-based: block_list, threatfox_ip, feodo_tracker, dshield, tor<br>* url-based: threatfox_url, urlhaus, open_phish<br>* domain-based: threatfox_domains |
| `SecintelAllowlistUrl` | `*string` | Optional | - |
| `SecintelBlocklistUrl` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "secintel_allowlist_url": "https://papi.s3.amazonaws.com/secintel_allowlist/xxx...",
  "secintel_blocklist_url": "https://papi.s3.amazonaws.com/secintel_blocklist/xxx...",
  "secintel": {
    "third_party_threat_feeds": [
      "third_party_threat_feeds0",
      "third_party_threat_feeds1"
    ]
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

