import api from "./api"

export const loginUser = async (email: string, password: string) => {

  const res = await api.post("/login", {
    email,
    password
  })

  return res.data
}


export const registerUser = async (data: {
  name: string
  email: string
  password: string
}) => {

  const res = await api.post("/register", data)

  return res.data
}


export const getProfile = async () => {

  const res = await api.get("/profile")

  return res.data

}