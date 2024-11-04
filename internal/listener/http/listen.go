package http

import "github.com/Nadzyoki/wowmq/internal/models"

func (hl *HTTPListener) GetChannelRawMessages() chan models.RawMessage {
	return hl.ch
}