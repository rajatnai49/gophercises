package main

import (
	"bufio"
	"database/sql"
	"os"
	"regexp"
	"strings"

	"github.com/rajatnai49/phone/db"
)

func main() {
	db_con, err := db.DBConnect()
	defer db_con.Close()
	if err != nil {
		panic(err)
	}
    // Insert Data Into
	// data := getData("input")
	// err = db.InsertData(db_con, data)
	// if err != nil {
	// 	panic(err)
	// }
	phoneNumberNormalizer(db_con)
}

func phoneNumberNormalizer(db_con *sql.DB) {
	data, err := db.GetPhoneNumbers(db_con)
	if err != nil {
		panic(err)
	}
	var re = regexp.MustCompile(`(?m)([0-9]+)`)
	for i, value := range data {
		match := re.FindAllSubmatch([]byte(value.Number), -1)
		var matchedString []string
		for _, v := range match {
			matchedString = append(matchedString, string(v[0]))
		}
		number_string := strings.Join(matchedString, "")
		data[i].Number = number_string
	}

    var duplicate_key []int
	for _, v := range data {
		err = db.UpdateEntry(db_con, v.Id, v.Number)
		if err != nil {
            duplicate_key = append(duplicate_key, v.Id)
		}
	}
    err = db.DeleteEntries(db_con, duplicate_key)
    if err != nil {
        panic(err)
    }
}

func getData(file string) []string {
	var data []string
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	return data
}
