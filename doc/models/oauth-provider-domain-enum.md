
# Oauth Provider Domain Enum

If `oauth_type`==`okta`, specifies the region-specific OAuth provider domain. enum: `okta.com`, `oktapreview.com`, `okta-emea.com`, `okta-gov.com`, `okta.mil`, `mtls.okta.com`

## Enumeration

`OauthProviderDomainEnum`

## Fields

| Name |
|  --- |
| `okta.com` |
| `oktapreview.com` |
| `okta-emea.com` |
| `okta-gov.com` |
| `okta.mil` |
| `mtls.okta.com` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    oauthProviderDomain := models.OauthProviderDomainEnum_ENUMOKTACOM

}
```

