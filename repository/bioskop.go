package repository

import (
	"bioskop/structs"
	"database/sql"
	"fmt"
)

func GetAllBioskop(db *sql.DB) (result []structs.Bioskop, err error) {
	sql := "SELECT * FROM bioskop"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var bioskop structs.Bioskop

		err = rows.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
		if err != nil {
			return
		}

		result = append(result, bioskop)
	}

	return
}

func InsertBioskop(db *sql.DB, bioskop structs.Bioskop) error {
	query := "INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating)
	return err
}

func UpdateBioskop(db *sql.DB, bioskop structs.Bioskop) error {
	query := "UPDATE bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4"
	result, err := db.Exec(query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating, bioskop.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no bioskop found with id %d", bioskop.ID)
	}

	return nil
}

func DeleteBioskop(db *sql.DB, bioskop structs.Bioskop) error {
	query := "DELETE FROM bioskop WHERE id = $1"
	result, err := db.Exec(query, bioskop.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no bioskop found with id %d", bioskop.ID)
	}

	return nil
}
