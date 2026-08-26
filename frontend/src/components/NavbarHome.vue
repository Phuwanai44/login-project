<script setup lang="ts">
import { useCartStore } from '@/stores/cart'

const cartStore = useCartStore()
</script>

<template>
  <nav class="navbar glass-nav px-4 d-flex justify-content-between align-items-center">
    <!-- Left -->
    <div class="d-flex align-items-center gap-2">
      <div class="logo-orb"></div>
      <h5 class="mb-0 brand-text">AetherAuth</h5>
    </div>

    <!-- Right -->
    <div class="d-flex gap-2 align-items-center">
      <router-link to="/store" class="btn glass-nav-btn glass-nav-btn-outline me-2">
        <i class="bi bi-shop me-1"></i> Store
      </router-link>

      <router-link to="/cart" class="btn glass-nav-btn glass-nav-btn-outline position-relative me-3">
        <i class="bi bi-cart3 fs-5"></i>
        <span v-if="cartStore.totalItems > 0" class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger">
          {{ cartStore.totalItems }}
        </span>
      </router-link>

      <router-link to="/login" class="btn glass-nav-btn glass-nav-btn-solid">
        Login
      </router-link>

      <router-link to="/register" class="btn glass-nav-btn glass-nav-btn-outline">
        Register
      </router-link>
    </div>
  </nav>
</template>

<style scoped>
.navbar {
  height: 70px;
  background: rgba(13, 10, 30, 0.5) !important;
  backdrop-filter: blur(12px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  position: sticky;
  top: 0;
  z-index: 1000;
}

.logo-orb {
  width: 12px;
  height: 12px;
  background: var(--accent-gradient);
  border-radius: 50%;
  box-shadow: 0 0 10px var(--accent);
}

.brand-text {
  font-weight: 700;
  letter-spacing: 0.5px;
  background: linear-gradient(135deg, #fff 0%, #cbd5e1 50%, #a78bfa 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.glass-nav-btn {
  padding: 6px 16px;
  border-radius: 8px;
  font-size: 0.9rem;
  font-weight: 600;
  transition: all 0.3s ease;
}

.glass-nav-btn-solid {
  background: var(--accent-gradient);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: #fff;
  box-shadow: 0 4px 10px rgba(139, 92, 246, 0.2);
}

.glass-nav-btn-solid:hover {
  transform: translateY(-1px);
  box-shadow: 0 6px 15px rgba(139, 92, 246, 0.4);
}

.glass-nav-btn-outline {
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: var(--text-secondary);
}

.glass-nav-btn-outline:hover {
  background: rgba(255, 255, 255, 0.08);
  border-color: rgba(255, 255, 255, 0.2);
  color: #fff;
  transform: translateY(-1px);
}
</style>