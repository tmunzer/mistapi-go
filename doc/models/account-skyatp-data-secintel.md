
# Account Skyatp Data Secintel

juniper secintel_feeds are enabled depending on your license tier: infected_host, geo_ip, attacker_ip, command_and_control.
third party:

* ip-based: block_list, threatfox_ip, feodo_tracker, dshield, tor
* url-based: threatfox_url, urlhaus, open_phish
* domain-based: threatfox_domains

## Structure

`AccountSkyatpDataSecintel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ThirdPartyThreatFeeds` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountSkyatpDataSecintel := models.AccountSkyatpDataSecintel{
        ThirdPartyThreatFeeds: []string{
            "third_party_threat_feeds0",
        },
    }

}
```

