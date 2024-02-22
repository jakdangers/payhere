package user

const createUserQuery = "INSERT INTO `users` (mobile_id, password, use_type) VALUES (?, ?, ?)"

const findUserByMobileIDQuery = "SELECT id, mobile_id, password, use_type FROM `users` WHERE mobile_id = ?"
