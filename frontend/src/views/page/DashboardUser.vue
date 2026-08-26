<script setup lang="ts">
import { ref, onMounted } from "vue"
import { getUsers, deleteUserById } from "@/services/userService"
import type { User } from "@/types/user"

const users = ref<User[]>([])
const loading = ref(true)

const loadUsers = async () => {
    loading.value = true
    try {
        const data = await getUsers()
        users.value = data
    } catch (err) {
        console.error(err)
    } finally {
        loading.value = false
    }
}

onMounted(() => {
    loadUsers()
})

const deleteUser = async (id?: string) => {
    if (!id) return
    const confirmDelete = confirm("Are you sure you want to delete this user?")
    if (!confirmDelete) return

    try {
        await deleteUserById(id)
        users.value = users.value.filter(u => u.id !== id)
    } catch (err) {
        console.error(err)
    }
}
</script>

<template>
    <div class="users-content">
        <div class="d-flex justify-content-between align-items-center mb-4">
            <div>
                <h2 class="fw-bold mb-1 text-gradient">Users Database</h2>
                <p class="text-muted small">Manage system users and administrative access roles.</p>
            </div>
        </div>

        <div class="glass-card">
            <div class="card-header-glass px-4 py-3 d-flex justify-content-between align-items-center">
                <h5 class="fw-bold mb-0">Registered Users</h5>
                <span class="badge glass-badge glass-badge-success">{{ users.length }} Users Active</span>
            </div>

            <div class="p-3">
                <div v-if="loading" class="text-center py-5">
                    <div class="spinner-border text-primary spinner-border-sm me-2" role="status"></div>
                    <span class="text-muted small">Loading user databases...</span>
                </div>

                <div v-else class="table-responsive glass-table-container">
                    <table class="glass-table">
                        <thead>
                            <tr>
                                <th>#</th>
                                <th>Name</th>
                                <th>Email</th>
                                <th class="text-end">Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-if="users.length === 0">
                                <td colspan="4" class="text-center text-muted small py-4">
                                    No registered users found
                                </td>
                            </tr>
                            <tr v-for="(user, index) in users" :key="user.id">
                                <td>{{ index + 1 }}</td>
                                <td class="fw-semibold">
                                    <div class="d-flex align-items-center gap-2">
                                        <div class="avatar-sm">
                                            {{ user.name.charAt(0).toUpperCase() }}
                                        </div>
                                        <span>{{ user.name }}</span>
                                    </div>
                                </td>
                                <td class="text-muted">{{ user.email }}</td>
                                <td class="text-end">
                                    <button class="btn glass-btn-danger btn-sm" @click="deleteUser(user.id)">
                                        <i class="bi bi-trash-fill me-1"></i>Delete
                                    </button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.users-content {
    color: var(--text-primary);
}

.card-header-glass {
    border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.avatar-sm {
    width: 28px;
    height: 28px;
    border-radius: 50%;
    background: rgba(139, 92, 246, 0.2);
    color: #c084fc;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 0.8rem;
    font-weight: 700;
    border: 1px solid rgba(139, 92, 246, 0.3);
}
</style>