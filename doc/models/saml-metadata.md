
# Saml Metadata

Read-only SAML and SCIM metadata generated for an SSO configuration

## Structure

`SamlMetadata`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcsUrl` | `*string` | Optional, Read-only | If `idp_type`==`saml`, Assertion Consumer Service URL that receives SAML responses for this Mist SSO configuration |
| `EntityId` | `*string` | Optional, Read-only | If `idp_type`==`saml`, service provider entity ID for this Mist SSO configuration |
| `LogoutUrl` | `*string` | Optional, Read-only | If `idp_type`==`saml`, Single Logout URL used by the identity provider to end the Mist SSO session |
| `Metadata` | `*string` | Optional, Read-only | If `idp_type`==`saml`, service provider metadata XML for configuring the identity provider |
| `ScimBaseUrl` | `*string` | Optional | If `idp_type`==`oauth` and `scim_enabled`==`true`, SCIM base URL that the identity provider uses to send provisioning requests to Mist |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    samlMetadata := models.SamlMetadata{
        AcsUrl:               models.ToPointer("https://api.mist.com/api/v1/saml/llDfa13f/login"),
        EntityId:             models.ToPointer("https://api.mist.com/api/v1/saml/llDfa13f/login"),
        LogoutUrl:            models.ToPointer("https://api.mist.com/api/v1/saml/llDfa13f/logout"),
        Metadata:             models.ToPointer("<?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://api.mist.com/api/v1/saml/llDfa13f/login\" validUntil=\"2027-10-12T21:59:01Z\" xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><md:SPSSODescriptor AuthnRequestsSigned=\"false\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://api.mist.com/api/v1/saml/llDfa13f/login\" index=\"0\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>"),
        ScimBaseUrl:          models.ToPointer("https://scim.nac-staging.mistsys.com/S_41b2525a-e8b8-4809-8168-f1d8dcbe9735/azure/4d72b1dc-7503-4717-81ea-80d0125b886e"),
    }

}
```

