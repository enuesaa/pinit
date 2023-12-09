package serve

type ConfigResponse struct {
	Token string `json:"token"`
}
type ChatRequest struct {
	Message string `json:"message"`
}
type ChatResponse struct {
	Message string `json:"message"`
}
type Action struct {
	Name     string `json:"name"`
	Template string `json:"template"`
}
type ActionResponse struct {
	Items []Action `json:"items"`
}
