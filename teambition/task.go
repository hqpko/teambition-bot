package teambition

type Task struct {
	ID         string `json:"_id,omitempty"`
	ExecutorId string `json:"_executorId,omitempty"`
	isDone     string `json:"is_done,omitempty"`
	Content    string `json:"content,omitempty"`
	Stage      struct {
		ID   string `json:"_id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"stage,omitempty"`
}
