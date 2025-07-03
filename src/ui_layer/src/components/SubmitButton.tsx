interface SubmitButtonProps {
  lable?: string;
  disabled?: boolean;
}

const SubmitButton: React.FC<SubmitButtonProps> = ({
  lable = "Send",
  disabled = false,
}) => {
  return (
    <>
      <button
        type="submit"
        disabled={disabled}
        className={`px-4 py-2 rounded-full text-sm  text-center border shadow-sm transition-transform duration-300 ease-in-out hover:scale-120 ${
          disabled
            ? "bg-indigo-300 text-white cursor-not-allowed"
            : "bg-gray-900 text-white hover:bg-indigo-700"
        }`}
      >
        {lable}
      </button>
    </>
  );
};

export default SubmitButton;
