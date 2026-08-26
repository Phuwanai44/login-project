<template>
  <nav class="navbar glass-dash-nav px-4 d-flex justify-content-between align-items-center">
    <div class="d-flex align-items-center gap-3">
      <button class="nav-toggle-btn d-md-none" @click="emit('toggleSidebar')">
        <i class="bi bi-list"></i>
      </button>
      <span class="page-title fs-5">Control Center</span>
    </div>

    <div class="d-flex align-items-center gap-3">
      <!-- Status Badge -->
      <span class="status-indicator d-none d-sm-flex align-items-center gap-2">
        <span class="pulse-dot"></span>
        <span class="small text-muted text-uppercase fw-bold">Live Server</span>
      </span>

      <button class="logout-btn btn d-flex align-items-center gap-2" @click="logout">
        <i class="bi bi-box-arrow-right"></i>
        <span>Logout</span>
      </button>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router"

const router = useRouter()
const emit = defineEmits(["toggleSidebar"])

const logout = () => {
  localStorage.removeItem("token")
  localStorage.removeItem("role")
  router.push("/login")
}
</script>

<style scoped>
.glass-dash-nav {
  height: 70px;
  background: rgba(13, 10, 30, 0.4) !important;
  backdrop-filter: blur(12px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  padding: 0 24px;
}

.page-title {
  font-weight: 700;
  color: var(--text-primary);
  letter-spacing: 0.5px;
}

.nav-toggle-btn {
  background: transparent;
  border: none;
  color: var(--text-primary);
  font-size: 1.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
}

.status-indicator {
  background: rgba(16, 185, 129, 0.1);
  border: 1px solid rgba(16, 185, 129, 0.2);
  padding: 6px 12px;
  border-radius: 20px;
}

.pulse-dot {
  width: 8px;
  height: 8px;
  background: #10b981;
  border-radius: 50%;
  box-shadow: 0 0 8px #10b981;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% {
    transform: scale(0.95);
    box-shadow: 0 0 0 0 rgba(16, 185, 129, 0.7);
  }
  70% {
    transform: scale(1);
    box-shadow: 0 0 0 6px rgba(16, 185, 129, 0);
  }
  100% {
    transform: scale(0.95);
    box-shadow: 0 0 0 0 rgba(16, 185, 129, 0);
  }
}

.logout-btn {
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.2);
  color: #f87171 !important;
  padding: 8px 16px;
  border-radius: 8px;
  font-size: 0.9rem;
  font-weight: 600;
  transition: all 0.25s ease;
}

.logout-btn:hover {
  background: var(--danger-gradient);
  border-color: rgba(255, 255, 255, 0.1);
  color: #fff !important;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3);
}
</style>