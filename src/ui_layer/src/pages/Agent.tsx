import MainLayout from "../layouts/MainLayout";

const params = {
  lable: "Agent management",
  content: "This page is designed for manage agents",
};

const Agent: React.FC = () => {
  return (
    <>
      <MainLayout lable={params.lable} content={params.content} />
    </>
  );
};

export default Agent;
