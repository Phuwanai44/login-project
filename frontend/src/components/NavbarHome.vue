<script setup lang="ts">

</script>

<template>
  <nav class="navbar bg-dark text-white border-bottom px-4 d-flex justify-content-between align-items-center">

    <!-- Left -->
    <div>
      <h5 class="mb-0">Login System</h5>
    </div>

    <!-- Right -->
    <div>
      <router-link to="/login" class="btn btn-primary me-2">
        Login
      </router-link>

      <router-link to="/register" class="btn btn-outline-primary">
        Register
      </router-link>
    </div>

  </nav>
</template>

<style scoped>
.navbar {
    height: 60px;
}
</style>