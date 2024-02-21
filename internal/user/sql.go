package user

const createUserQuery = "INSERT INTO `users` (user_id, password, user_type) VALUES (?, ?, ?)"

const findByUserIDQuery = "SELECT id, user_id, password, user_type FROM `users` WHERE user_id = ?"
