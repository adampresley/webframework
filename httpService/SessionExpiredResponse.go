package httpService

type SessionExpiredResponse struct {
	Success        bool `json:"success"`
	SessionExpired bool `json:"sessionExpired"`
}

/*
Creates a new SessionExpiredResponse structure.
*/
func NewSessionExpiredResponse() SessionExpiredResponse {
	return SessionExpiredResponse{
		Success:        false,
		SessionExpired: true,
	}
}
