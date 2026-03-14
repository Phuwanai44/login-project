<script setup lang="ts">
import { useRouter } from "vue-router"

const router = useRouter()

const logout = () => {

  localStorage.removeItem("token")
  localStorage.removeItem("role")

  router.push("/login")

}
</script>

<template>
  <nav class="navbar bg-white border-bottom px-3 d-flex justify-content-between align-items-center">

    <div class="d-flex align-items-center gap-3">
      <span class="fw-semibold">
        Dashboard
      </span>
    </div>

    <button class="btn btn-outline-danger btn-sm"
    @click="logout">
      Logout
    </button>

  </nav>
</template>

<style scoped>
.navbar {
  height: 60px;
}
</style>