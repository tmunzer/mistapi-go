
# Events Client Wan

## Structure

`EventsClientWan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `When` | `*string` | Optional | - |
| `EvType` | `*string` | Optional | - |
| `Metadata` | `*interface{}` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RandomMac` | `*bool` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Text` | `*string` | Optional | - |
| `Wcid` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "When": "2022-12-31 23:59:59.293000+00:00",
  "ev_type": "CLIENT_IP_ASSIGNED",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "text": "DHCP Ack IP 192.168.88.216",
  "wcid": "62bbfb75-10d8-49d1-dec7-d2df91624287",
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  },
  "random_mac": false
}
```

