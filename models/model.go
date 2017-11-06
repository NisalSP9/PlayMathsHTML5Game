package models



type USER struct {

	UserID int `db:"userid"json:"userid"`
	Username string `db:"username"json:"username"`
	Password string `db:"password"json:"password"`
	Level int `db:"level"json:"level"`


}
