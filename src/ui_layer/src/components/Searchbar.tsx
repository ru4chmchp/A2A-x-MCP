import { Search } from "lucide-react";

interface SearchbarProps {
  query: string;
  onQueryChange: (value: string) => void;
  // khúc này khi cha gọi Searchbar thì truyển vào một hàm
}

const Searchbar: React.FC<SearchbarProps> = ({ query, onQueryChange }) => {
  return (
    <>
      <div className="relative w-full">
        <Search
          className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
          size={18}
        />
        <input
          type="text"
          placeholder="Enter your web application you want to recon (e.g., example.com)"
          value={query}
          // onChange sẽ gọi hàm setQuery với value được user nhập vào
          onChange={(e) => onQueryChange(e.target.value)}
          className="w-full pl-10 pr-2 py-3 rounded-xl bg-white border border-gray-300 text-sm focus:outline-none focus:ring-2 forucs: ring-indigo-500 shadow-sm"
        />
      </div>
    </>
  );
};

export default Searchbar;
