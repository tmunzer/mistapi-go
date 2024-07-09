
# Org Apitoken

Org API Token

## Structure

`OrgApitoken`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedBy` | `models.Optional[string]` | Optional | email of the token creator / null if creator is deleted |
| `CreatedTime` | `*int` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Key` | `*string` | Optional | - |
| `LastUsed` | `models.Optional[int]` | Optional | - |
| `Name` | `models.Optional[string]` | Optional | name of the token |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Privileges` | [`[]models.PrivilegeOrg`](../../doc/models/privilege-org.md) | Optional | list of privileges the token has on the orgs/sites<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `SrcIps` | `[]string` | Optional | to restrict where the API can be called from |

## Example (as JSON)

```json
{
  "created_by": "user@mycorp.com",
  "created_time": 1626875902,
  "id": "497f6eca-6276-4993-bfeb-53ecbbba6f08",
  "key": "1qkb...QQCL",
  "last_used": 1690115110,
  "name": "org_token_xyz",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "privileges": [
    {
      "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
      "role": "admin",
      "scope": "org"
    }
  ],
  "src_ips": [
    "198.51.0.42/32"
  ]
}
```

