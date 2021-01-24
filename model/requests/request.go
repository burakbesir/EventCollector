package requests

type EventCreateRequest struct {
	SessionId string `json:"sessionId" validate:"required"`
	ClickId string `json:"clickId" validate:"required"`
}
