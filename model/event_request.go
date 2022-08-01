package model

import "event-api/shared"

type EventRequest struct {
	Type               EventType `json:"eventType"`
	WebsiteUrl         string
	SessionId          string
	ResizeFrom         *shared.Dimension
	ResizeTo           *shared.Dimension
	Pasted             bool
	FormId             InputType
	FormCompletionTime *int `json:"timeTaken"`
}

type EventType string

const (
	InputPasted   EventType = "copyAndPaste"
	ScreenResized EventType = "screenResize"
	TimeTaken     EventType = "timeTaken"
)

type InputType string

const (
	Email      InputType = "inputEmail"
	CardNumber InputType = "inputCardNumber"
	Cvv        InputType = "inputCVV"
)
