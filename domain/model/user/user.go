package user

import (
	"fmt"
	"reflect"

	"github.com/SakataAtsuki/itddd-03-entity/iterrors"
)

type User struct {
	userId UserId
	name   string
}

func NewUser(userId UserId, name string) (*User, error) {
	user := new(User)

	user.userId = userId

	if err := user.ChangeUserName(name); err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) ChangeUserName(name string) (err error) {
	defer iterrors.Wrap(&err, "user.ChangeUserName(%q)", name)
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if len(name) < 3 {
		return fmt.Errorf("name %v is less than three characters long", name)
	}
	user.name = name
	return nil
}

func (user *User) Equals(other *User) bool {
	return reflect.DeepEqual(user.userId, other.userId)
}
