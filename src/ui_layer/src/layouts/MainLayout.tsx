import type { ReactNode } from "react";
import Sidebar from "../components/Sidebar/Sidebar";
import Header from "../components/Header/Header";

interface MainLayoutProps {
  children?: ReactNode;
}

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => {
  return (
    <>
      <div className="flex bg-red-200 p-5">
        <Sidebar />
        <main className="flex-1 p-6 bg-gray-100 ml-2 rounded-xl">
          <Header />
          {/* nơi mà các page sẽ implement */}
          {children}
        </main>
      </div>
    </>
  );
};

export default MainLayout;
