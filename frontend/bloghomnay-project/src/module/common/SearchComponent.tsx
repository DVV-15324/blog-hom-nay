import { Search } from 'lucide-react';
import { useState, useEffect } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';

const ISearchComponent = () => {
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();

    // Lấy query từ URL, nếu không có thì rỗng
    const urlQuery = searchParams.get('q') || '';

    // State query nội bộ, khởi tạo từ urlQuery
    const [query, setQuery] = useState<string>(urlQuery);

    // Đồng bộ state query khi URL thay đổi (vd: khi bấm tag thay đổi URL)
    useEffect(() => {
        setQuery(urlQuery);
    }, [urlQuery]);

    const handleSearch = () => {
        if (!query.trim()) return;
        navigate(`/search?q=${encodeURIComponent(query)}`, { replace: true });
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

export default ISearchComponent;
