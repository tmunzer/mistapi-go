
# Js Inventory Item

## Structure

`JsInventoryItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Claimed` | `*bool` | Optional | Indicates if the device is claimed by any org |
| `DeviceName` | `*string` | Optional | Name of the device |
| `EolTime` | `*int` | Optional | End of life time |
| `EosTime` | `*int` | Optional | End of support time |
| `HasSupport` | `*bool` | Optional | Indicates if the device is covered under active support contract |
| `Master` | `*bool` | Optional | Indicates whether it is Master |
| `Model` | `*string` | Optional | Model of the install base inventory |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Serial` | `*string` | Optional | Serial Number of the inventory |
| `Sku` | `*string` | Optional | Serviceable device stock |
| `Status` | `*string` | Optional | Status of the connected device |
| `SuggestedVersion` | `*string` | Optional | Suggested SW version |
| `Type` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Optional | enum: `ap`, `gateway`, `switch` |
| `Version` | `*string` | Optional | SW version running |
| `VersionEosTime` | `*int` | Optional | End of Service of version |
| `VersionTime` | `*int` | Optional | FRS date of the version |
| `Warranty` | `*string` | Optional | warranty description |
| `WarrantyTime` | `*int` | Optional | Time when warranty needs to be renewed |
| `WarrantyType` | [`*models.JsiWarrantyTypeEnum`](../../doc/models/jsi-warranty-type-enum.md) | Optional | Warranty type for Juniper Support Insight (JSI) devices. The warranty type<br>is used to determine the support level and duration of the warranty for the<br>device. enum:<br><br>* WTY00001: Standard Hardware Warranty<br>* WTY00002: Enhanced Hardware Warranty<br>* WTY00003: Dead On Arrival Warranty<br>* WTY00004: Limited Lifetime Warranty<br>* WTY00005: Software Warranty<br>* WTY00006: Limited Lifetime Warranty for WLA<br>* WTY00007: Warranty-JCPO EOL (DOA Not Included)<br>* WTY00008: MIST Enhanced Hardware Warranty<br>* WTY00009: MIST Standard Warranty<br>* WTY00099: Determine Lifetime warranty |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "claimed": false,
  "device_name": "device_name8",
  "eol_time": 58,
  "eos_time": 104,
  "has_support": false
}
```

