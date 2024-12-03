
# Other Device Update

*This model accepts additional fields of type interface{}.*

## Structure

`OtherDeviceUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "43e9c864-a7e4-4310-8031-d9817d2c5a43",
  "device_mac": "device_mac2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

