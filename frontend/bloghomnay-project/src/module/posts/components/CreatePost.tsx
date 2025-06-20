//Lấy mẫu từ https://youtu.be/LE5PoCg35GI

import { Plus } from "lucide-react";
import {
    useState,
    ChangeEvent,
    FormEvent,
    useRef,
    useEffect,
} from "react";
import ReactQuill from "react-quill-new";
import "quill/dist/quill.snow.css";

import {
    CreatePostType,
    ImgReponse,
    TagBase,
} from "../models/post";
import parse from "html-react-parser";
import {
    ApiCreatePost,
    ApiGetAllCategories,
    ApiGetAllTags,
    ApiGetImg,
    ApiGoCreateImg,
    ApiNodeCreateImg,
} from "../services/api";
import { enqueueSnackbar } from "notistack";
import { AxiosError } from "axios";
import { Categories } from "../../categories/module/categories";
import { Response } from "../../auth/model/auth"


export default function CreatePost() {
    const [uploadedImages, setImage] = useState<ImgReponse[]>([]);
    const [title, setTitle] = useState("");
    const [selectedCategoryId, setSelectedCategoryId] = useState("");
    const [selectedTags, setSelectedTags] = useState<TagBase[]>([]);
    const [categories, setCategories] = useState<Categories[]>([]);
    const [tags, setTags] = useState<TagBase[]>([]);
    const [description, setDescription] = useState("");
    const [content, setContent] = useState("");
    const quillRef = useRef<ReactQuill | null>(null);

    const insertImageToEditor = (url: string) => {
        const editor = quillRef.current?.getEditor();
        const range = editor?.getSelection();
        if (range) {
            editor!.insertEmbed(range.index, "image", url);
            editor!.setSelection(range.index + 1);
        }
    };

    useEffect(() => {
        const fetchData = async () => {
            try {
                const [imgRes, catRes, tagRes] = await Promise.all([
                    ApiGetImg<Response<ImgReponse[]>>(),
                    ApiGetAllCategories<Response<Categories[]>>(),
                    ApiGetAllTags<Response<TagBase[]>>(),
                ]);
                setImage(imgRes.data);
                setCategories(catRes.data);
                setTags(tagRes.data);
                console.log(imgRes, catRes, tagRes)
            } catch (error) {
                enqueueSnackbar("Lỗi tải dữ liệu", { variant: "error" });
            }
        };
        fetchData();
    }, []);

    const handleUploadImg = async (event: ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        if (!file) return;

        try {
            const nodeRes = await ApiNodeCreateImg<{ img: string }>(file);
            const imageUrl = nodeRes.img;
            console.log(imageUrl)
            await ApiGoCreateImg({ img: imageUrl });
            setImage((prev) => [...prev, { img: imageUrl }]);
            enqueueSnackbar("Tải ảnh thành công!", { variant: "success" });
        } catch (error) {
            const err = error as AxiosError;
            enqueueSnackbar(err.message || "Lỗi tải ảnh", { variant: "error" });
        }
    };

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const newBlog: CreatePostType = {
            title,
            categories_id: selectedCategoryId,
            description,
            content,
            tags: selectedTags,

        };
        try {
            console.log(newBlog)
            await ApiCreatePost<CreatePostType>(newBlog);

            enqueueSnackbar("Tạo Post thành công!", { variant: "success" });
        } catch (error) {
            const err = error as AxiosError;
            enqueueSnackbar(err.message || "Lỗi tải ảnh", { variant: "error" });
        }
    };

    const modules = {
        toolbar: [
            [{ 'header': [1, 2, 3, false] }],
            ["bold", "italic", "underline", "strike", "blockquote"],
            [{ align: "" }, { align: "center" }, { align: "right" }, { align: "justify" }],
            [{ list: "ordered" }, { list: "bullet" }],
            ["link", "color"],
            [{ "code-block": true }],
            ["clean"],

        ],
        resizeImage: {
            displaySize: true,
            modules: ['Resize', 'DisplaySize'],
        }
    };
    const formats = [
        "header",
        "bold",
        "italic",
        "underline",
        "strike",
        "blockquote",
        "list",
        "link",
        "indent",
        "image",
        "code-block",
        "color",
        "align",
    ];

    return (
        <div className="flex h-full w-full flex-col items-center justify-center">
            <div className="sm:col-span-2">
                <label className="block mb-2 text-sm font-medium text-gray-900">
                    Image Gallery
                </label>
                <div className="flex gap-4 overflow-x-auto border">
                    {Array.isArray(uploadedImages) &&
                        uploadedImages.map((img, index) => (
                            <img
                                key={index}
                                src={img.img}
                                alt={`uploaded-${index}`}
                                className="w-24 h-24 object-contain aspect-square cursor-pointer border rounded hover:opacity-80"
                                onClick={() => insertImageToEditor(img.img)}
                            />
                        ))}
                </div>
                <p className="text-xs text-gray-500 mt-1">
                    Click an image to insert into content
                </p>
            </div>
            {/* Upload ảnh mới */}
            <div className="sm:col-span-2 mt-4">
                <label className="block mb-2 text-sm font-medium text-gray-900">
                    Upload Image
                </label>
                <input
                    type="file"
                    accept="image/*"
                    onChange={handleUploadImg}
                    className="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 focus:outline-none"
                />
                <p className="mt-1 text-xs text-gray-500">
                    Chọn ảnh từ máy của bạn để tải lên thư viện.
                </p>
            </div>

            <div className="sm:col-span-2">
                <label className="block mb-2 text-sm font-medium text-gray-900">
                    Category
                </label>
                <select
                    value={selectedCategoryId}
                    onChange={(e) => setSelectedCategoryId(e.target.value)}
                    className="block w-full rounded-md border-0 py-2 px-4 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:ring-2 focus:ring-purple-600 sm:text-sm"
                >
                    <option value="" disabled>
                        Chọn danh mục
                    </option>
                    {categories.map((cat) => (
                        <option key={cat.id} value={cat.id}>
                            {cat.name}
                        </option>
                    ))}
                </select>
            </div>

            <div className="sm:col-span-2">
                <label className="block mb-2 text-sm font-medium text-gray-900">Tags</label>
                <div className="flex flex-wrap gap-3">
                    {tags.map((tag) => {
                        const isChecked = selectedTags.some((t) => t.id === tag.id);

                        return (
                            <label key={tag.id} className="flex items-center space-x-2 cursor-pointer border px-3 py-1 rounded-md hover:bg-gray-100">
                                <input
                                    type="checkbox"
                                    checked={isChecked}
                                    onChange={(e) => {
                                        if (e.target.checked) {
                                            setSelectedTags((prev) => [...prev, tag]);
                                        } else {
                                            setSelectedTags((prev) => prev.filter((t) => t.id !== tag.id));
                                        }
                                    }}
                                    className="accent-purple-600"
                                />
                                <span className="text-sm">{tag.name}</span>
                            </label>
                        );
                    })}
                </div>
            </div>


            <div className="grid grid-cols-1 lg:grid-cols-2 p-8 gap-4">
                <div className="w-full max-w-3xl p-5 my-6 bg-white border border-gray-200 rounded-lg shadow mx-auto">
                    <h2 className="text-3xl font-bold border-b border-gray-400 pb-2 mb-5">
                        Chỉnh Sửa BLog
                    </h2>
                    <form onSubmit={handleSubmit} onKeyDown={(e) => {
                        if (e.key === "Enter" && e.target instanceof HTMLInputElement) {
                            e.preventDefault();
                        }
                    }}>
                        <div className="grid gap-4 sm:grid-cols-2 sm:gap-6">
                            <div className="sm:col-span-2">
                                <label htmlFor="title" className="block text-sm font-medium text-gray-900 mb-2">
                                    Tiêu Đề Blog
                                </label>
                                <input
                                    onChange={(e) => setTitle(e.target.value)}
                                    type="text"
                                    value={title}
                                    id="title"
                                    className="block w-full rounded-md border-0 py-2 px-4 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-purple-600 sm:text-sm"
                                    placeholder="Nhập tiêu đề blog"
                                />
                            </div>

                            <div className="sm:col-span-2">
                                <label htmlFor="description" className="block mb-2 text-sm font-medium text-gray-900">
                                    Miêu Tả Blog
                                </label>
                                <textarea
                                    id="description"
                                    rows={4}
                                    onChange={(e) => setDescription(e.target.value)}
                                    value={description}
                                    className="block w-full p-2.5 text-sm text-gray-900 bg-white rounded-lg border border-gray-300 focus:ring-purple-500 focus:border-purple-500"
                                    placeholder="Viết một mô tả ngắn gọn..."
                                />
                            </div>

                            <div className="sm:col-span-2">
                                <label htmlFor="content" className="block mb-2 text-sm font-medium text-gray-900">
                                    Nội Dung Blog
                                </label>
                                <ReactQuill
                                    ref={quillRef}
                                    theme="snow"
                                    value={content}
                                    onChange={setContent}
                                    modules={modules}
                                    formats={formats}
                                />
                            </div>
                        </div>

                        <button
                            type="submit"
                            className="inline-flex items-center px-5 py-2.5 mt-4 sm:mt-6 text-sm font-medium text-white bg-purple-700 rounded-lg hover:bg-purple-800 focus:ring-4 focus:ring-purple-200"
                        >
                            <Plus className="w-5 h-5 mr-2" />
                            <span>Tạo Blog</span>
                        </button>
                    </form>
                </div>

                <div className="w-full max-w-3xl p-8 my-6 bg-white border border-gray-200 rounded-lg shadow mx-auto">
                    <h2 className="text-3xl font-bold border-b border-gray-400 pb-2 mb-5">Xem Trước Blog</h2>
                    <div className="grid gap-4 sm:grid-cols-2 sm:gap-6">
                        <div className="sm:col-span-2">
                            <h2 className="text-sm font-medium text-gray-900 mb-2">Tiêu Đề Blog</h2>
                            <p className="text-2xl font-bold">{title}</p>
                        </div>
                        <div className="sm:col-span-2">
                            <h2 className="text-sm font-medium text-gray-900 mb-2">Miêu Tả Blog</h2>
                            <p>{description}</p>
                        </div>
                        <div className="sm:col-span-full">
                            <h2 className="text-sm font-medium text-gray-900 mb-2">Nội Dung Blog</h2>
                            <div className="ql-editor">{parse(content)}</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}