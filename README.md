# go_crud
simple crud API with [echo](https://pkg.go.dev/github.com/labstack/echo/v4) and [gorm](https://pkg.go.dev/gorm.io/gorm)

## Model
    
```go
type Data struct {
	Value     uint64 `json:"value"`
	
	ID        uint64   `gorm:"primaryKey"`
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

### GET `/hello`
