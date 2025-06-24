
# License Usage Org

*This model accepts additional fields of type interface{}.*

## Structure

`LicenseUsageOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForSite` | `*bool` | Optional | - |
| `FullyLoaded` | `map[string]int` | Optional | Maximum number of licenses that may be required if the service is enabled on all the Organization Devices. Property key is the service name (e.g. "SUB-MAN"). |
| `NumDevices` | `int` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Summary` | `map[string]int` | Optional | Number of licenses currently consumed. Property key is license type (e.g. SUB-MAN). |
| `Usages` | `map[string]int` | Required | Number of available licenes. Property key is the service name (e.g. "SUB-MAN"). name (e.g. "SUB-MAN") |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "num_devices": 254,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "usages": {
    "key0": 205
  },
  "for_site": false,
  "fully_loaded": {
    "key0": 56
  },
  "summary": {
    "key0": 120,
    "key1": 121
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

