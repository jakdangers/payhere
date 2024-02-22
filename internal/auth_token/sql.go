package auth_token

const createAuthTokenQuery = "INSERT INTO `auth_tokens` (user_id, jwt_token, creation_time, expiration_time, active) VALUES (?, ?, ?, ?, ?)"

const findAuthTokenByUserIDAndJwtTokenQuery = "SELECT id, user_id, jwt_token, creation_time, expiration_time, active FROM `auth_tokens` WHERE user_id = ? AND jwt_token = ?"

const deactivateAuthTokenQuery = "UPDATE `auth_tokens` SET `active` = 0 WHERE `user_id` = ? AND `jwt_token` = ?"
