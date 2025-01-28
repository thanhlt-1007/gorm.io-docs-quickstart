package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "fmt"
)

type Product struct {
    gorm.Model
    Code string
    Price uint
}

func main() {
    db, err := gorm.Open(
        sqlite.Open("db.sqlite"),
        &gorm.Config{},
    )

    if err != nil {
        panic(fmt.Sprintf("gorm.Open err: %v\n", err))
    }

    // Migrate
    db.AutoMigrate(&Product{})

    // Create
    db.Create(&Product{
        Code: "D42",
        Price: 100,
    })

    // Read
    var product Product
    db.First(&product, 1)
    db.Find(&product, "code = ?", "D42")

    // Update - update product's price to 200
    db.Model(&product).Update("Price", 200)
    db.Model(&product).Updates(Product{
        Code: "F42",
        Price: 100,
    })
    db.Model(&product).Updates(map[string]interface{} {
        "Code": "F42",
        "Price": 200,
    })

    // Delete
    db.Delete(&product, 1)
}
