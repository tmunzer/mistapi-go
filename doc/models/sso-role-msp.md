
# Sso Role Msp

SSO Role response

## Structure

`SsoRoleMsp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `MspId` | `*uuid.UUID` | Optional | - |
| `Name` | `string` | Required | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Privileges` | [`[]models.PrivilegeMsp`](../../doc/models/privilege-msp.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "privileges": [
    {
      "org_id": "00000cc8-0000-0000-0000-000000000000",
      "org_name": "org_name6",
      "orggroup_id": "000010e2-0000-0000-0000-000000000000",
      "role": "admin",
      "scope": "org",
      "views": "super_observer"
    }
  ],
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 45.5,
  "for_site": false,
  "id": "00001d60-0000-0000-0000-000000000000",
  "modified_time": 33.46,
  "msp_id": "00001b5c-0000-0000-0000-000000000000"
}
```

