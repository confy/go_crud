# go_crud
simple crud API with [echo](https://pkg.go.dev/github.com/labstack/echo/v4) and [gorm](https://pkg.go.dev/gorm.io/gorm)

## Model
    
```go
type Data struct {
	gorm.Model
	
	Value     uint `json:"value"` // All other values are added by gorm

	ID        uint   `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

## Routes

### POST `/data`

### GET `/data/:id`
### PUT `/data/:id`
### DELETE `/data/:id`
### GET `/average`
