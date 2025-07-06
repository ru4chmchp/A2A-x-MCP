package agentcard

// thì lớp này sẽ triển khai định nghĩa về agent card, các chức năng,
type AgentCard struct {
	ID           string               `json:"id"`
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	Version      string               `json:"version"`
	Capabilities []FunctionCapability `json:"capabilities"` // này là các hàm mà agent này có
	Auth         AuthInfo             `json:"auth"`         // yêu cầu xác thực
	Metadata     Metadata             `json:"metadata"`     // đây là dữ liệu mô tả các thông tin thêm về agent này
}

type FunctionCapability struct {
	FunctionName string      `json:"function_name"`
	Description  string      `json:"description"`
	Params       []ParamSpec `json:"params"`
	Returns      []ParamSpec `json:"returns"`
}

type ParamSpec struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type AuthInfo struct {
	AuthRequired bool     `json:"auth_required"`
	Methods      []string `json:"methods"` // ví dụ như các api key hay Oauth2
	Scopes       []string `json:"scopes"`  // read, write, execute, ...
}

type Metadata struct {
	Category string   `json:"category"` // Cái này dùng để detect gì đó .... Recon, burforcing....
	Tags     []string `json:"tags"`     // dns, domain, subdomain,...
	Creator  string   `json:"creator"`  // tên tui
}
