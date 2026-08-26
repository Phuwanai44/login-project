<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import NavbarHome from '@/components/NavbarHome.vue'
import FooterHome from '@/components/FooterHome.vue'

const router = useRouter()

interface Order {
  id: string
  totalPrice: number
  status: string
  createdAt: string
  items: { name: string; quantity: number }[]
}

const orders = ref<Order[]>([])
const loading = ref(true)

const statusClass = (status: string) => {
  const map: Record<string, string> = {
    pending: 'warning',
    shipped: 'info',
    delivered: 'success',
    cancelled: 'danger',
  }
  return `badge bg-${map[status] || 'secondary'}`
}

const fetchOrders = async () => {
  const token = localStorage.getItem('token')
  if (!token) { router.push('/login'); return }

  try {
    const res = await axios.get('http://localhost:8080/orders', {
      headers: { Authorization: `Bearer ${token}` },
    })
    orders.value = res.data.data
  } catch {
    orders.value = []
  } finally {
    loading.value = false
  }
}

onMounted(fetchOrders)
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
      <h2 class="display-5 fw-extrabold text-gradient mb-4">My Orders</h2>

      <div v-if="loading" class="text-center py-5">
        <div class="spinner-border text-primary"></div>
      </div>

      <div v-else-if="orders.length === 0" class="text-center py-5 glass-card">
        <i class="bi bi-box-seam fs-1 text-muted"></i>
        <h4 class="text-white mt-3">No orders yet</h4>
        <router-link to="/store" class="btn btn-primary mt-3">Start Shopping</router-link>
      </div>

      <div v-else class="d-flex flex-column gap-3 mb-5">
        <div
          v-for="order in orders"
          :key="order.id"
          class="glass-card p-4 d-flex justify-content-between align-items-center flex-wrap gap-3"
        >
          <div>
            <p class="text-muted small mb-1">Order ID: <span class="text-white">{{ order.id }}</span></p>
            <p class="text-muted small mb-1">{{ order.items.map(i => i.name).join(', ') }}</p>
            <p class="text-muted small mb-0">{{ new Date(order.createdAt).toLocaleDateString('th-TH') }}</p>
          </div>
          <div class="text-end">
            <p class="text-success fw-bold fs-5 mb-2">฿{{ order.totalPrice.toLocaleString() }}</p>
            <span :class="statusClass(order.status)">{{ order.status }}</span>
          </div>
          <router-link :to="`/orders/${order.id}`" class="btn btn-outline-light btn-sm">
            View Details
          </router-link>
        </div>
      </div>
    </main>

    <FooterHome />
  </div>
</template>

<style scoped>
.page-wrapper { min-height: 100vh; display: flex; flex-direction: column; position: relative; }
</style>
