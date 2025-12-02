
# Ssr Version

## Structure

`SsrVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Default` | `*bool` | Optional | - |
| `Package` | `string` | Required | - |
| `Tags` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Version` | `string` | Required | - |

## Example (as JSON)

```json
{
  "default": false,
  "package": "package4",
  "tags": [
    "tags9",
    "tags0",
    "tags1"
  ],
  "version": "version0"
}
```

