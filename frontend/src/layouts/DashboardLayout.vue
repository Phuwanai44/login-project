<template>
  <div class="dashboard-wrapper">
    <!-- Glowing background orbs -->
    <div class="glow-orb-container">
      <div class="glow-orb orb-purple"></div>
      <div class="glow-orb orb-blue"></div>
      <div class="glow-orb orb-cyan"></div>
    </div>

    <!-- Sidebar -->
    <DashboardSidebar :isOpen="isSidebarOpen" @toggleSidebar="toggleSidebar" />

    <div class="main-layout">
      <!-- Navbar -->
      <NavbarDashbord @toggleSidebar="toggleSidebar" />

      <!-- Page Content -->
      <main class="content-container">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import DashboardSidebar from "@/components/DashboardSidebar.vue"
import NavbarDashbord from "@/components/NavbarDashbord.vue"

const isSidebarOpen = ref(true)

const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}
</script>

<style scoped>
.dashboard-wrapper {
  display: flex;
  min-height: 100vh;
  position: relative;
  overflow: hidden;
}

.main-layout {
  flex: 1;
  display: flex;
  flex-direction: column;
  min-width: 0;
  transition: all 0.3s ease;
}

.content-container {
  padding: 30px;
  flex: 1;
  overflow-y: auto;
}

/* Page transitions */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s ease, transform 0.25s ease;
}

.fade-enter-from {
  opacity: 0;
  transform: translateY(10px);
}

.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}
</style>