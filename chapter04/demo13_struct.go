package main

type User struct {
	ID     int
	Name   string
	Age    int
	Leader *User
}

/*
func main() {

	var user User
	user.ID = 1
	user.Name = "tom"
	user.Age = 20

	fmt.Printf("%#v\n", user)

	var u2 *User = &user
	u2.Age += 1

	fmt.Printf("%#v\n", user)
}
*/
