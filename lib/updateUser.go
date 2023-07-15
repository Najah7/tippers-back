package lib

import "tippers-back/db/table"

// password以外
func UpdateUser(oldUser, newUser *table.User) *table.User {
	if oldUser.Email != "" {
		newUser.Email = oldUser.Email
	}
	if oldUser.Name != "" {
		newUser.Name = oldUser.Name
	}
	if oldUser.ProfileImageURL != "" {
		newUser.ProfileImageURL = oldUser.ProfileImageURL
	}
	if oldUser.Dream != "" {
		newUser.Dream = oldUser.Dream
	}
	if oldUser.Major != "" {
		newUser.Major = oldUser.Major
	}
	if oldUser.IsStudent != 0 {
		newUser.IsStudent = oldUser.IsStudent
	}
	if oldUser.IsEmployed != 0 {
		newUser.IsEmployed = oldUser.IsEmployed
	}
	if oldUser.PeriodOfWorkings != "" {
		newUser.PeriodOfWorkings = oldUser.PeriodOfWorkings
	}
	if oldUser.RestaurantID != 0 {
		newUser.RestaurantID = oldUser.RestaurantID
	}

	return newUser
}
