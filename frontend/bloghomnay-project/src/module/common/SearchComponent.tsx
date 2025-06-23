import { Search } from 'lucide-react';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const ISearchComponent = () => {
    const navigate = useNavigate();
    const [query, setQuery] = useState<string>("")

    const encodedString = encodeURIComponent(query);

    const handleSearch = () => {
        navigate(`/search/${encodedString}`);
    };

    return (
        <div className="p-4 max-w-4xl mx-auto">
            <div className="flex">
                <input
                    type="text"
                    value={query}
                    onChange={e => setQuery(e.target.value)}
                    onKeyDown={e => {
                        if (e.key === "Enter") handleSearch();
                    }}
                    placeholder="e.g. [golang][docker] backend"
                    className="w-full p-2 border rounded"
                />

                <button
                    onClick={handleSearch}
                    disabled={!query.trim()}
                    className={`px-2 text-black rounded cursor-pointer transition hover:shadow-xl hover:scale-[1.02] duration-300
        ${!query.trim() ? "opacity-50 cursor-not-allowed" : ""}
    `}
                >
                    <Search />
                </button>

            </div>
        </div>
    );
};
export default ISearchComponent