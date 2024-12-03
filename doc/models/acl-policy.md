
# Acl Policy

- for GBP-based policy, all src_tags and dst_tags have to be gbp-based
- for ACL-based policy, `network` is required in either the source or destination so that we know where to attach the policy to

*This model accepts additional fields of type interface{}.*

## Structure

`AclPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Actions` | [`[]models.AclPolicyAction`](../../doc/models/acl-policy-action.md) | Optional | - for GBP-based policy, all src_tags and dst_tags have to be gbp-based<br>- for ACL-based policy, `network` is required in either the source or destination so that we know where to attach the policy to |
| `Name` | `*string` | Optional | - |
| `SrcTags` | `[]string` | Optional | - for GBP-based policy, all src_tags and dst_tags have to be gbp-based<br>- for ACL-based policy, `network` is required in either the source or destination so that we know where to attach the policy to |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "name": "guest access",
  "actions": [
    {
      "action": "allow",
      "dst_tag": "dst_tag0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "src_tags": [
    "src_tags9",
    "src_tags0",
    "src_tags1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

