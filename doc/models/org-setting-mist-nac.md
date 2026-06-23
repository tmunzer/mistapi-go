
# Org Setting Mist Nac

Organization-level Mist NAC configuration

## Structure

`OrgSettingMistNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowTeapMachineAuthOnly` | `*bool` | Optional | allow clients to connect even when the user cert failed. TEAP authenticates both Machine Cert and User Cert. When enabled, clients who only succeed Machine Cert authentication will be accepted.<br><br>**Default**: `false` |
| `Cacerts` | `[]string` | Optional | List of PEM-encoded ca certs |
| `DefaultIdpId` | `*string` | Optional | use this IDP when no explicit realm present in the incoming username/CN OR when no IDP is explicitly mapped to the incoming realm. |
| `DisableRsaeAlgorithms` | `*bool` | Optional | to disable RSAE_PSS_SHA256, RSAE_PSS_SHA384, RSAE_PSS_SHA512 from server side. see https://www.openssl.org/docs/man3.0/man1/openssl-ciphers.html<br><br>**Default**: `false` |
| `EapSslSecurityLevel` | `*int` | Optional | eap ssl security level, see https://www.openssl.org/docs/man1.1.1/man3/SSL_CTX_set_security_level.html#DEFAULT-CALLBACK-BEHAVIOUR<br><br>**Default**: `2`<br><br>**Constraints**: `>= 1`, `<= 4` |
| `EuOnly` | `*bool` | Optional | By default, NAC POD failover considers all NAC pods available around the globe, i.e. EU, US, or APAC based, failover happens based on geo IP of the originating site. For strict GDPR compliance NAC POD failover would only happen between the PODs located within the EU environment, and no authentication would take place outside of EU. This is an org setting that is applicable to WLANs, switch templates, Mist Edge clusters that have mist_nac enabled<br><br>**Default**: `false` |
| `Fingerprinting` | [`*models.OrgSettingMistNacFingerprinting`](../../doc/models/org-setting-mist-nac-fingerprinting.md) | Optional | Client fingerprinting settings used for Mist NAC policy enforcement |
| `IdpMachineCertLookupField` | [`*models.IdpMachineCertLookupFieldEnum`](../../doc/models/idp-machine-cert-lookup-field-enum.md) | Optional | allow customer to choose the EAP-TLS client certificate's field to use for IDP Machine Groups lookup. enum: `automatic`, `cn`, `dns`<br><br>**Default**: `"automatic"` |
| `IdpUserCertLookupField` | [`*models.IdpUserCertLookupFieldEnum`](../../doc/models/idp-user-cert-lookup-field-enum.md) | Optional | allow customer to choose the EAP-TLS client certificate's field. To use for IDP User Groups lookup. enum: `automatic`, `cn`, `email`, `upn`<br><br>**Default**: `"automatic"` |
| `Idps` | [`[]models.OrgSettingMistNacIdp`](../../doc/models/org-setting-mist-nac-idp.md) | Optional | Identity provider realm mappings used by Mist NAC |
| `Mdm` | [`*models.OrgSettingMistNacMdm`](../../doc/models/org-setting-mist-nac-mdm.md) | Optional | MDM (Mobile Device Management) CoA configuration |
| `ServerCert` | [`*models.OrgSettingMistNacServerCert`](../../doc/models/org-setting-mist-nac-server-cert.md) | Optional | RADIUS server certificate presented by Mist NAC during EAP-TLS |
| `UseIpVersion` | [`*models.OrgSettingMistNacIpVersionEnum`](../../doc/models/org-setting-mist-nac-ip-version-enum.md) | Optional | by default, NAS devices(switches/aps) and proxies(mxedge) are configured to reach mist-nac via IPv4. enum: `v4`, `v6`<br><br>**Default**: `"v4"` |
| `UseSslPort` | `*bool` | Optional | By default, NAS devices (switches/aps) and proxies(mxedge) are configured to use port TCP2083(RadSec) to reach mist-nac. Set `use_ssl_port`==`true` to override that port with TCP43 (ssl), This is an org level setting that is applicable to wlans, switch_templates, and mxedge_clusters that have mist-nac enabled<br><br>**Default**: `false` |
| `UsermacExpiry` | `*int` | Optional | Allow customer to configure an expiry time for usermacs by attaching a Quarantine label to those which have been inactive for the configured period of time (in days). 0 means no expiry<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 1095` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingMistNac := models.OrgSettingMistNac{
        AllowTeapMachineAuthOnly:  models.ToPointer(false),
        Cacerts:                   []string{
            "cacerts2",
            "cacerts3",
            "cacerts4",
        },
        DefaultIdpId:              models.ToPointer("default_idp_id2"),
        DisableRsaeAlgorithms:     models.ToPointer(false),
        EapSslSecurityLevel:       models.ToPointer(2),
        EuOnly:                    models.ToPointer(false),
        IdpMachineCertLookupField: models.ToPointer(models.IdpMachineCertLookupFieldEnum_AUTOMATIC),
        IdpUserCertLookupField:    models.ToPointer(models.IdpUserCertLookupFieldEnum_AUTOMATIC),
        UseIpVersion:              models.ToPointer(models.OrgSettingMistNacIpVersionEnum_V4),
        UseSslPort:                models.ToPointer(false),
        UsermacExpiry:             models.ToPointer(30),
    }

}
```

