package models

import "gorm.io/gorm"

type Nasabah struct {
	gorm.Model
	NamaLengkap  string `gorm:"type:varchar(50);" form:"nama_lengkap" json:"nama_lengkap"`
	Alamat       string `gorm:"type:varchar(255);" form:"alamat" json:"alamat"`
	TempatLahir  string `gorm:"type:varchar(50);" form:"tempat_lahir" json:"tempat_lahir"`
	TanggalLahir string `gorm:"type:varchar(50);" form:"tanggal_lahir" json:"tanggal_lahir"`
	NoKtp        int    `gorm:"not null;unique;" form:"no_ktp" json:"no_ktp"`
	NoTelepon    string `gorm:"type:varchar(14);" form:"no_telepon" json:"no_telepon"`
}

type ResNasabah struct {
	ID           uint
	NamaLengkap  string
	Alamat       string
	TempatLahir  string
	TanggalLahir string
	NoKtp        int
	NoTelepon    string
}

// type Users struct {
// 	gorm.Model
// 	Email    string `gorm:"unique" json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// 	Token    string
// }

// type ResLoginUser struct {
// 	ID       uint
// 	Email    string
// 	Token    string
// 	Password string
// }
// type ResCreateUser struct {
// 	ID       uint
// 	Email    string
// 	Password string
// }
