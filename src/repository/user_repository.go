package repository

type User struct {
	Id   int
	Name string
	Role string
}

func GetUsersByRole(role string) []User {
	var users []User
	rows, errQuery := db.Query("SELECT id, name, role FROM jira.users WHERE role = ? ORDER BY id ASC", role)
	if errQuery != nil {
		return users
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name, &user.Role)
		users = append(users, user)
	}
	return users
}
