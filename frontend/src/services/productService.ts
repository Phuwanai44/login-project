import axios from "axios"
import type { ProductResponse } from "@/types/product"

export const getProducts = async (
    page: number,
    search: string
): Promise<ProductResponse> => {

    const res = await axios.get<ProductResponse>(
        "http://localhost:3000/products",
        {
            params: {
                page,
                limit: 10,
                search
            }
        }
    )

    return res.data
}