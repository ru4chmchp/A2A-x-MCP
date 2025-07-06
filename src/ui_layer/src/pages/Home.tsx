import SubmitForm from "../components/SubmitForm";
import MainLayout from "../layouts/MainLayout";

const params = {
  lable: "Welcome back 👋",
  content: "Here's what's happening in your web application",
};
const Home: React.FC = () => {
  return (
    <>
      <MainLayout lable={params.lable} content={params.content}>
        <SubmitForm />
      </MainLayout>
    </>
  );
};

export default Home;
