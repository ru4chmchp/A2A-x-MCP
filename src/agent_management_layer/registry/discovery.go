package registry

import (
	"strings"

	"github.com/ru4chmchp/A2A-x-MCP/agentcard"
)

type DiscoveryCriteria struct {
	Tag        string
	Category   string
	Capability string
}

func (r *AgentRegistry) FindByTag(tag string) []*agentcard.AgentCard {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*agentcard.AgentCard

	for _, card := range r.agents {

		if hasTag(card, tag) {
			result = append(result, card)
		}
	}
	return result
}

func (r *AgentRegistry) FindByCategory(category string) []*agentcard.AgentCard {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*agentcard.AgentCard

	for _, card := range r.agents {
		if strings.EqualFold(card.Metadata.Category, category) {
			result = append(result, card)
		}
	}
	return result
}

func (r *AgentRegistry) FindByCapability(functionName string) []*agentcard.AgentCard {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*agentcard.AgentCard
	for _, card := range r.agents {
		// for _, cap := range card.Capabilities {
		// 	if strings.EqualFold(cap.FunctionName, functionName) {
		// 		result = append(result, card)
		// 		break
		// 	}
		// }

		if hasCapability(card, functionName) {
			result = append(result, card)
		}
	}
	return result
}

func (r *AgentRegistry) FindByCriteria(mtdt DiscoveryCriteria) []*agentcard.AgentCard {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*agentcard.AgentCard

	for _, card := range r.agents {
		if mtdt.Tag != "" && !hasTag(card, mtdt.Tag) {
			continue
		}
		if mtdt.Category != "" && !strings.EqualFold(card.Metadata.Category, mtdt.Category) {
			continue
		}
		if mtdt.Capability != "" && !hasCapability(card, mtdt.Capability) {
			continue
		}
		result = append(result, card)
	}
	return result
}

func hasTag(card *agentcard.AgentCard, tag string) bool {
	// hàm strings.EqualFold sẽ so sánh hai chuỗi có khác nhau hay không, bỏ qua sự khác biệt về chữ hoa chữ thường
	for _, t := range card.Metadata.Tags {
		if strings.EqualFold(t, tag) {
			return true
		}
	}
	return false
}

func hasCapability(card *agentcard.AgentCard, capability string) bool {
	for _, c := range card.Capabilities {
		if strings.EqualFold(c.FunctionName, capability) {
			return true
		}
	}
	return false
}
