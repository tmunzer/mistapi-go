
# Radius Auth Port

RADIUS Auth Port, value from 1 to 65535, default is 1812

## Class Name

`RadiusAuthPort`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.RadiusAuthPortContainer.FromNumber(int number) |
| `string` | models.RadiusAuthPortContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.RadiusAuthPortContainer.FromNumber(1)
```

## string

### Initialization Code

#### Example

```go
value := models.RadiusAuthPortContainer.FromString("String0")
```

