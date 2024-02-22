package domain

type FindByUserIDAndJwtTokenParams struct {
	UserID   int
	JwtToken string
}

type DeactivateAuthTokenParams struct {
	UserID   int
	JwtToken string
}
