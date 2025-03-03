
# Inventory Search Result

*This model accepts additional fields of type interface{}.*

## Structure

`InventorySearchResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | - |
| `Master` | `*bool` | Optional | - |
| `Members` | [`[]models.InventorySearchResultMember`](../../doc/models/inventory-search-result-member.md) | Optional | - |
| `Model` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Serial` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Sku` | `*string` | Optional | - |
| `Status` | `*string` | Optional | - |
| `Type` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Optional | enum: `ap`, `gateway`, `switch`<br>**Default**: `"ap"` |
| `VcMac` | `*string` | Optional | - |
| `Version` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "f01c2df166e0",
  "master": true,
  "model": "EX4300-48P",
  "name": "mist-wa-ex4300-VC",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "serial": "PD3714460200",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "sku": "EX4300-48P",
  "status": "disconnected",
  "type": "ap",
  "vc_mac": "f01c2df166e0",
  "version": "21.4R3.5",
  "members": [
    {
      "mac": "mac2",
      "model": "model6",
      "serial": "serial8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "mac": "mac2",
      "model": "model6",
      "serial": "serial8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "mac": "mac2",
      "model": "model6",
      "serial": "serial8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

