<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import { useCartStore } from '@/stores/cart'
import NavbarHome from '@/components/NavbarHome.vue'
import FooterHome from '@/components/FooterHome.vue'

const router = useRouter()
const cartStore = useCartStore()

const form = ref({
  fullName: '',
  phone: '',
  address: '',
  city: '',
})
const loading = ref(false)
const error = ref('')

const totalPrice = computed(() => cartStore.totalPrice)

const handleSubmit = async () => {
  if (!form.value.fullName || !form.value.phone || !form.value.address || !form.value.city) {
    error.value = 'Please fill in all fields.'
    return
  }
  if (cartStore.items.length === 0) {
    error.value = 'Your cart is empty.'
    return
  }

  const token = localStorage.getItem('token')
  if (!token) {
    router.push('/login')
    return
  }

  loading.value = true
  error.value = ''

  try {
    const payload = {
      items: cartStore.items.map(item => ({
        productId: item.productId,
        name: item.name,
        color: item.color,
        size: item.size,
        quantity: item.quantity,
        price: item.price,
      })),
      shippingAddress: {
        fullName: form.value.fullName,
        phone: form.value.phone,
        address: form.value.address,
        city: form.value.city,
      },
    }

    await axios.post('http://localhost:8080/orders', payload, {
      headers: { Authorization: `Bearer ${token}` },
    })

    cartStore.clearCart()
    router.push('/orders')
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Failed to place order.'
  } finally {
    loading.value = false
  }
}
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
      <h2 class="display-5 fw-extrabold text-gradient mb-4">Checkout</h2>

      <div class="row g-4">
        <!-- Shipping Form -->
        <div class="col-lg-7">
          <div class="glass-card p-4">
            <h5 class="fw-bold text-white mb-4"><i class="bi bi-truck me-2"></i>Shipping Address</h5>

            <div v-if="error" class="alert alert-danger">{{ error }}</div>

            <div class="mb-3">
              <label class="form-label text-muted">Full Name</label>
              <input v-model="form.fullName" type="text" class="form-control bg-dark text-white border-secondary" placeholder="John Doe">
            </div>
            <div class="mb-3">
              <label class="form-label text-muted">Phone</label>
              <input v-model="form.phone" type="text" class="form-control bg-dark text-white border-secondary" placeholder="08X-XXX-XXXX">
            </div>
            <div class="mb-3">
              <label class="form-label text-muted">Address</label>
              <textarea v-model="form.address" class="form-control bg-dark text-white border-secondary" rows="3" placeholder="123 Main St."></textarea>
            </div>
            <div class="mb-3">
              <label class="form-label text-muted">City</label>
              <input v-model="form.city" type="text" class="form-control bg-dark text-white border-secondary" placeholder="Bangkok">
            </div>
          </div>
        </div>

        <!-- Order Summary -->
        <div class="col-lg-5">
          <div class="glass-card p-4 sticky-top" style="top: 100px;">
            <h5 class="fw-bold text-white mb-4"><i class="bi bi-receipt me-2"></i>Order Summary</h5>

            <div v-for="item in cartStore.items" :key="item.id" class="d-flex justify-content-between mb-3 text-muted small">
              <span>{{ item.name }} ({{ item.color }}/{{ item.size }}) x{{ item.quantity }}</span>
              <span class="text-white">฿{{ (item.price * item.quantity).toLocaleString() }}</span>
            </div>

            <hr class="border-secondary">

            <div class="d-flex justify-content-between mb-4">
              <span class="fw-bold text-white">Total</span>
              <span class="fw-bold text-success fs-5">฿{{ totalPrice.toLocaleString() }}</span>
            </div>

            <button
              class="btn btn-primary w-100 py-3 fw-bold"
              @click="handleSubmit"
              :disabled="loading || cartStore.items.length === 0"
            >
              <span v-if="loading" class="spinner-border spinner-border-sm me-2"></span>
              Place Order
            </button>

            <router-link to="/cart" class="btn btn-link text-muted w-100 mt-2">← Back to Cart</router-link>
          </div>
        </div>
      </div>
    </main>

    <FooterHome />
  </div>
</template>

<style scoped>
.page-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  position: relative;
}
.form-control:focus {
  border-color: #8b5cf6;
  box-shadow: 0 0 0 0.25rem rgba(139, 92, 246, 0.25);
}
</style>
