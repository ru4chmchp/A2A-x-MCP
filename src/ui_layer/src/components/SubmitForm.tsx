import { useState } from "react";
import Searchbar from "./Searchbar";
import SubmitButton from "./SubmitButton";
import TaskTypeSelector from "./TaskTypeSelector";

const SubmitForm: React.FC = () => {
  const [query, setQuery] = useState("");
  const [taskTypes, setTaskTypes] = useState<string[]>([]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log("Gửi form với:", { query, taskTypes });
    // TODO: gọi API hoặc trigger agent ở đây
  };
  return (
    <>
      <form onSubmit={handleSubmit} className="flex flex-col gap-5">
        <div>
          <Searchbar query={query} onQueryChange={setQuery} />
        </div>
        <div>
          <TaskTypeSelector value={taskTypes} onChange={setTaskTypes} />
        </div>
        <div className="text-end">
          <SubmitButton lable="run RECON" />
        </div>
      </form>
    </>
  );
};

export default SubmitForm;
