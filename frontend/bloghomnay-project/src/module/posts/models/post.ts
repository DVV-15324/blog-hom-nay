export type CreatePostType = {
    description: string;
    categories_id: string;
    title: string;
    content: string;
    tags: TagBase[];
}

export type UpdatePostType = {
    categories_id: string;
    title: string;
    description: string;
    content: string;
    tags: TagBase[];
}

export type TagBase = {
    [x: string]: any;
    id: string;
    name: string;
};

export type CommentBase = {
    id: string;
    post_id: string;
    user_id: string;
    content: string;
    user: UserBase;
    created_at: string;
};

export type PostResponse = {
    id: string;
    user_id: string;
    category_id: string;
    title: string;
    content: string;
    description: string;
    islike: boolean;
    like: number;
    comments: CommentBase[];
    tags: TagBase[];
    user: UserBase;
    count_comment: number;
    updated_at: string;
    created_at: string;
};

export type UserBase = {
    user_id: string;
    first_name: string;
    last_name: string;
    avatar: {
        String: string;
        Valid: boolean;
    };
    updated_at: string;
    created_at: string;
    deleted_at: string;
};

export interface CreateCommentPayload {
    post_id: string;
    content: string;
}

export interface ImgReponse {
    img: string;
}


