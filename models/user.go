package models

import "time"

// urutannya : nama field, tipe data, key di jsonnya
// pakai automigrate, kita bisa tambah field sesuka kita
// namun ketika di code dihapus, di db ngga langsung kehapus
type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
