package db

import (
	"database/sql"
	"fmt"
)

type PhoneNumber struct {
	Id     int
	Number string
}

func InsertData(db *sql.DB, data []string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Could not begin Transaction: %w", err)
	}
	defer tx.Rollback()
	query := "INSERT INTO phone_numbers (mobile_number) VALUES ($1);"
	for _, value := range data {
		_, err := tx.Exec(query, value)
		if err != nil {
			return fmt.Errorf("Could not insert the data: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Could not commit transaction: %w", err)
	}
	fmt.Println("Data Inserted :)")
	return nil
}

func GetPhoneNumbers(db *sql.DB) ([]PhoneNumber, error) {
	query := "SELECT id, mobile_number FROM phone_numbers"
	var data []PhoneNumber
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("could not query data: %w", err)
	}
	for rows.Next() {
		var id int
		var phone_number string
		err = rows.Scan(&id, &phone_number)
		if err != nil {
			return nil, fmt.Errorf("error in the scanning row: %w", err)
		}
		data = append(data, PhoneNumber{
			Id:     id,
			Number: phone_number,
		})
	}
	return data, nil
}

func DeleteEntries(db *sql.DB, ids []int) error {
	query := "DELETE FROM phone_numbers WHERE id=$1"
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Could not begin Transaction: %w", err)
	}
	defer tx.Rollback()
	for _, value := range ids {
		_, err := tx.Exec(query, value)
		if err != nil {
			return fmt.Errorf("Could not delete the data: %w", err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Could not commit transaction: %w", err)
	}
	return nil
}

func UpdateEntry(db *sql.DB, id int, number string) error {
	query := "UPDATE phone_numbers SET mobile_number = $2 WHERE id = $1"
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Could not begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Exec(query, id, number)
	if err != nil {
		return fmt.Errorf("Could not update the row: %w", err)
	}
	return nil
}

func GetPhoneNumberByNumber(db *sql.DB, number string) (*PhoneNumber, error) {
	query := "SELECT id, mobile_number FROM phone_numbers WHERE mobile_number = $1 LIMIT 1"
	var phoneNumber PhoneNumber
	err := db.QueryRow(query, number).Scan(&phoneNumber.Id, &phoneNumber.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Error querying phone number: %w", err)
	}
	return &phoneNumber, nil
}
