package agentcard

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Load agent card từ file JSON vào cấu trúc AgentCard
func LoadAgentCard(filePath string) (*AgentCard, error) {
	// đọc dữ liệu từ file JSON và lưu vào data
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, fmt.Errorf("thất bại khi đọc file agentcard, lỗi đường dẫn khum đúng or, ...: %w", err)

	}
	var card AgentCard
	// điền dữ liệu từ biến data vào cấu trúc AgentCard
	if err := json.Unmarshal(data, &card); err != nil {
		return nil, fmt.Errorf("thất bại khi truyền dữ liệu từ JSON vào struct Agent %w", err)

	}

	return &card, nil
}

// Lưu agentcard thành file json
func SaveAgentCard(card *AgentCard, filePath string) error {
	// đọc dữ liệu từ agentcard sau đo lưu vào file json
	data, err := json.MarshalIndent(card, "", "  ")

	if err != nil {
		return fmt.Errorf("thất bại khi truyền agentcard thành file json %w", err)
	}
	// 0644 đây là permission , số 0 ở đầu biểu diễn cho quyền hạn của tệp trên linux và có nghĩa là không có quyền hạn đặc biệt nào
	// rwx thì r là 4 w là 2 x là 1
	// 0644 : 6 là quyền chủ sở hữu có thể rw-, 4 là quyền của nhóm r--, 4 là quyền của others r--
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("thất bại khi write file json %w", err)
	}
	return nil
}

// đây là hàm load nguyên một thư mục agentcard
func LoadAllAgentCards(dirPath string) ([]*AgentCard, error) {
	// filepath.Glob là một hàm trong path/filepath sẽ nhận một chuỗi đường dẫn và trả về một []string chưa tất cả đường dẫn tệp và thư mục khớp với chuỗi đường dẫn đó
	// filepath.Join là hàm nối các phần từ đương dẫn lại với nhau nếu dirpath là registry/cards thì sẽ tạo ra registry/cards/*.json
	files, err := filepath.Glob(filepath.Join(dirPath, "*.json"))

	if err != nil {
		return nil, fmt.Errorf("thất bại khi đọc files agentcard : %w", err)
	}

	var cards []*AgentCard

	for _, file := range files {
		card, err := LoadAgentCard(file)
		if err != nil {
			fmt.Printf("thất bại khi load agentcard từ file %s : %v\n", file, err)
			continue
		}
		cards = append(cards, card)

	}

	if len(cards) == 0 {
		return nil, errors.New("không có cái agentcard nào được tìm thấy or chúng bị lỗi hết :v")
	}
	return cards, nil
}
