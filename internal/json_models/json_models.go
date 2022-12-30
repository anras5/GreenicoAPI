package json_models

type SensorData struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Pressure    float32 `json:"pressure"`
	VOC         float32 `json:"voc"`
	UV          float32 `json:"uv"`
	Light       float32 `json:"light"`
}
