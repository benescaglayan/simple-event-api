package entity

import "event-api/shared"

type EventEntity struct {
	SessionId          string
	WebsiteUrl         string
	ResizeFrom         *shared.Dimension
	ResizeTo           *shared.Dimension
	CopyAndPaste       map[string]bool
	FormCompletionTime *int
}
