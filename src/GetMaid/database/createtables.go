package database

import (
	"database/sql"
	"log"
)

func checkErr(err error, str ...string) {

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	if err != nil {
		if len(str) == 0 {
			panic(err)
		} else {
			panic(str[0])
		}
	}
}

func createTables(db *sql.DB) {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS address (
  id int(11) NOT NULL AUTO_INCREMENT,
  Pincode varchar(6) NOT NULL,
  Locality varchar(100) NOT NULL,
  City varchar(100) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS hirer (
  Hirer_id int(11) NOT NULL AUTO_INCREMENT,
  Name varchar(100) NOT NULL,
  Email varchar(100) DEFAULT NULL,
  Phone varchar(10) DEFAULT NULL,
  HouseNumber varchar(5) DEFAULT NULL,
  AddressId int(11) DEFAULT NULL,
  Password varchar(1000) NOT NULL,
  Active int(1) NOT NULL,
  PRIMARY KEY (Hirer_id),
  UNIQUE KEY Email (Email),
  UNIQUE KEY Phone (Phone),
  CONSTRAINT fk_hirer FOREIGN KEY(AddressId) REFERENCES address(id) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS maid (
  Maid_id int(11) NOT NULL AUTO_INCREMENT,
  Name varchar(100) NOT NULL,
  Email varchar(100) DEFAULT NULL,
  Phone varchar(10) DEFAULT NULL,
  AddressId int(11) DEFAULT NULL,
  Password varchar(1000) NOT NULL,
  PRIMARY KEY (Maid_id),
  Active int(1) NOT NULL,
  UNIQUE KEY Email (Email),
  UNIQUE KEY Phone (Phone),
  CONSTRAINT fk_maid FOREIGN KEY(AddressId) REFERENCES address(id) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS pincodes(
  Pincode1 varchar(15) NOT NULL,
  Pincode2 varchar(15) NOT NULL,
  UNIQUE(Pincode1, Pincode2)
)ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;`,
	}

	gos := 5

	proc := make(chan bool)
	allDone := make(chan bool)

	for i := 0; i < gos; i++ {
		go func(n int) {
			for j := n; j < len(tables); j += gos {
				_, err := db.Exec(tables[j])

				if err != nil {
					checkErr(err, err.Error())
				}
			}
			proc <- true
		}(i)
	}

	go func() {
		for i := 0; i < gos; i++ {
			<-proc
		}
		allDone <- true
	}()

	for {
		select {
		case <-allDone:
			return
		}
	}
}
