import { FC, useState } from "react";
import parse, { HTMLReactParserOptions, Element } from "html-react-parser";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { oneDark } from "react-syntax-highlighter/dist/esm/styles/prism";

interface Props {
    content: string;
}

const PreviewWithCodeBlock: FC<Props> = ({ content }) => {
    const options: HTMLReactParserOptions = {
        replace: (domNode) => {
            if (
                domNode instanceof Element &&
                domNode.name === "div" &&
                domNode.attribs?.class?.includes("ql-code-block-container")
            ) {
                // Lấy tất cả các thẻ con ql-code-block
                const codeBlocks = domNode.children.filter(
                    (child) =>
                        child instanceof Element && child.attribs?.class?.includes("ql-code-block")
                ) as Element[];

                // Lấy text từng dòng
                const lines = codeBlocks.map((block) => {
                    // Nếu có thẻ <br> hoặc rỗng thì trả về dòng trống
                    if (
                        block.children.length === 1 &&
                        block.children[0].type === "tag" &&
                        (block.children[0] as Element).name === "br"
                    ) {
                        return "";
                    }

                    // Lấy text bên trong block
                    return block.children
                        .map((child) => {
                            if (child.type === "text") return child.data || "";
                            return "";
                        })
                        .join("");
                });

                const codeContent = lines.join("\n");

                return <CodeBlockWithCopy code={codeContent} />;
            }
        },
    };

    return (
        <div className="prose max-w-none ql-editor" style={{ whiteSpace: "pre-wrap" }}>
            {parse(content, options)}
        </div>
    );
};

interface CodeBlockProps {
    code: string;
}

const CodeBlockWithCopy: FC<CodeBlockProps> = ({ code }) => {
    const [copied, setCopied] = useState(false);

    const handleCopy = () => {
        navigator.clipboard.writeText(code).then(() => {
            setCopied(true);
            setTimeout(() => setCopied(false), 1500);
        });
    };

    return (
        <div className="relative group">
            <button
                onClick={handleCopy}
                className="absolute top-1 right-1  group-hover:opacity-100 bg-gray-700 text-white text-xs px-5 py-3 rounded"
                aria-label="Copy code"
                type="button"
            >
                {copied ? "Copied!" : "Copy"}
            </button>
            <SyntaxHighlighter
                language="text"
                style={oneDark}

            >
                {code}
            </SyntaxHighlighter>



        </div>
    );
};

export default PreviewWithCodeBlock;
