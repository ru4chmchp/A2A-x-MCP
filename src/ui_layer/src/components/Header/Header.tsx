interface HeaderProps {
  lable: string;
  content: string;
}

const Header: React.FC<HeaderProps> = ({ lable, content }) => {
  return (
    <>
      <div className="flex flex-col items-start mb-5">
        <h1 className="text-2xl font-bold">{lable}</h1>
        <p className="text-gray-500">{content}</p>
      </div>
    </>
  );
};

export default Header;
