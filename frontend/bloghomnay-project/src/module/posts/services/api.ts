
import { axiosInstance } from "../../auth/services/api"
import axios, { AxiosResponse } from "axios"
import { CreateCommentPayload, CreatePostType, UpdatePostType } from "../models/post"
export const ApiGetAllPost = async <T>(): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post("/v1/post/get_post_all")
    return response.data
}
export const ApiGetPostByCategories = async <T>(id: string): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post(`/v1/post/get_post_categories_id/${id}`)
    return response.data
}

export const ApiGetPostById = async <T>(id: string): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post(`/v1/post/get_post_id/${id}`)
    return response.data
}

export const ApiGetPostByUser = async <T>(): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post(`/v2/post/get_post_user_id`)
    return response.data
}

export const ApiSearchPost = async <T>(str: string): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post(`/v1/post/get_post_tag?search=${str}`)
    return response.data
}

export const ApiLike = async <T>(str: string): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post(`/v2/postlikes/create_postlike/${str}`)
    return response.data
}
export const ApiDisLike = async <T>(str: string): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post(`/v2/postlikes/delete_postlike/${str}`)
    return response.data
}


export const ApiCreateComment = async <T>(data: CreateCommentPayload): Promise<T> => {
    const response = await axiosInstance.post<T>('/v2/comment/create_comment', data);
    return response.data;
};

export const ApiCreatePost = async <T>(data: CreatePostType): Promise<T> => {
    const response = await axiosInstance.post<T>('/v2/post/create_post', data)
    return response.data
}


export const ApiUpdatePost = async <T>(data: { data: UpdatePostType, id: string }): Promise<T> => {
    const response = await axiosInstance.post<T>(`/v2/post/update_post_id/${data.id}`, data.data)
    console.log(data.data)
    return response.data
}
export const ApiGetImg = async <T>(): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post(`/v2/img/get_img`)
    return response.data
}
export const ApiGoCreateImg = async <T>(data: { img: string }): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post(
        `/v2/img/create_img`,
        data
    );
    return response.data;
};


export const ApiNodeCreateImg = async <T>(file: File): Promise<T> => {
    const formData = new FormData();
    formData.append("image", file);

    const response: AxiosResponse<T> = await axios.post(
        `http://image.bloghomnay.com/upload-image`,
        formData,
        {
            headers: {
                "Content-Type": "multipart/form-data"
            }
        }
    );
    return response.data;
};
export const ApiGetAllCategories = async <T>(): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post("/v1/categories/get_categories")
    return response.data
}
export const ApiGetAllTags = async <T>(): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post("/v1/tag/get_all_tag")
    return response.data
}