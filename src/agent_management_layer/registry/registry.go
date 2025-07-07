package registry

import (
	"errors"
	"sync"

	"github.com/ru4chmchp/A2A-x-MCP/agentcard"
)

// struct này khai báo đăng ký agent với một map trống khai báo bằng make để lưu các agent với key là agent.ID và value là con trỏ đến agentcard
// sync là packg đồng hóa cơ bản của go, RWMutex là một loại khóa lock xử lý dữ liệu đọc thường xuyên hơn là ghi nếu không có nó bạn sẽ bị data race khi đọc ghi cùng một lúc
// ví dụ dễ hiểu là nếu bạn có một tài nguyên chia sẽ như một map or một ds dữ liệu, sau đó 2 hay hai hay nhiều goroutine có gắng truy cập và sửa đổi nội dung cùng một lúc thì sẽ dẫn đến tình trạng tranh chấp dữ liệu (data race)
// và vì thế chúng ta cần Mutex, nó hoạt động như một khóa (lock), khi một goroutine muốn truy cập tài nguyên chia sẽ thì nó phải cố gắng khóa (lock)Mutex.
// Nếu một mutex đang không bị khóa bởi goroutine nào khác thì, cái goroutine hiện tại sẽ dành được khóa và có thể truy cập tài nguyên, nếu nó bị khóa bởi một goroutine khác thì cái goroutine hiện tại sẽ bị chặn và sẽ đợi cho đến khi cái goroutine đó nhả khóa ra.
// Sau khi hoàn thành công việc thì goroutine phải bắt buộc nhả khóa và cho phép các goroutine khác truy cập.
type AgentRegistry struct {
	mu     sync.RWMutex
	agents map[string]*agentcard.AgentCard
}

func NewRegistry() *AgentRegistry {
	return &AgentRegistry{
		agents: make(map[string]*agentcard.AgentCard),
	}
}

// cấu trúc func (receiver) NameFunction(parameters) Typereturn {} thì phần receiver đó khi một hàm định nghĩa với một receiver thì nó được coi là phương thức với kiểu dữ liệu đó, r là tên biến receiver dùng để truy cập các trường và method của đối tượng registry đó.

func (r *AgentRegistry) Register(card *agentcard.AgentCard) error {
	// phương thức này như đã nói trên nó sẽ cố gắng dành khóa (Lock) mutex để có thể ghi vào dữ liệu
	r.mu.Lock()
	// phương thức này sẽ nhả khóa sau khi hoàn thành việc ghi ở trên
	defer r.mu.Unlock()

	if _, exits := r.agents[card.ID]; exits {
		return errors.New("đã tồn tại agent rồi :v")
	}
	r.agents[card.ID] = card
	return nil
}

func (r *AgentRegistry) Get(ID string) (*agentcard.AgentCard, error) {
	// phương thức này giành lấy khóa đọc, khi nhiều goroutine cần đọc dữ liệu, chúng có thể cùng giữ khóa đọc và truy cập, đảm bảo không có goroutine nào khác ghi dữ liệu
	r.mu.RLock()
	// như trên thôi, phương thức nhả khóa đọc
	defer r.mu.RUnlock()

	card, exits := r.agents[ID]

	if !exits {
		return nil, errors.New("không tìm thấy ID của agentcard này")
	}
	return card, nil
}

func (r *AgentRegistry) Delete(ID string) error {
	// xóa là ghi nên dùng Lock
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exits := r.agents[ID]; !exits {
		return errors.New("không tìm thấy ID của agentcard này ")
	}
	delete(r.agents, ID)
	return nil
}

func (r *AgentRegistry) List() []*agentcard.AgentCard {
	// chỉ đọc nên dùng RLock
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]*agentcard.AgentCard, 0, len(r.agents))

	for _, card := range r.agents {
		list = append(list, card)
	}
	return list
}
