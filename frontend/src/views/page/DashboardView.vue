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
  { title: "Total Users", value: users.value.length, icon: "bi-people", color: "purple" },
  { title: "Active Users", value: users.value.length, icon: "bi-activity", color: "green" },
  { title: "New Users", value: users.value.length, icon: "bi-person-plus", color: "cyan" },
  { title: "System Status", value: "Online", icon: "bi-hdd-network", color: "success" }
])

const recentUsers = computed(() => users.value.slice(0, 5))

onMounted(() => {
  loadUsers()
})
</script>

<template>
  <div class="dashboard-content">
    <div class="d-flex justify-content-between align-items-center mb-4">
      <div>
        <h2 class="fw-bold mb-1 text-gradient">Console Overview</h2>
        <p class="text-muted small">Real-time status updates and user logs.</p>
      </div>
    </div>

    <!-- Stats Grid -->
    <div class="row g-4 mb-4">
      <div class="col-md-3 col-sm-6" v-for="s in stats" :key="s.title">
        <div class="glass-card stat-card p-4">
          <div class="d-flex justify-content-between align-items-start">
            <div>
              <p class="stat-label small fw-semibold text-uppercase mb-2">
                {{ s.title }}
              </p>
              <h3 class="fw-extrabold mb-0 val-text">
                {{ s.value }}
              </h3>
            </div>
            <div class="stat-icon-wrapper" :class="s.color">
              <i class="bi" :class="s.icon"></i>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Grid -->
    <div class="row g-4">
      <!-- Users Table -->
      <div class="col-lg-8">
        <div class="glass-card h-100">
          <div class="card-header-glass d-flex justify-content-between align-items-center px-4 py-3">
            <h5 class="fw-bold mb-0">Recent Users</h5>
            <router-link to="/users" class="btn glass-btn-secondary py-1 px-3 fs-7">
              Manage All
            </router-link>
          </div>

          <div class="p-3">
            <div v-if="loading" class="text-center py-5">
              <div class="spinner-border text-primary spinner-border-sm me-2" role="status"></div>
              <span class="text-muted small">Retrieving user roster...</span>
            </div>

            <div v-else class="table-responsive glass-table-container">
              <table class="glass-table">
                <thead>
                  <tr>
                    <th>#</th>
                    <th>Name</th>
                    <th>Email</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-if="recentUsers.length === 0">
                    <td colspan="3" class="text-center text-muted small py-4">
                      No registered users found
                    </td>
                  </tr>
                  <tr v-for="(user, index) in recentUsers" :key="user.id">
                    <td>{{ index + 1 }}</td>
                    <td class="fw-semibold">{{ user.name }}</td>
                    <td class="td-email">{{ user.email }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <!-- System Info -->
      <div class="col-lg-4">
        <div class="glass-card h-100">
          <div class="card-header-glass px-4 py-3">
            <h5 class="fw-bold mb-0">System Specifications</h5>
          </div>

          <div class="p-4 spec-body">
            <div class="spec-item mb-3">
              <span class="text-muted small d-block">FRONTEND LAYOUT</span>
              <span class="fw-bold">Vue 3 + Tailwind CSS / BootStrap</span>
            </div>
            <hr class="spec-hr">
            <div class="spec-item mb-3">
              <span class="text-muted small d-block">BACKEND SERVER</span>
              <span class="fw-bold">Go Gin API (v1.10)</span>
            </div>
            <hr class="spec-hr">
            <div class="spec-item mb-3">
              <span class="text-muted small d-block">DATABASE ROUTING</span>
              <span class="fw-bold">MongoDB Cluster</span>
            </div>
            <hr class="spec-hr">
            <div class="spec-item">
              <span class="text-muted small d-block">AUTHENTICATION GATE</span>
              <span class="badge glass-badge glass-badge-success">JWT-SECURED</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dashboard-content {
  color: var(--text-primary);
}

.stat-card {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.val-text {
  font-family: 'Outfit', sans-serif;
  letter-spacing: -0.5px;
}

.stat-icon-wrapper {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(255, 255, 255, 0.08);
  font-size: 1.3rem;
}

.stat-icon-wrapper.purple {
  background: rgba(139, 92, 246, 0.15);
  color: #c084fc;
  box-shadow: 0 0 12px rgba(139, 92, 246, 0.25);
}

.stat-icon-wrapper.green {
  background: rgba(16, 185, 129, 0.15);
  color: #34d399;
  box-shadow: 0 0 12px rgba(16, 185, 129, 0.25);
}

.stat-icon-wrapper.cyan {
  background: rgba(6, 182, 212, 0.15);
  color: #22d3ee;
  box-shadow: 0 0 12px rgba(6, 182, 212, 0.25);
}

.stat-icon-wrapper.success {
  background: rgba(16, 185, 129, 0.15);
  color: #34d399;
  box-shadow: 0 0 12px rgba(16, 185, 129, 0.25);
}

.card-header-glass {
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.fs-7 {
  font-size: 0.8rem;
}

.spec-body {
  display: flex;
  flex-direction: column;
  height: calc(100% - 60px);
  justify-content: space-between;
}

.spec-hr {
  border-color: rgba(255, 255, 255, 0.06);
  margin: 12px 0;
}

/* Stat card label – visible light slate on dark bg */
.stat-label {
  color: #94a3b8;
  letter-spacing: 0.06em;
}

/* Email column – slightly dimmed but readable */
.td-email {
  color: #cbd5e1 !important;
  font-size: 0.88rem;
}
</style>