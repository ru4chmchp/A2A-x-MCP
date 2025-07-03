# UI layer

Thì mính sẽ vừa thiết kế vừa tìm hiểu các kiến thức liên quan :v

Đầu tiến là tạo project thông qua vite như sau :

```javascript
    npm create vite@latest ui_layer -- --template react-ts
```

Project ban đầu chưa có gì cả nên mình sẽ tạo cấu trúc thư mục, import tailwindcss vào.

```javascript
    npm install tailwindcss @tailwindcss/vite
```

và config `@tailwindcss/vite` trong file vite.config.ts

```typescript
import { defineConfig } from 'vite'import tailwindcss from '@tailwindcss/vite'
export default defineConfig({
    plugins: [    tailwindcss(),
    ],
})
```

**Cấu trúc project sẽ bao gồm các phần sau**

- public
  - assets # đây là thư mục chứa ảnh, logo, icon
- src/
  - components/ # chứa các component thu nhỏ như button, input, card ...
  - pages/ # chứa các trang như dashboard, ...
  - layouts/ # chứa layout của trang web nhwu header, sidebar,...
  - libs/ # chứa api, utils
  - hooks/ # chứa custom hook (useSearch, use Alert)
  - stores/ # chứa Redux store

Cấu trúc trên sẽ mở rộng tùy vào khả năng và giai đoạn phát triển.

Sau tất cả các setup cơ bản trên thì mình sẽ đi vào phần react/vite, chúng ta cần hiểu một chút về vite với vài câu lệnh cơ bản sau :

- Chúng ta đã biết lệnh tạo project react với vite ở câu lệnh trên giờ thì chúng ta sẽ đi vào các lệnh cơ bản sau :

```javascript
    npm run dev
```

Thì dòng lệnh này sẽ chạy server dev tại localhost:5173 mặc định, nếu bạn muốn chạy ở port khác thì vẫn có thể thay đổi thông qua thay đổi trong vite.config.ts

```typescript
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";
// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  server: {
    port: 5124, // Thay đổi thành port bạn muốn
  },
});
```

hoặc có thể chạy bằng lệnh sau, nhưng chỉ có tác dụng tạm thời cho lần chạy đó, không lưu vào config file, muốn chạy vv thì dùng cái trên nha

```bash
  npx vite --port 3001
```

```bash
  npm run build
```

lệnh này sẽ build production ra thư mục dist/ dùng để deploy khi đã hoàn thành.

```bash
  npm run preview
```

lệnh trên sẽ xem bản build thử giống như bản deploy

ở trên là một vài lệnh cơ bản mà mình muốn bạn hiểu để có thể triển khai project một cách cơ bản.

- Một số keyword cần biết trong react
  - Components :
    - Component : là một khối UI trong react, có thể là một nút, một thẻ, một trang,... Các tên component phải viết hoa chữ cái đầu tiên.
    - JSX : là một cú pháp mở rộng của javascript cho phép viết HTML trong js.
    - Props vs State:
      - Props : là tham số truyền từ component cha, 0 thể thay đổi, mục đích là truyền data giữa các component
      - State : là dữ liệu nội bộ của component, có thể thay đổi qua setState/useState, mục đích quản lí trạng thái bên trong.
    - Conditional Rendering : hiển thị theo điều kiện dùng if, ternary, && hoặc return sớm gì đó
    - Composition : là cách xây dựng các component phức tạp bằng cách lông các component nhỏ vào.
    - Functional component : là một hàm javascript(or typescript) trả về JSX/TSX để hiển thị giao diện
  - Rendering
    - Component Lifecycle
    - Lists and Keys
    - Render Props
    - Refs
    - Events
    - High Order Components
  - Hooks
    - Basic hooks
      - useState : dùng để lưu trạng thái nội bộ giống this.state trong class component
      - useEffect : dùng dẻ chạy tác vu phụ sau khi render, gọi API, lắng nghe sự kiện, thay đổi tiêu đề tab.
