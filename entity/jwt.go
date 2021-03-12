package entity

type JWT struct {
	accessToken      string `json:"access_token"` //tell the access information of the user
	idToken          string `json:"id_token"`
	expiresIn        int    `json:"expires_in"`
	refreshExpiresIn int    `json:"refresh_expires_in"`
	refreshToken     string `json:"refresh_token"`
	tokenType        string `json:"token_type"`
	notBeforePolicy  int    `json:"not-before-policy"`
	sessionState     string `json:"session_state"`
	scope            string `json:"scope"`
}
