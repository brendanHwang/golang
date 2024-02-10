package page6

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authentication: Basic %s:%s",
		authI.username,
		authI.password,
	)
}

func StructMethod() {
	authI := authenticationInfo{
		username: "username",
		password: "password",
	}
	fmt.Println(authI.getBasicAuth())
}
