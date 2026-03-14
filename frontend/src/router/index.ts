import { createRouter, createWebHistory } from "vue-router"

import LoginView from "@/views/page/LoginView.vue"
import RegisterView from "@/views/page/RegisterView.vue"

import DashboardLayout from "@/layouts/DashboardLayout.vue"
import DashboardView from "@/views/page/DashboardView.vue"
import UserDashboard from "@/views/page/DashboardUser.vue"
import ProfileDashboard from "@/views/page/DadhboardProfile.vue"

const routes = [

  {
    path: "/login",
    component: LoginView
  },

  {
    path: "/register",
    component: RegisterView
  },

  {
    path: "/dashboard",
    component: DashboardLayout,
    children: [
      {
        path: "",
        component: DashboardView
      }
    ]
  },

  {
    path: "/users",
    component: DashboardLayout,
    children: [
      {
        path: "",
        component: UserDashboard
      }
    ]
  },

  {
    path: "/profile",
    component: DashboardLayout,
    children: [
      {
        path: "",
        component: ProfileDashboard
      }
    ]
  }

]

const router = createRouter({
  history: createWebHistory(),
  routes
})


// Route Guard (กันเข้า dashboard ถ้าไม่ login)
router.beforeEach((to, from, next) => {

  const token = localStorage.getItem("token")

  if (!token && to.path !== "/login" && to.path !== "/register") {
    next("/login")
  } else {
    next()
  }

})

export default router