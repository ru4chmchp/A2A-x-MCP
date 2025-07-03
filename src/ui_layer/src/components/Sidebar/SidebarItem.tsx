import type { ReactNode } from "react";
import { NavLink } from "react-router";

// khai báo các props mà mình sẽ truyền vào SidebarItem
interface SidebarItemProps {
  icon: ReactNode; // kiểu dữ liệu trả về, có thể mô tả bất cứ thứ gì được render bởi react
  lable: string;
  to: string;
  collapsed: boolean; // trạng thái thu gọn sidebar hay không
}

const SidebarItem: React.FC<SidebarItemProps> = ({
  icon,
  lable,
  to,
  collapsed,
}) => {
  return (
    <>
      <NavLink
        to={to}
        className={({ isActive }) =>
          `flex items-center gap-3 p-3 rounded-md transition-transform duration-300 ease-in-out hover:scale-120 ${
            isActive ? "bg-gray-800 text-white" : "text-gray-400"
          }`
        }
      >
        <span>{icon}</span>
        {!collapsed && <span className="text-sm">{lable}</span>}
      </NavLink>
    </>
  );
};

export default SidebarItem;
