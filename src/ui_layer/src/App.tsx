import { BrowserRouter, Route, Routes } from "react-router";
import "./index.css";
import Home from "./pages/Home";
import Agent from "./pages/Agent";
//
const App: React.FC = () => {
  return (
    <>
      {/*  */}
      <BrowserRouter>
        {/*  */}
        <Routes>
          {/*  */}
          <Route path="/" element={<Home />}></Route>
          <Route path="/agent" element={<Agent />}></Route>
        </Routes>
      </BrowserRouter>
    </>
  );
};

export default App;
