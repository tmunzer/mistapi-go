
# Sso Role Org

SSO Role response

## Structure

`SsoRoleOrg`

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
| `Privileges` | [`[]models.PrivilegeOrg`](../../doc/models/privilege-org.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name2",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "privileges": [
    {
      "role": "admin",
      "scope": "site",
      "site_id": "00002366-0000-0000-0000-000000000000",
      "sitegroup_id": "000006ce-0000-0000-0000-000000000000",
      "views": "lobby_admin"
    }
  ],
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 65.12,
  "for_site": false,
  "id": "0000205a-0000-0000-0000-000000000000",
  "modified_time": 13.84,
  "msp_id": "00001e56-0000-0000-0000-000000000000"
}
```

