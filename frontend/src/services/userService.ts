import api from "./api"
import type { User } from "@/types/user"

export const getUsers = async (): Promise<User[]> => {

  const res = await api.get("/admin/users")

  console.log("API RESPONSE:", res.data)

  return res.data

}

export const deleteUserById = async (id: string) => {

  const res = await api.delete(`/admin/users/${id}`)

  return res.data
}