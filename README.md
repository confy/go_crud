# go_crud
simple crud API with echo and gorm

## Model
    
```go
type Data struct {
	gorm.Model

	Value     string `json:"value"` // All other values are added by gorm

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

