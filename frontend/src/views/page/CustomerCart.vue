<script setup lang="ts">
import { useCartStore } from '@/stores/cart'
import { useRouter } from 'vue-router'
import NavbarHome from '@/components/NavbarHome.vue'
import FooterHome from '@/components/FooterHome.vue'

const cartStore = useCartStore()
const router = useRouter()

const handleQuantityChange = (item: any, amount: number) => {
  const newQuantity = item.quantity + amount
  if (newQuantity > 0) {
    cartStore.updateQuantity(item.id, newQuantity)
  }
}

const handleRemove = (itemId: string) => {
  cartStore.removeItem(itemId)
}

const handleCheckout = () => {
  router.push('/checkout')
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
      <h2 class="display-5 fw-extrabold text-gradient mb-4">Your Cart</h2>
      
      <div class="row g-4" v-if="cartStore.items.length > 0">
        <!-- Cart Items -->
        <div class="col-lg-8">
          <div class="glass-card p-4 h-100">
            <div 
              v-for="(item, index) in cartStore.items" 
              :key="item.id"
              class="d-flex align-items-center mb-4 pb-4 border-bottom border-secondary cart-item-row"
            >
              <!-- Image -->
              <img :src="item.image" class="rounded cart-img me-4" :alt="item.name">
              
              <!-- Details -->
              <div class="flex-grow-1">
                <h5 class="fw-bold text-white mb-1">{{ item.name }}</h5>
                <p class="text-muted small mb-2">Variant: {{ item.color }} - {{ item.size }}</p>
                <h6 class="text-success fw-bold">฿{{ item.price.toLocaleString() }}</h6>
              </div>
              
              <!-- Quantity & Remove -->
              <div class="d-flex flex-column align-items-end gap-3">
                <div class="input-group input-group-sm" style="width: 100px;">
                  <button class="btn btn-outline-secondary text-white" @click="handleQuantityChange(item, -1)">-</button>
                  <input type="text" class="form-control text-center bg-dark text-white border-secondary" :value="item.quantity" readonly>
                  <button class="btn btn-outline-secondary text-white" @click="handleQuantityChange(item, 1)">+</button>
                </div>
                
                <button class="btn btn-sm btn-link text-danger text-decoration-none" @click="handleRemove(item.id)">
                  <i class="bi bi-trash"></i> Remove
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Summary -->
        <div class="col-lg-4">
          <div class="glass-card p-4 sticky-top" style="top: 100px;">
            <h4 class="fw-bold text-white mb-4">Summary</h4>
            
            <div class="d-flex justify-content-between mb-3 text-muted">
              <span>Items ({{ cartStore.totalItems }})</span>
              <span>฿{{ cartStore.totalPrice.toLocaleString() }}</span>
            </div>
            <div class="d-flex justify-content-between mb-4 text-muted">
              <span>Shipping</span>
              <span>Free</span>
            </div>
            
            <hr class="border-secondary">
            
            <div class="d-flex justify-content-between mb-4 mt-3">
              <span class="fs-5 text-white fw-bold">Total</span>
              <span class="fs-5 text-success fw-bold">฿{{ cartStore.totalPrice.toLocaleString() }}</span>
            </div>
            
            <button class="btn btn-primary w-100 py-3 fw-bold" @click="handleCheckout">
              Proceed to Checkout
            </button>
          </div>
        </div>
      </div>
      
      <!-- Empty Cart -->
      <div v-else class="text-center py-5 glass-card">
        <i class="bi bi-cart-x fs-1 text-muted"></i>
        <h3 class="text-white mt-3">Your cart is empty</h3>
        <p class="text-muted mb-4">Looks like you haven't added anything to your cart yet.</p>
        <router-link to="/store" class="btn btn-primary px-4 py-2">
          Start Shopping
        </router-link>
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

.cart-img {
  width: 80px;
  height: 100px;
  object-fit: cover;
}

.cart-item-row:last-child {
  border-bottom: none !important;
  margin-bottom: 0 !important;
  padding-bottom: 0 !important;
}
</style>
