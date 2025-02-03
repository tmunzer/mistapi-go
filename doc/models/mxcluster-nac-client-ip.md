
# Mxcluster Nac Client Ip

*This model accepts additional fields of type interface{}.*

## Structure

`MxclusterNacClientIp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RequireMessageAuthenticator` | `*bool` | Optional | Whether to require Message-Authenticator in requests<br>**Default**: `false` |
| `Secret` | `*string` | Optional | If different from above |
| `SiteId` | `*uuid.UUID` | Optional | Present only for 3rd party clients |
| `Vendor` | [`*models.MxclusterNacClientVendorEnum`](../../doc/models/mxcluster-nac-client-vendor-enum.md) | Optional | convention to be followed is : "<vendor>-<variant>", <variant> could be an os/platform/model/company. For ex: for cisco vendor, there could variants wrt os (such as ios, nxos etc), platforms (asa etc), or acquired companies (such as meraki, airnonet) etc. enum: `aruba`, `cisco-aironet`, `cisco-ios`, `cisco-meraki`, `generic`, `juniper`, `paloalto` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "require_message_authenticator": false,
  "site_id": "00000000-0000-0000-1234-000000000000",
  "vendor": "cisco-ios",
  "secret": "secret8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

