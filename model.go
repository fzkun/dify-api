package dify_api

type DifyGenerateRequest struct {
	Inputs       map[string]any `json:"inputs"`        //dify自定义参数
	User         string         `json:"user"`          //调用用户(用于统计)
	ResponseMode string         `json:"response_mode"` //SSE:streaming blocking
}

type DifyGenerateResponse struct {
	TaskId        string `json:"task_id"`
	WorkflowRunId string `json:"workflow_run_id"`
	Data          struct {
		Id         string `json:"id"`
		WorkflowId string `json:"workflow_id"`
		Status     string `json:"status"`
		Outputs    struct {
			Text string `json:"text"`
		} `json:"outputs"`
		Error       string  `json:"error"`
		ElapsedTime float64 `json:"elapsed_time"`
		TotalTokens int     `json:"total_tokens"`
		TotalSteps  int     `json:"total_steps"`
		CreatedAt   int     `json:"created_at"`
		FinishedAt  int     `json:"finished_at"`
	} `json:"data"`
}

type DifySSEResponse struct {
	Event         string `json:"event"`
	WorkflowRunId string `json:"workflow_run_id"`
	TaskId        string `json:"task_id"`
	Data          struct {
		Id             string `json:"id"`
		WorkflowId     string `json:"workflow_id"`
		SequenceNumber int    `json:"sequence_number"`
		Status         string `json:"status"`
		Outputs        struct {
			Text string `json:"text"`
		} `json:"outputs"`
		Error       interface{} `json:"error"`
		ElapsedTime float64     `json:"elapsed_time"`
		TotalTokens int         `json:"total_tokens"`
		TotalSteps  int         `json:"total_steps"`
		CreatedBy   struct {
			Id   string `json:"id"`
			User string `json:"user"`
		} `json:"created_by"`
		CreatedAt       int           `json:"created_at"`
		FinishedAt      int           `json:"finished_at"`
		ExceptionsCount int           `json:"exceptions_count"`
		Files           []interface{} `json:"files"`
	} `json:"data"`
}
