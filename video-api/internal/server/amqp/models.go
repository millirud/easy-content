package amqp

type videoTransformPayload struct {
	Filename string `json:"filename"`
	Bucket   string `json:"bucket"`
}
