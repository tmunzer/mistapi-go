
# Org Crl Import File

## Structure

`OrgCrlImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `*[]byte` | Optional | a binary .crl file |
| `Json` | `*string` | Optional | json string with name for .crl file (optional) |

## Example (as JSON)

```json
{
  "file": "data:text/plain;name=dummy_file;base64,",
  "json": "json2"
}
```

