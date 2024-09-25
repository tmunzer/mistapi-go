
# Privilege Org

Privilieges settings

## Structure

`PrivilegeOrg`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | if `scope`==`org` |
| `Role` | [`models.PrivilegeOrgRoleEnum`](../../doc/models/privilege-org-role-enum.md) | Required | access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write` |
| `Scope` | [`models.PrivilegeOrgScopeEnum`](../../doc/models/privilege-org-scope-enum.md) | Required | enum: `org`, `site`, `sitegroup` |
| `SiteId` | `*uuid.UUID` | Optional | if `scope`==`site` |
| `SitegroupId` | `*uuid.UUID` | Optional | if `scope`==`sitegroup` |

## Example (as JSON)

```json
{
  "org_id": "0000270e-0000-0000-0000-000000000000",
  "role": "write",
  "scope": "sitegroup",
  "site_id": "0000169c-0000-0000-0000-000000000000",
  "sitegroup_id": "00002114-0000-0000-0000-000000000000"
}
```

