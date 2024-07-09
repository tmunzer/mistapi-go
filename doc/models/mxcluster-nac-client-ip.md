
# Mxcluster Nac Client Ip

## Structure

`MxclusterNacClientIp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Secret` | `*string` | Optional | if different from above |
| `SiteId` | `*uuid.UUID` | Optional | present only for 3rd party clients |
| `Vendor` | [`*models.MxclusterNacClientVendorEnum`](../../doc/models/mxcluster-nac-client-vendor-enum.md) | Optional | convention to be followed is : "<vendor>-<variant>"<br><variant> could be an os/platform/model/company<br>for ex: for cisco vendor, there could variants wrt os (such as ios, nxos etc), platforms (asa etc), or acquired companies (such as meraki, airnonet) etc. |

## Example (as JSON)

```json
{
  "site_id": "00000000-0000-0000-1234-000000000000",
  "vendor": "cisco-ios",
  "secret": "secret8"
}
```

