
# License Usage Org

## Structure

`LicenseUsageOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForSite` | `*bool` | Optional | - |
| `FullyLoaded` | `map[string]int` | Optional | Property key is the service name (e.g. "SUB-MAN") |
| `NumDevices` | `int` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Usages` | `map[string]int` | Required | subscriptions and their quantities. Property key is the service name (e.g. "SUB-MAN") |

## Example (as JSON)

```json
{
  "num_devices": 42,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "usages": {
    "key0": 249
  },
  "for_site": false,
  "fully_loaded": {
    "key0": 156
  }
}
```

