# 🌐 AI-Powered OSINT Research Agent System

Thì đây sẽ là dự án triển khai integration among A2A and MCP, tất cả source code mình sẽ viết bằng `go` để vừa luyện `go` vừa triển khai dự án thử.

## System architecture

Hệ thống tích hợp này sẽ gồm có 5 layer chính như sau :

<div align="center">

```mermaid
graph TD
  O[System] --> A[User interface layer]
  A --> B[Agent management layer]
  B --> C[Core protocol layer]
  C --> D[Tool integration layer]
  D --> E[Security layer]
```

</div>

Chúng ta có thể thấy đó là 5 layer này chịu trách nhiệm cho 5 nhiệm vụ khác nhau, những layer này bổ sung cho nhau, góp phần tạo nên hệ thống, thiếu một trong số trong không thể nào hoàn thành hệ thống được.

### User interface layer

Lớp giao diện này để người dùng tương tác với hệ thống, sử dụng React/Vite + Tailwind CSS để xây dựng giao diện. Chủ yếu xử lí đầu vào của người dùng là văn bản.

Đơn giản thì mình muốn xây dựng một giao diện giúp người dùng dễ dàng giao tiếp với hệ thống agent và tiếp nhận và phân tích yêu cầu người dùng để gửi xuống các lớp bên dưới.

<div align="center">

```mermaid
graph TD
    subgraph L1[UI interaction]
    A["React + Vite +Tailwind CSS"] --> B["Web application"]
    end
    subgraph L2[User request handler]
    E["User requests"] --> F["Messages"]
    end
    subgraph O[User interface layer]
    end

    O --> L1 --> L2
```

</div>

### Agent management layer

Lớp này sẽ làm nhiệm vụ quản lý agent bao gồm các vòng đời ,trạng thái, các chức năng của agent, quản lý tạo, xóa các agent card, điều phối những agent giao tiếp với nhau, những agent có thể cộng tác linh hoạt qua tự động khám quá agent và kết nối thông qua khả năng đàm phán. (mình cũng khá khó hiểu khúc khả năng đàm phán này :v)

```mermaid
graph TD
    subgraph L1[Management]
    A["Lifecycle"]
    B["State"]
    C["Functionality"]
    end
    subgraph L2[Agent card]
    T["Create"]
    Y["Delete"]
    end
    subgraph L3[Agent combination]
    D["Agent A"] --> |Call| E["Agent B"]
    end
    subgraph L4[Task coordination]
    F["Agent managemnt"] --> |divide tasks|G["Agent"]
    end
    subgraph L5[Mornitoring]
    H["Log"]
    J["Agent status"]
    end
    subgraph L6[Re-used]
    K["agent"] -->|re-used|K["agent"]
    end
    subgraph O[Agent management layer]
    end
    subgraph L7[Agent enviroment]
    U["Agent 1"]-->|A2A protocol|W["Agent ..."]
    end

    O --> L1
    O --> L2
    O --> L3
    O --> L4
    O --> L5
    O --> L6
    O --> L7


```

### Core protocol layer

Lớp này sẽ làm nhiệm vụ triển khai cơ bản giao thức A2A và MCP, cung cấp những chức năng cốt lõi như message formats, communication rules, error handling, giao tiếp với các công cụ bên ngoài theo mô hình MCP, quản lí các artifact sinh ra từ task sau khi hoàn thành, theo giỏi.

```mermaid
graph TD
    subgraph L1[A2A message format]
    A["Message reception"]
    B["Message parsing"]
    C["Message validation"]
    D["Message processing"]
    E["Response generation"]
    F["Response transmission"]
    end

    subgraph L2[MCP context protocol]
    G["Tool description"]
    H["handling function call"]
    J["result conversion"]
    end

    subgraph O[Core protocol layer]
    end
    O --> L1
    O --> L2


```

### Tool integration layer

Lớp này sẽ kết hợp các công cự và các api thông qua mcp, chúng quản lí tool và xử lí goi hàm.

```mermaid
graph TD
    subgraph L1[Tool description manager]
    A["Define Tool Basic Information"]
    B["Define Function List"]
    C["Define Parameter Schema"]
    D["Define Return Value Schema"]
    end
    subgraph L2[Function caller]
    E["Function Call Request Reception"]
    F["Parameter Validation"]
    G["Function Execution"]
    H["Result conversion"]
    J["Response return"]
    end
    subgraph L3[Schema validator]
    K["Input check"]
    L["Output check"]
    end
    subgraph L4[Result handler]
    I["Save result"]
    U["forwarded to the another agent"]
    end
    subgraph L5[External Tools Ecosystem]
    Q["Database"]
    W["Api service"]
    E["LLM provider"]
    R["Knowledge base"]
    T["File system"]
    end

    subgraph O[Tool integration layer]
    end
    O --> L1
    O --> L2
    O --> L3
    O --> L4
    O --> L5


```

### Security layer

Lớp này sẽ chịu trách nhiệm về việc bảo nhật giao tiếp giữa các agent và công cụ sử dụng, xử lý xác thực, ủy quyền, và mã hóa dữ liệu.

```mermaid
graph TD
    subgraph O[Security layer]
    end
    O --> L1["Authentication"]
    O --> L2["Authorization"]
    O --> L3["Encryption"]
    O --> L4["Access control"]


```
