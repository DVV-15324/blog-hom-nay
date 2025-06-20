export type LoginType = {
    email: string;
    password: string;
}
export type ResponseLoginType = {

    access_token: {
        token: string;
        expire_at: string;
    }

}
export type RegisterType = {
    email: string;
    password: string;
    last_name: string;
    first_name: string;
    phone: string;
}

export type ProfileType = {
    id: string;
    email: string;
    last_name: string;
    first_name: string;
    phone: string;
    address: {
        String: string;
        Valid: boolean;
    };
    avatar: {
        String: string;
        Valid: boolean;
    };
}



export type Response<T> = {
    data: T
}