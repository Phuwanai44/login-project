<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'
import NavbarHome from '@/components/NavbarHome.vue'
import FooterHome from '@/components/FooterHome.vue'

const route = useRoute()
const router = useRouter()

interface OrderItem {
  productId: string
  name: string
  color: string
  size: string
  quantity: number
  price: number
}

interface Order {
  id: string
  totalPrice: number
  status: string
  createdAt: string
  items: OrderItem[]
  shippingAddress: {
    fullName: string
    phone: string
    address: string
    city: string
  }
}

const order = ref<Order | null>(null)
const loading = ref(true)

const statusClass = (status: string) => {
  const map: Record<string, string> = {
    pending: 'warning',
    shipped: 'info',
    delivered: 'success',
    cancelled: 'danger',
  }
  return `badge bg-${map[status] || 'secondary'} fs-6`
}

const fetchOrder = async () => {
  const token = localStorage.getItem('token')
  if (!token) { router.push('/login'); return }

  try {
    const res = await axios.get(`http://localhost:8080/orders/${route.params.id}`, {
      headers: { Authorization: `Bearer ${token}` },
    })
    order.value = res.data.data
  } catch {
    router.push('/orders')
  } finally {
    loading.value = false
  }
}

onMounted(fetchOrder)
</script>

<template>
  <div class="page-wrapper">
    <div class="glow-orb-container">
      <div class="glow-orb orb-purple"></div>
      <div class="glow-orb orb-blue"></div>
      <div class="glow-orb orb-cyan"></div>
    </div>

    <NavbarHome />

    <main class="container mt-5 flex-grow-1">
      <div v-if="loading" class="text-center py-5">
        <div class="spinner-border text-primary"></div>
      </div>

      <div v-else-if="order">
        <div class="d-flex justify-content-between align-items-start flex-wrap gap-3 mb-4">
          <div>
            <h2 class="display-6 fw-extrabold text-gradient mb-1">Order Detail</h2>
            <p class="text-muted small">ID: {{ order.id }}</p>
          </div>
          <span :class="statusClass(order.status)">{{ order.status }}</span>
        </div>

        <div class="row g-4 mb-5">
          <!-- Items -->
          <div class="col-lg-7">
            <div class="glass-card p-4">
              <h5 class="fw-bold text-white mb-3"><i class="bi bi-bag me-2"></i>Items</h5>
              <div v-for="item in order.items" :key="item.productId + item.color + item.size"
                class="d-flex justify-content-between align-items-center mb-3 pb-3 border-bottom border-secondary">
                <div>
                  <p class="fw-bold text-white mb-0">{{ item.name }}</p>
                  <p class="text-muted small mb-0">{{ item.color }} / {{ item.size }} × {{ item.quantity }}</p>
                </div>
                <span class="text-success fw-bold">฿{{ (item.price * item.quantity).toLocaleString() }}</span>
              </div>
              <div class="d-flex justify-content-between mt-3">
                <span class="fw-bold text-white">Total</span>
                <span class="fw-bold text-success fs-5">฿{{ order.totalPrice.toLocaleString() }}</span>
              </div>
            </div>
          </div>

          <!-- Shipping & Info -->
          <div class="col-lg-5">
            <div class="glass-card p-4">
              <h5 class="fw-bold text-white mb-3"><i class="bi bi-truck me-2"></i>Shipping Address</h5>
              <p class="text-white mb-1">{{ order.shippingAddress.fullName }}</p>
              <p class="text-muted mb-1">{{ order.shippingAddress.phone }}</p>
              <p class="text-muted mb-1">{{ order.shippingAddress.address }}</p>
              <p class="text-muted mb-0">{{ order.shippingAddress.city }}</p>
            </div>
            <div class="glass-card p-4 mt-3">
              <h5 class="fw-bold text-white mb-3"><i class="bi bi-calendar3 me-2"></i>Order Date</h5>
              <p class="text-muted">{{ new Date(order.createdAt).toLocaleString('th-TH') }}</p>
            </div>
          </div>
        </div>

        <router-link to="/orders" class="btn btn-outline-light">← Back to Orders</router-link>
      </div>
    </main>

    <FooterHome />
  </div>
</template>

<style scoped>
.page-wrapper { min-height: 100vh; display: flex; flex-direction: column; position: relative; }
</style>
