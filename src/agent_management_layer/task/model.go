package task

import "time"

type TaskStatus string

const (
	StatusCreated   TaskStatus = "created"
	StatusRunning   TaskStatus = "in_progress"
	StatusCompleted TaskStatus = "completed"
	StatusFailed    TaskStatus = "failed"
)

type Artifact struct {
	Type      string    `json:"type"` // loại result trả về ví dụ json, text, ...
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type Task struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`   // subdomain, ip, ...
	Params    map[string]interface{} `json:"params"` // tham số đầu vào động nên dùng map[string]interface{}
	Status    TaskStatus             `json:"status"`
	Result    []Artifact             `json:"result"`
	Error     *string                `json:"error,omitempty"` // lưu lỗi nếu task thất bại, không dùng string mà dùng con trỏ *string vì để làm rõ hơn lỗi trả về, nếu để *string mà lỗi rỗng thì nó sẽ thể hiện rõ hơn &"" hoặc nếu không có lỗi trả về thì nó sẽ là nil, nó sẽ bị lượt bỏ qua trong JSON output, nếu string ngay cả chuỗi rỗng thì nó sẽ hiện trong JSON là "" không thể phân biệt được
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	AgentID   string                 `json:"agent_id"` // id agent làm task này
}
