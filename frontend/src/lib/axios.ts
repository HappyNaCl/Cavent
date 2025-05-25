import axios from "axios";
import { env } from "./schema/EnvSchema";

const api = axios.create({
    baseURL: env.VITE_BACKEND_URL + "/api/v1"
})

export default api