import React from "react";

export interface TOCItem {
    id: string;
    text: string;
    tag: "h1" | "h2" | "h3";

}

interface TableOfContentsProps {
    tocList: TOCItem[];
    activeId: string | null;


}

const TableOfContents: React.FC<TableOfContentsProps> = ({ tocList, activeId }) => {


    return (
        <nav
            className="fixed top-24 mt-9 w-64 p-4 bg-white rounded-md border border-gray-200 shadow overflow-y-auto h-full"
            aria-label="Table of contents"
        >
            <h3 className="text-lg font-semibold mb-2 text-gray-800 text-center">Mục lục</h3>
            <ul className="space-y-1 text-sm">
                {tocList.map(({ id, text, tag }) => {
                    const paddingLeft =
                        tag === "h1" ? "pl-2" : tag === "h2" ? "pl-6" : "pl-10";

                    const activeClass = activeId === id
                        ? "border-l-4 border-red-500 bg-red-50 font-semibold text-red-600"
                        : "border-l-4 border-transparent text-blue-600";

                    return (
                        <li key={id}>
                            <button
                                onClick={() => {

                                    const el = document.getElementById(id);
                                    if (el) {
                                        el.scrollIntoView({ behavior: "smooth" });

                                    }
                                }}

                                className={`block w-full text-left px-2 py-1 rounded hover:underline cursor-pointer ${paddingLeft} ${activeClass}`}
                                aria-current={activeId === id ? "true" : undefined}
                                type="button"
                            >
                                {text}
                            </button>
                        </li>
                    );
                })}
            </ul>
        </nav>
    );
};

export default TableOfContents;
