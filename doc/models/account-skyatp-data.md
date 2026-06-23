
# Account Skyatp Data

Sky ATP SecIntel feed data and generated list URLs

*This model accepts additional fields of type interface{}.*

## Structure

`AccountSkyatpData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Secintel` | [`*models.AccountSkyatpDataSecintel`](../../doc/models/account-skyatp-data-secintel.md) | Optional | juniper secintel_feeds are enabled depending on your license tier: infected_host, geo_ip, attacker_ip, command_and_control.<br>third party:<br><br>* ip-based: block_list, threatfox_ip, feodo_tracker, dshield, tor<br>* url-based: threatfox_url, urlhaus, open_phish<br>* domain-based: threatfox_domains |
| `SecintelAllowlistUrl` | `*string` | Optional, Read-only | URL for the Sky ATP SecIntel allowlist |
| `SecintelBlocklistUrl` | `*string` | Optional, Read-only | URL for the Sky ATP SecIntel blocklist |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountSkyatpData := models.AccountSkyatpData{
        Secintel:             models.ToPointer(models.AccountSkyatpDataSecintel{
            ThirdPartyThreatFeeds: []string{
                "third_party_threat_feeds0",
                "third_party_threat_feeds1",
            },
        }),
        SecintelAllowlistUrl: models.ToPointer("https://papi.s3.amazonaws.com/secintel_allowlist/xxx..."),
        SecintelBlocklistUrl: models.ToPointer("https://papi.s3.amazonaws.com/secintel_blocklist/xxx..."),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

