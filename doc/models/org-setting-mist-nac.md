
# Org Setting Mist Nac

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingMistNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cacerts` | `[]string` | Optional | List of PEM-encoded ca certs |
| `DefaultIdpId` | `*string` | Optional | use this IDP when no explicit realm present in the incoming username/CN OR when no IDP is explicitly mapped to the incoming realm. |
| `DisableRsaeAlgorithms` | `*bool` | Optional | to disable RSAE_PSS_SHA256, RSAE_PSS_SHA384, RSAE_PSS_SHA512 from server side. see https://www.openssl.org/docs/man3.0/man1/openssl-ciphers.html<br><br>**Default**: `false` |
| `EapSslSecurityLevel` | `*int` | Optional | eap ssl security level, see https://www.openssl.org/docs/man1.1.1/man3/SSL_CTX_set_security_level.html#DEFAULT-CALLBACK-BEHAVIOUR<br><br>**Default**: `2`<br><br>**Constraints**: `>= 1`, `<= 4` |
| `EuOnly` | `*bool` | Optional | By default, NAC POD failover considers all NAC pods available around the globe, i.e. EU, US, or APAC based, failover happens based on geo IP of the originating site. For strict GDPR compliance NAC POD failover would only happen between the PODs located within the EU environment, and no authentication would take place outside of EU. This is an org setting that is applicable to WLANs, switch templates, mxedge clusters that have mist_nac enabled<br><br>**Default**: `false` |
| `IdpMachineCertLookupField` | [`*models.IdpMachineCertLookupFieldEnum`](../../doc/models/idp-machine-cert-lookup-field-enum.md) | Optional | allow customer to choose the EAP-TLS client certificate's field to use for IDP Machine Groups lookup. enum: `automatic`, `cn`, `dns`<br><br>**Default**: `"automatic"` |
| `IdpUserCertLookupField` | [`*models.IdpUserCertLookupFieldEnum`](../../doc/models/idp-user-cert-lookup-field-enum.md) | Optional | allow customer to choose the EAP-TLS client certificate's field. To use for IDP User Groups lookup. enum: `automatic`, `cn`, `email`, `upn`<br><br>**Default**: `"automatic"` |
| `Idps` | [`[]models.OrgSettingMistNacIdp`](../../doc/models/org-setting-mist-nac-idp.md) | Optional | - |
| `ServerCert` | [`*models.OrgSettingMistNacServerCert`](../../doc/models/org-setting-mist-nac-server-cert.md) | Optional | radius server cert to be presented in EAP TLS |
| `UseIpVersion` | [`*models.OrgSettingMistNacIpVersionEnum`](../../doc/models/org-setting-mist-nac-ip-version-enum.md) | Optional | by default, NAS devices(switches/aps) and proxies(mxedge) are configured to reach mist-nac via IPv4. enum: `v4`, `v6`<br><br>**Default**: `"v4"` |
| `UseSslPort` | `*bool` | Optional | By default, NAS devices (switches/aps) and proxies(mxedge) are configured to use port TCP2083(RadSec) to reach mist-nac. Set `use_ssl_port`==`true` to override that port with TCP43 (ssl), This is an org level setting that is applicable to wlans, switch_templates, and mxedge_clusters that have mist-nac enabled<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disable_rsae_algorithms": false,
  "eap_ssl_security_level": 2,
  "eu_only": false,
  "idp_machine_cert_lookup_field": "automatic",
  "idp_user_cert_lookup_field": "automatic",
  "use_ip_version": "v4",
  "use_ssl_port": false,
  "cacerts": [
    "cacerts6",
    "cacerts5",
    "cacerts4"
  ],
  "default_idp_id": "default_idp_id0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

