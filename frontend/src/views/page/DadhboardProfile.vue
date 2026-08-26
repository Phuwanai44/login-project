<script setup lang="ts">
import { ref, onMounted } from "vue"
import { getProfile } from "@/services/authService"

const user = ref({
    name: "",
    email: "",
    role: ""
})

const loadProfile = async () => {
    try {
        const data = await getProfile()
        user.value = data
    } catch (error) {
        console.error(error)
    }
}

onMounted(() => {
    loadProfile()
})
</script>

<template>
    <div class="profile-content">
        <div class="d-flex justify-content-between align-items-center mb-4">
            <div>
                <h2 class="fw-bold mb-1 text-gradient">My Account</h2>
                <p class="text-muted small">Manage security settings and view profile records.</p>
            </div>
        </div>

        <div class="row g-4">
            <!-- Profile Card -->
            <div class="col-md-4">
                <div class="glass-card text-center p-4 h-100 d-flex flex-column justify-content-center align-items-center">
                    <div class="avatar-large mb-3">
                        <i class="bi bi-person-circle text-white"></i>
                    </div>

                    <h4 class="fw-bold mb-1">{{ user.name || 'User Profile' }}</h4>
                    <p class="text-muted small mb-3">{{ user.email }}</p>

                    <span class="badge glass-badge w-50 py-2" :class="user.role === 'admin' ? 'glass-badge-admin' : 'glass-badge-user'">
                        {{ user.role || 'Guest' }}
                    </span>
                </div>
            </div>

            <!-- Account Details -->
            <div class="col-md-8">
                <div class="glass-card p-4">
                    <h5 class="fw-bold mb-3 pb-2 border-bottom border-light border-opacity-10">Account Information</h5>

                    <div class="row mb-3 align-items-center">
                        <div class="col-sm-4 text-muted small fw-bold">FULL NAME</div>
                        <div class="col-sm-8 fw-semibold">{{ user.name }}</div>
                    </div>

                    <hr class="profile-hr">

                    <div class="row mb-3 align-items-center">
                        <div class="col-sm-4 text-muted small fw-bold">EMAIL ADDRESS</div>
                        <div class="col-sm-8 text-muted">{{ user.email }}</div>
                    </div>

                    <hr class="profile-hr">

                    <div class="row mb-2 align-items-center">
                        <div class="col-sm-4 text-muted small fw-bold">SECURITY ACCESS</div>
                        <div class="col-sm-8">
                            <span class="badge glass-badge" :class="user.role === 'admin' ? 'glass-badge-admin' : 'glass-badge-user'">
                                {{ user.role }}
                            </span>
                        </div>
                    </div>
                </div>

                <!-- System Info -->
                <div class="glass-card p-4 mt-4">
                    <h5 class="fw-bold mb-3 pb-2 border-bottom border-light border-opacity-10">System Information</h5>
                    <div class="row g-3 text-center text-sm-start">
                        <div class="col-sm-6">
                            <p class="mb-1 text-muted small">FRONTEND CORE</p>
                            <p class="fw-bold mb-0">Vue 3 + Vite + TypeScript</p>
                        </div>
                        <div class="col-sm-6">
                            <p class="mb-1 text-muted small">API ORCHESTRATOR</p>
                            <p class="fw-bold mb-0">Go Gin Framework</p>
                        </div>
                        <div class="col-sm-6">
                            <p class="mb-1 text-muted small">PERSISTENT CACHE</p>
                            <p class="fw-bold mb-0">MongoDB (NoSQL Cluster)</p>
                        </div>
                        <div class="col-sm-6">
                            <p class="mb-1 text-muted small">SIGNING ALGORITHM</p>
                            <p class="badge glass-badge glass-badge-success mb-0">JWT-HS256</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.profile-content {
    color: var(--text-primary);
}

.avatar-large {
    width: 100px;
    height: 100px;
    border-radius: 50%;
    background: radial-gradient(circle, rgba(139, 92, 246, 0.3) 0%, rgba(139, 92, 246, 0.05) 100%);
    border: 2px solid rgba(139, 92, 246, 0.4);
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 3rem;
    box-shadow: 0 0 20px rgba(139, 92, 246, 0.2);
}

.profile-hr {
    border-color: rgba(255, 255, 255, 0.06);
    margin: 15px 0;
}
</style>