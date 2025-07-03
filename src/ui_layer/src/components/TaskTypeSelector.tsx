interface TaskTypeSelectorProps {
  value: string[];
  onChange: (value: string[]) => void;
}

const TASK_TYPES = [
  "IP",
  "Domain",
  "Technologies",
  "Port",
  "Subdomain Enumeration",
  "File bruteforcing",
  "Parameter bruteforcing",
  "Wayback machine",
  "Crawl",
];

const TaskTypeSelector: React.FC<TaskTypeSelectorProps> = ({
  value,
  onChange,
}) => {
  const handleToggle = (type: string) => {
    if (value.includes(type)) {
      onChange(value.filter((t) => t !== type));
    } else {
      onChange([...value, type]);
    }
  };
  return (
    <>
      <div className="flex  flex-wrap gap-2">
        {TASK_TYPES.map((type) => (
          <button
            key={type}
            type="button"
            onClick={() => handleToggle(type)}
            className={`px-4 py-2 rounded-full text-sm border shadown-ws  ${
              value.includes(type)
                ? "bg-indigo-600 text-white border-indigo-600 border ring"
                : "bg-white text-gray-700 border-gray-300 hover:bg-gray-200"
            } `}
          >
            {type}
          </button>
        ))}
      </div>
    </>
  );
};

export default TaskTypeSelector;
