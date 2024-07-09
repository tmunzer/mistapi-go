
# Wxlan Rule

WXlan

## Structure

`WxlanRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.WxlanRuleActionEnum`](../../doc/models/wxlan-rule-action-enum.md) | Optional | type of action, allow / block |
| `ApplyTags` | `[]string` | Optional | - |
| `BlockedApps` | `[]string` | Optional | blocked apps (always blocking, ignoring action), the key of Get Application List |
| `CreatedTime` | `*float64` | Optional | - |
| `DstAllowWxtags` | `[]string` | Optional | tag list to indicate these tags are allowed access |
| `DstDenyWxtags` | `[]string` | Optional | tag list to indicate these tags are blocked access |
| `Enabled` | `*bool` | Optional | **Default**: `true` |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Order` | `int` | Required | the order how rules would be looked up, > 0 and bigger order got matched first, -1 means LAST, uniqueness not checked<br>**Constraints**: `>= 1` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SrcWxtags` | `[]string` | Required | tag list to determine if this rule would match |
| `TemplateId` | `*uuid.UUID` | Optional | Only for Org Level WxRule |

## Example (as JSON)

```json
{
  "action": "allow",
  "blocked_apps": [
    "mist",
    "all-videos"
  ],
  "dst_allow_wxtags": [
    "fff34466-eec0-3756-6765-381c728a6037",
    "eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
  ],
  "dst_deny_wxtags": [
    "aaa34466-eec0-3756-6765-381c728a6037",
    "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
  ],
  "enabled": true,
  "order": 1,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "src_wxtags": [
    "8bfc2490-d726-3587-038d-cb2e71bd2330",
    "3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"
  ],
  "template_id": "6aa54cbd-e039-4878-846a-04f270de8a5c",
  "apply_tags": [
    "apply_tags8",
    "apply_tags9",
    "apply_tags0"
  ],
  "created_time": 51.9
}
```

