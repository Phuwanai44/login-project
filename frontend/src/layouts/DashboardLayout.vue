<template>
  <div class="dashboard">

    <!-- Sidebar -->
    <DashboardSidebar :isOpen="isSidebarOpen" @toggleSidebar="toggleSidebar" />

    <div class="main">

      <!-- Navbar -->
      <NavbarDashbord />

      <!-- Page Content -->
      <div class="content">
        <router-view />
      </div>

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
.dashboard {
  display: flex;
  min-height: 100vh;
}

.main {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.content {
  padding: 24px;
  background: #f5f6fa;
  flex: 1;
}
</style>