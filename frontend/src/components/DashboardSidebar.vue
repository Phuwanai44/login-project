<script setup lang="ts">

defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits(["toggleSidebar"])

</script>

<template>

  <div :class="['sidebar', { close: !isOpen }]">

    <div class="sidebar-header d-flex justify-content-between align-items-center mb-4">

      <h5 v-if="isOpen" class="mb-0 text-white fw-semibold">
        Dashboard
      </h5>

      <button class="btn btn-dark btn-sm" @click="emit('toggleSidebar')">
        ☰
      </button>

    </div>

    <ul class="nav flex-column">

      <li class="nav-item">
        <router-link class="nav-link text-white d-flex align-items-center gap-2" to="/dashboard">
          <span class="icon"></span>
          <span v-if="isOpen">Dashboard</span>
        </router-link>
      </li>

      <li class="nav-item">
        <router-link class="nav-link text-white d-flex align-items-center gap-2" to="/profile">
          <span class="icon"></span>
          <span v-if="isOpen">Profile</span>
        </router-link>
      </li>

      <li class="nav-item">
        <router-link class="nav-link text-white d-flex align-items-center gap-2" to="/users">
          <span class="icon"></span>
          <span v-if="isOpen">Users</span>
        </router-link>
      </li>

      <li class="nav-item">
        <router-link class="nav-link text-white d-flex align-items-center gap-2" to="/products">
          <span class="icon"></span>
          <span v-if="isOpen">product</span>
        </router-link>
      </li>

    </ul>

  </div>

</template>

<style scoped>

.sidebar{
  width:240px;
  min-height:100vh;
  background:#1f2937;
  padding:20px;
  transition:0.3s;
}

.sidebar.close{
  width:70px;
}

.sidebar.close span{
  display:none;
}

.nav-link:hover{
  background:#374151;
  border-radius:6px;
}

.icon{
  width:24px;
  text-align:center;
}

</style>