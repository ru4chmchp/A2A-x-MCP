// thư viện lucide-react là library mã nguồn mở được tạo để import các icon vào sử dụng cho tiện hơn
import { Bot, FileText, LayoutDashboard, Menu, Share2 } from "lucide-react";
import { useState } from "react";
import SidebarItem from "./SidebarItem";

const Sidebar: React.FC = () => {
  // useState thay đổi trạng thái của Sidebar xem có thu gọn hay không
  const [collapsed, setCollapsed] = useState(false);
  const toggleSidebar = () => setCollapsed(!collapsed);

  const items = [
    { to: "/", icon: <LayoutDashboard size={20} />, lable: "Home" },
    { to: "/user", icon: <Share2 size={20} />, lable: "History" },
    // { to: "/cloud", icon: <Bot size={20} />, lable: "Agents" },
    // { to: "/report", icon: <FileText size={20} />, lable: "Reports" },
  ];

  return (
    <>
      <aside
        className={`h-screen bg-gray-900 text-white p-3 rounded-xl transition-all duration-300 ${
          collapsed ? "w-[60px]" : "w-[200px]"
        }`}
      >
        <div
          className={`flex items-center mb-6 ${
            !collapsed ? "justify-between" : "justify-center"
          } `}
        >
          {!collapsed && <h1 className="text-xl font-bold">PRAISE</h1>}
          <button
            onClick={toggleSidebar}
            className="transition-transform duration-300 ease-in-out hover:scale-120"
          >
            <Menu size={20} />
          </button>
        </div>
        <nav
          className={`flex flex-col gap-4 ${!collapsed ? "" : "items-center"} `}
        >
          {items.map((item) => (
            <SidebarItem key={item.to} collapsed={collapsed} {...item} />
          ))}
        </nav>
      </aside>
    </>
  );
};

export default Sidebar;
