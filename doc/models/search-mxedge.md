
# Search Mxedge

## Structure

`SearchMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Distro` | `*string` | Optional | - |
| `LastSeen` | `*float64` | Optional | - |
| `Model` | `*string` | Optional | - |
| `MxclusterId` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `MxedgeId` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Name` | `*string` | Optional | The name of the tunnel |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `TuntermVersion` | `*string` | Optional | - |
| `Uptime` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "model": "ME-VM",
  "mxcluster_id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "mxedge_id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "me-vm-1",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "distro": "distro0",
  "last_seen": 135.34
}
```

