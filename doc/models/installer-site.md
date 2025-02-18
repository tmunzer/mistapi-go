
# Installer Site

*This model accepts additional fields of type interface{}.*

## Structure

`InstallerSite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `string` | Required | - |
| `CountryCode` | `string` | Required | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Latlng` | [`models.LatLng`](../../doc/models/lat-lng.md) | Required | - |
| `Name` | `string` | Required | - |
| `RftemplateName` | `*string` | Optional | - |
| `SitegroupNames` | `[]string` | Optional | - |
| `Timezone` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "address": "1601 S. Deanza Blvd., Cupertino, CA, 95014",
  "country_code": "US",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "latlng": {
    "lat": 37.295833,
    "lng": -122.032946,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "name": "Mist Office",
  "rftemplate_name": "rftemplate1",
  "sitegroup_names": [
    "sg1",
    "sg2"
  ],
  "timezone": "America/Los_Angeles",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

