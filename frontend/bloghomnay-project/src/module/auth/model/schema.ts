import * as yup from "yup";

export const LoginSchema = yup.object({
    email: yup.string().email("email invalid").required("email is required"),
    password: yup.string().required("password is required")
})

export const RegisterSchema = yup.object({
    first_name: yup.string().required("first name is required"),
    last_name: yup.string().required("last name is required"),
    email: yup.string().email("email invalid").required("email is required"),
    password: yup.string().required("password is required"),
    phone: yup.string().required("phone is required"),
}) 