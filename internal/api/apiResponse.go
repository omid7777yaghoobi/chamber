package api

import (
	"encoding/json"
)

type ApiResponse struct {
	messageAttributes MessageAttributes
	messageBody       MessageBody
}

type MessageAttributes struct {
	CorrelationId string
}

type MessageBody interface {
	Serialize() ([]byte, error)
	// Deserialize(data []byte) error
}

func NewApiResponse() *ApiResponse {
	return &ApiResponse{}
}

func (ar *ApiResponse) WithMessageAttributes(msgAttrs MessageAttributes) *ApiResponse {
	ar.messageAttributes = msgAttrs
	return ar
}

func (ar *ApiResponse) WithMessageBody(msgbody MessageBody) *ApiResponse {
	ar.messageBody = msgbody
	return ar
}

func (ar *ApiResponse) Serialize() ([]byte, error) {
	return json.Marshal(ar)
}

func NewMessageAttributes(correlationId string) *MessageAttributes {
	return &MessageAttributes{CorrelationId: correlationId}
}
