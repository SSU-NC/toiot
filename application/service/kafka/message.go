package kafka

const (
	Init = iota
	NewNode
	UpdateNode
	DeleteNode
	NewSensor
	DeleteSensor
)

type KafkaMessage struct {
	Type int         `json:"type"`
	Msg  interface{} `json:"message"`
}
