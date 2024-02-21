package auth_token

const findAuthTokenByUserIDAndJwtTokenQuery = "SELECT id, user_id, jwt_token, creation_time, expiration_time, last_access_time, is_logged_out  FROM `auth_tokens` WHERE user_id = ? AND jwt_token = ?"
