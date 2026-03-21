import axios from 'axios'
import type { Product, ProductResponse, ProductPayload } from '@/types/product'

export const getProducts = async (page: number, search: string): Promise<ProductResponse> => {
  const res = await axios.get<ProductResponse>('http://localhost:3000/products', {
    params: {
      page,
      limit: 10,
      search,
    },
  })

  return res.data
}
// ✅ CREATE
export const createProduct = async (product: Product) => {
  const res = await axios.post('http://localhost:3000/products', product)
  return res.data
}

export const deleteProduct = async (id: string) => {
  const res = await axios.delete(`http://localhost:3000/products/${id}`)
  return res.data
}

export const getProductById = async (id: string) => {
  const res = await axios.get(`http://localhost:3000/products/${id}`)
  return res.data
}

export const editProduct = async (id: string, data: ProductPayload) => {
  const res = await axios.put(
    `http://localhost:3000/products/${id}`,
    data
  )
  return res.data
}
