package databases

import (
	"test-b/config"
	"test-b/models"
)

// function database untuk menambahkan user baru (registrasi)
func CreateNasabah(Nasabah *models.Nasabah) (interface{}, error) {
	if err := config.DB.Create(&Nasabah).Error; err != nil {
		return nil, err
	}
	res := models.ResNasabah{
		ID:           Nasabah.ID,
		NamaLengkap:  Nasabah.NamaLengkap,
		Alamat:       Nasabah.Alamat,
		NoKtp:        Nasabah.NoKtp,
		NoTelepon:    Nasabah.NoTelepon,
		TanggalLahir: Nasabah.TanggalLahir,
		TempatLahir:  Nasabah.TempatLahir,
	}

	return res, nil
}

func GetAllNasabah() (interface{}, error) {
	res := []models.ResNasabah{}
	query := config.DB.Table("nasabahs").Select("*").Where("nasabahs.deleted_at IS NULL").Find(&res)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return res, nil
}

func GetByKtpNasabah(ktp int) (interface{}, error) {
	res := []models.ResNasabah{}
	query := config.DB.Table("nasabahs").Select("*").Where("nasabahs.deleted_at IS NULL && nasabahs.no_ktp =? ", ktp).Find(&res)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return res, nil
}

func DeleteByIdNasabah(id int) (interface{}, error) {
	nasabah := models.Nasabah{}
	query := config.DB.Where("id = ?", id).Delete(&nasabah)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return nasabah, nil
}

func UpdateNasabah(id int, Nasabah *models.Nasabah) (interface{}, error) {
	query := config.DB.Where("id = ?", id).Updates(&Nasabah)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	res := models.ResNasabah{
		ID:           uint(id),
		NamaLengkap:  Nasabah.NamaLengkap,
		Alamat:       Nasabah.Alamat,
		NoKtp:        Nasabah.NoKtp,
		NoTelepon:    Nasabah.NoTelepon,
		TanggalLahir: Nasabah.TanggalLahir,
		TempatLahir:  Nasabah.TempatLahir,
	}
	return res, nil
}
