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
        <div className="p-4 max-w-3xl mx-auto">
            <div className="flex gap-2">
                <input
                    type="text"
                    value={query}
                    onChange={e => setQuery(e.target.value)}
                    placeholder="e.g. [golang][docker] backend"
                    className="w-full p-2 border rounded"
                />
                <button
                    onClick={handleSearch}
                    className="px-3 text-black rounded cursor-pointer transition hover:shadow-xl hover:scale-[1.02] duration-300"
                >
                    <Search />
                </button>
            </div>
        </div>
    );
};
export default ISearchComponent