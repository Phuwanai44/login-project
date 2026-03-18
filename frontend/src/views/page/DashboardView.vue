<script setup lang="ts">
import { ref, onMounted, computed } from "vue"
import { getUsers } from "@/services/userService"
import type { User } from "@/types/user"

const users = ref<User[]>([])
const loading = ref(true)

const loadUsers = async () => {
  try {
    const data = await getUsers()
    users.value = data
  } catch (err) {
    console.error(err)
  } finally {
    loading.value = false
  }
}

const stats = computed(() => [
  { title: "Total Users", value: users.value.length },
  { title: "Active Users", value: users.value.length },
  { title: "New Users", value: users.value.length },
  { title: "System Status", value: "Online" }
])

const recentUsers = computed(() => users.value.slice(0, 5))

onMounted(() => {
  loadUsers()
})
</script>

<template>

  <div class="container mt-4">

    <h2 class="mb-4">Dashboard</h2>

    <!-- Stats -->

    <div class="row mb-4">

      <div class="col-md-3" v-for="s in stats" :key="s.title">

        <div class="card shadow-sm h-100">

          <div class="card-body">

            <h6 class="text-muted">
              {{ s.title }}
            </h6>

            <h3 class="fw-bold">
              {{ s.value }}
            </h3>

          </div>

        </div>

      </div>

    </div>


    <!-- Users Table -->

    <div class="card shadow-sm mb-4">

      <div class="card-header d-flex justify-content-between align-items-center">

        <span class="fw-bold">Recent Users</span>

        <router-link to="/users" class="btn btn-sm btn-primary">
          Manage
        </router-link>

      </div>

      <div class="card-body">

        <div v-if="loading" class="text-center py-4">
          Loading users...
        </div>

        <table v-else class="table table-hover">

          <thead>

            <tr>
              <th>#</th>
              <th>Name</th>
              <th>Email</th>
            </tr>

          </thead>

          <tbody>

            <tr v-if="recentUsers.length === 0">

              <td colspan="3" class="text-center text-muted">
                No users found
              </td>

            </tr>

            <tr v-for="(user, index) in recentUsers" :key="user.id">

              <td>{{ index + 1 }}</td>
              <td>{{ user.name }}</td>
              <td>{{ user.email }}</td>

            </tr>

          </tbody>

        </table>

      </div>

    </div>


    <!-- Bottom Section -->

    <div class="row">

      <!-- System Info -->

        <div class="card shadow-sm">

          <div class="card-header fw-bold">
            System Information
          </div>

          <div class="card-body">

            <p><b>Frontend:</b> Vue 3 + TypeScript</p>
            <p><b>Backend:</b> Go + Gin</p>
            <p><b>Database:</b> MongoDB</p>
            <p><b>Authentication:</b> JWT</p>

          </div>

        </div>

      </div>

    </div>


</template>