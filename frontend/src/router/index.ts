import { createRouter, createWebHistory } from 'vue-router'

import LoginView from '@/views/page/LoginView.vue'
import RegisterView from '@/views/page/RegisterView.vue'
import HomeView from '@/views/page/HomeView.vue'

import DashboardLayout from '@/layouts/DashboardLayout.vue'
import DashboardView from '@/views/page/DashboardView.vue'
import UserDashboard from '@/views/page/DashboardUser.vue'
import ProfileDashboard from '@/views/page/DadhboardProfile.vue'
import ProductDashboard from '@/views/page/ProductDashboard.vue'
import ProductDashboardAdd from '@/views/page/ProductDashboardAdd.vue'
import EditProduct from '@/views/page/EditProduct.vue'
import CustomerProducts from '@/views/page/CustomerProducts.vue'
import CustomerProductDetail from '@/views/page/CustomerProductDetail.vue'
import CustomerCart from '@/views/page/CustomerCart.vue'
import CustomerCheckout from '@/views/page/CustomerCheckout.vue'
import CustomerOrders from '@/views/page/CustomerOrders.vue'
import CustomerOrderDetail from '@/views/page/CustomerOrderDetail.vue'

const routes = [
  {
    path: '/',
    component: HomeView,
  },

  {
    path: '/login',
    component: LoginView,
  },
  {
    path: '/store',
    component: CustomerProducts,
  },
  {
    path: '/store/:id',
    component: CustomerProductDetail,
  },
  {
    path: '/cart',
    component: CustomerCart,
  },
  {
    path: '/checkout',
    component: CustomerCheckout,
  },
  {
    path: '/orders',
    component: CustomerOrders,
  },
  {
    path: '/orders/:id',
    component: CustomerOrderDetail,
  },

  {
    path: '/register',
    component: RegisterView,
  },
  {
    path: '/dashboard',
    component: DashboardLayout,
    children: [
      {
        path: '',
        component: DashboardView,
      },
    ],
  },

  {
    path: '/users',
    component: DashboardLayout,
    children: [
      {
        path: '',
        component: UserDashboard,
      },
    ],
  },

  {
    path: '/profile',
    component: DashboardLayout,
    children: [
      {
        path: '',
        component: ProfileDashboard,
      },
    ],
  },

  {
    path: '/products',
    component: DashboardLayout,
    children: [
      {
        path: '',
        component: ProductDashboard,
      },
    ],
  },

  {
    path: '/add-products',
    component: DashboardLayout,
    children: [
      {
        path: '',
        component: ProductDashboardAdd,
      },
    ],
  },

  {
    path: '/products/edit/:id',
    component: DashboardLayout,
    children: [
      {
        path: '',
        name: 'EditProduct',
        component: EditProduct,
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

const publicPaths = ['/', '/login', '/register', '/store', '/cart']

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  const isPublic = publicPaths.includes(to.path) || to.path.startsWith('/store/')

  if (!token && !isPublic) {
    next('/login')
  } else {
    next()
  }
})

export default router
