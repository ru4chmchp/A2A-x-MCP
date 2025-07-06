package agentcard

import (
	"errors"
	"fmt"
	"regexp"
)

func ValidateAgentCard(card *AgentCard) error {
	if card.ID == "" || !isValidID(card.ID) {
		return errors.New("phải có ID của agentcard và ID phải là số như 001,002")
	}
	if card.Name == "" {
		return errors.New("phải có Name của agentcard thì người ta mới biết nó tên giè")
	}
	if card.Description == "" {
		return errors.New("có mô tả mới biết nó là cái gì chớ")
	}
	if card.Version == "" {
		return errors.New("nhập version vô, mới thì 1.0.0")
	}

	for i, cap := range card.Capabilities {
		if cap.FunctionName == "" {
			return fmt.Errorf("capability[%d]: yêu cầu tên hàm vào", i)
		}
		if cap.Description == "" {
			return fmt.Errorf("capability[%d]: mô tả ro vào", i)
		}
		for j, param := range cap.Params {
			if param.Name == "" || param.Type == "" {
				return fmt.Errorf("capability[%d]: tham số[%d] null tề", i, j)
			}
		}
		for j, ret := range cap.Returns {
			if ret.Name == "" || ret.Type == "" {
				return fmt.Errorf("capability[%d]: return[%d] null kìa", i, j)
			}
		}
	}

	if card.Metadata.Category == "" {
		return errors.New("phân loại nó ra cho dễ")
	}
	if len(card.Metadata.Tags) == 0 {
		return errors.New("gắn tag vào sau dễ tìm, dễ làm phần discovery")
	}
	if card.Metadata.Creator == "" {
		return errors.New("tên người tạo agent, mặc định là Nguyễn Hữu Dũng :V")
	}

	return nil

}

func isValidID(id string) bool {
	var idPattern = regexp.MustCompile(`^\d{3}$`)
	return idPattern.MatchString(id)
}
