
# Sso Role Org

SSO Role response

*This model accepts additional fields of type interface{}.*

## Structure

`SsoRoleOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional | - |
| `Name` | `string` | Required | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Privileges` | [`[]models.PrivilegeOrg`](../../doc/models/privilege-org.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SitegroupId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
  "name": "name4",
  "org_id": "60f6bfdb-2f45-4022-8e2a-e00d977953fe",
  "privileges": [
    {
      "org_id": "00000cc8-0000-0000-0000-000000000000",
      "role": "admin",
      "scope": "sitegroup",
      "site_id": "00002366-0000-0000-0000-000000000000",
      "sitegroup_id": "000006ce-0000-0000-0000-000000000000",
      "view": "view4",
      "views": [
        "location",
        "lobby_admin",
        "switch_admin"
      ]
    }
  ],
  "site_id": "580eb9a7-e6cb-4d05-aa20-4c60808fb011",
  "sitegroup_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "created_time": 76.64,
  "for_site": false,
  "modified_time": 2.32,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

