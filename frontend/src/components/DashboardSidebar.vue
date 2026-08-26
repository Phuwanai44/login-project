<template>
  <div :class="['sidebar-glass-panel', { close: !isOpen }]">
    <div class="sidebar-header d-flex justify-content-between align-items-center mb-4">
      <h5 v-if="isOpen" class="mb-0 brand-logo">
        <i class="bi bi-shield-lock-fill me-2"></i>AetherAuth
      </h5>
      <button class="toggle-btn" @click="emit('toggleSidebar')">
        <i class="bi" :class="isOpen ? 'bi-chevron-left' : 'bi-list'"></i>
      </button>
    </div>

    <ul class="nav flex-column sidebar-nav">
      <li class="nav-item">
        <router-link class="nav-link-glass d-flex align-items-center gap-3" to="/dashboard">
          <i class="bi bi-grid-1x2-fill"></i>
          <span v-if="isOpen">Dashboard</span>
        </router-link>
      </li>

      <li class="nav-item">
        <router-link class="nav-link-glass d-flex align-items-center gap-3" to="/profile">
          <i class="bi bi-person-circle"></i>
          <span v-if="isOpen">Profile</span>
        </router-link>
      </li>

      <li class="nav-item">
        <router-link class="nav-link-glass d-flex align-items-center gap-3" to="/users">
          <i class="bi bi-people-fill"></i>
          <span v-if="isOpen">Users</span>
        </router-link>
      </li>

      <li class="nav-item">
        <router-link class="nav-link-glass d-flex align-items-center gap-3" to="/products">
          <i class="bi bi-box-seam-fill"></i>
          <span v-if="isOpen">Products</span>
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits(["toggleSidebar"])
</script>

<style scoped>
.sidebar-glass-panel {
  width: 260px;
  min-height: 100vh;
  background: rgba(13, 10, 30, 0.6);
  backdrop-filter: blur(20px);
  border-right: 1px solid rgba(255, 255, 255, 0.08);
  padding: 24px 16px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  flex-direction: column;
  z-index: 100;
}

.sidebar-glass-panel.close {
  width: 80px;
  padding: 24px 12px;
}

.brand-logo {
  font-weight: 800;
  letter-spacing: 0.5px;
  background: var(--accent-gradient);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  display: flex;
  align-items: center;
}

.brand-logo i {
  color: var(--accent);
  -webkit-text-fill-color: initial;
}

.toggle-btn {
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: var(--text-secondary);
  border-radius: 8px;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s ease;
}

.toggle-btn:hover {
  background: rgba(255, 255, 255, 0.15);
  color: #fff;
  border-color: rgba(255, 255, 255, 0.2);
}

.sidebar-nav {
  margin-top: 20px;
}

.sidebar-nav .nav-item {
  margin-bottom: 8px;
}

.nav-link-glass {
  color: var(--text-secondary);
  padding: 12px 16px;
  border-radius: 10px;
  text-decoration: none;
  font-weight: 500;
  display: flex;
  align-items: center;
  transition: all 0.25s ease;
  border: 1px solid transparent;
}

.nav-link-glass i {
  font-size: 1.2rem;
  transition: transform 0.25s ease;
}

.nav-link-glass:hover {
  background: rgba(255, 255, 255, 0.05);
  color: #fff;
  border-color: rgba(255, 255, 255, 0.05);
}

.nav-link-glass:hover i {
  transform: scale(1.1);
}

/* Active route styling */
.nav-link-glass.router-link-active {
  background: rgba(139, 92, 246, 0.15);
  color: #c084fc;
  border-color: rgba(139, 92, 246, 0.2);
  box-shadow: 0 0 15px rgba(139, 92, 246, 0.1);
  font-weight: 600;
}

.sidebar-glass-panel.close .nav-link-glass {
  justify-content: center;
  padding: 12px 0;
}
</style>