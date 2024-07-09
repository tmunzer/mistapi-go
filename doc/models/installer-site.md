
# Installer Site

## Structure

`InstallerSite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `string` | Required | - |
| `CountryCode` | `string` | Required | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Latlng` | [`models.LatLng`](../../doc/models/lat-lng.md) | Required | - |
| `Name` | `string` | Required | - |
| `RftemplateName` | `*string` | Optional | - |
| `SitegroupNames` | `[]string` | Optional | - |
| `Timezone` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "address": "1601 S. Deanza Blvd., Cupertino, CA, 95014",
  "country_code": "US",
  "id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
  "latlng": {
    "lat": 37.295833,
    "lng": -122.032946
  },
  "name": "Mist Office",
  "rftemplate_name": "rftemplate1",
  "sitegroup_names": [
    "sg1",
    "sg2"
  ],
  "timezone": "America/Los_Angeles"
}
```

