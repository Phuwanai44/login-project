export interface Product {
  id: string
  name: string
  price: number
  stock: number
}

export interface ProductResponse {
  data: Product[]
  total: number
  totalPages: number
  page: number
}

export interface ProductPayload {
  name: string
  price: number
  stock: number
}