package mysql

import (
	"database/sql"
	"fmt"
	"game/entity"
	"time"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	var createdAt time.Time

	row := d.db.QueryRow(`SELECT * FROM user WHERE phone_number = ?`, phoneNumber)

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			return true, nil
		}
		return false, fmt.Errorf("can't scan query result: %w", err)
	}
	return false, nil
}

func (d *MySQLDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`INSERT INTO user (name, phone_number) VALUES (?,?)`, u.Name, u.PhoneNumber)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %w", err)
	}
	// error is always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)
	fmt.Println(u)
	return u, nil
}
