<script setup lang="ts">
import { ref } from "vue"
import axios from "axios"
import type { User } from "@/types/user"
import { useRouter } from "vue-router"

const router = useRouter()
const user = ref<User>({
    name: "",
    email: "",
    password: ""
})
const loading = ref(false)
const errorMessage = ref("")
const successMessage = ref("")

const register = async () => {
    if (!user.value.name || !user.value.email || !user.value.password) {
        errorMessage.value = "Please fill in all fields"
        return
    }

    loading.value = true
    errorMessage.value = ""
    successMessage.value = ""
    try {
        const res = await axios.post(
            "http://localhost:3000/register",
            user.value
        )

        if (res.status === 201) {
            successMessage.value = "Registration successful! Redirecting to login..."
            setTimeout(() => {
                router.push("/login")
            }, 1500)
        }
    } catch (error: unknown) {
        if (axios.isAxiosError(error)) {
            const status = error.response?.status
            if (status === 400) {
                errorMessage.value = "Invalid details or email already registered."
            } else {
                errorMessage.value = error.response?.data?.message || "Registration failed. Please try again."
            }
        } else {
            errorMessage.value = "Network error. Please try again later."
        }
    } finally {
        loading.value = false
    }
}
</script>

<template>
    <div class="register-page">
        <!-- Glowing background orbs -->
        <div class="glow-orb-container">
          <div class="glow-orb orb-purple"></div>
          <div class="glow-orb orb-blue"></div>
        </div>

        <div class="register-card shadow-lg">
            <!-- LEFT - FORM -->
            <div class="register-form">
                <div class="mb-4 text-center text-md-start">
                    <span class="badge glass-badge glass-badge-success mb-2">Registration</span>
                    <h2 class="fw-bold mb-1 text-gradient">Create Account</h2>
                    <p class="text-muted small">Sign up to access the administrative system.</p>
                </div>

                <div v-if="errorMessage" class="alert alert-danger border-0 bg-danger bg-opacity-10 text-danger small py-2 px-3 rounded-3 mb-3">
                    <i class="bi bi-exclamation-triangle-fill me-2"></i>{{ errorMessage }}
                </div>

                <div v-if="successMessage" class="alert alert-success border-0 bg-success bg-opacity-10 text-success small py-2 px-3 rounded-3 mb-3">
                    <i class="bi bi-check-circle-fill me-2"></i>{{ successMessage }}
                </div>

                <form @submit.prevent="register">
                    <div class="mb-3 position-relative">
                        <label class="form-label text-muted small fw-bold">YOUR NAME</label>
                        <div class="input-icon-wrapper">
                            <i class="bi bi-person-fill input-icon"></i>
                            <input class="form-control glass-input" placeholder="John Doe" v-model="user.name" required />
                        </div>
                    </div>

                    <div class="mb-3 position-relative">
                        <label class="form-label text-muted small fw-bold">EMAIL ADDRESS</label>
                        <div class="input-icon-wrapper">
                            <i class="bi bi-envelope-fill input-icon"></i>
                            <input class="form-control glass-input" type="email" placeholder="name@company.com" v-model="user.email" required />
                        </div>
                    </div>

                    <div class="mb-4 position-relative">
                        <label class="form-label text-muted small fw-bold">PASSWORD</label>
                        <div class="input-icon-wrapper">
                            <i class="bi bi-shield-lock-fill input-icon"></i>
                            <input class="form-control glass-input" type="password" placeholder="••••••••" v-model="user.password" required />
                        </div>
                    </div>

                    <button class="btn glass-btn w-100 py-3" type="submit" :disabled="loading">
                        <span v-if="loading" class="spinner-border spinner-border-sm role-status me-2" aria-hidden="true"></span>
                        {{ loading ? 'Creating Account...' : 'Sign Up' }}
                    </button>
                </form>

                <p class="text-center mt-4 small text-muted">
                    Already have an account? 
                    <router-link to="/login" class="login-link text-decoration-none">
                        Sign In
                    </router-link>
                </p>
            </div>

            <!-- RIGHT - HERO DECORATION -->
            <div class="register-welcome text-white text-center">
                <div class="hero-glow"></div>
                <div class="welcome-content z-2">
                    <div class="badge-icon mb-4">
                        <i class="bi bi-rocket-takeoff text-white fs-2"></i>
                    </div>
                    <h1 class="fw-extrabold mb-3 title-main">AetherAuth</h1>
                    <p class="mb-4 text-slate small">
                        Join AetherAuth today to unlock real-time analytical panels, role customizers, and secure JWT-protected databases.
                    </p>
                    <p class="small text-muted mb-2">Already registered?</p>
                    <router-link to="/login" class="btn glass-btn-secondary w-75 py-2">
                        Log In Instead
                    </router-link>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.register-page {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 20px;
    z-index: 10;
}

.register-card {
    display: flex;
    width: 900px;
    min-height: 540px;
    background: rgba(18, 14, 38, 0.45);
    backdrop-filter: blur(25px);
    -webkit-backdrop-filter: blur(25px);
    border: 1px solid rgba(255, 255, 255, 0.08);
    border-radius: 20px;
    overflow: hidden;
}

.register-form {
    width: 50%;
    padding: 50px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    background: rgba(10, 8, 22, 0.3);
}

.input-icon-wrapper {
    position: relative;
}

.input-icon {
    position: absolute;
    left: 14px;
    top: 50%;
    transform: translateY(-50%);
    color: var(--text-muted);
    z-index: 5;
}

.glass-input {
    padding-left: 42px !important;
}

.login-link {
    font-weight: 600;
    color: var(--accent);
    transition: color 0.2s ease;
}

.login-link:hover {
    color: #a78bfa;
}

.register-welcome {
    width: 50%;
    background: radial-gradient(circle at 100% 100%, rgba(139, 92, 246, 0.25) 0%, rgba(13, 10, 30, 0.95) 100%);
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 50px;
    position: relative;
    border-left: 1px solid rgba(255, 255, 255, 0.06);
}

.hero-glow {
    position: absolute;
    width: 300px;
    height: 300px;
    background: radial-gradient(circle, rgba(139, 92, 246, 0.3) 0%, rgba(139, 92, 246, 0) 70%);
    filter: blur(40px);
    pointer-events: none;
    z-index: 1;
}

.welcome-content {
    position: relative;
}

.badge-icon {
    width: 60px;
    height: 60px;
    background: rgba(255, 255, 255, 0.04);
    border: 1px solid rgba(255, 255, 255, 0.15);
    border-radius: 16px;
    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.2);
    margin: 0 auto;
}

.title-main {
    font-family: 'Outfit', sans-serif;
    font-weight: 800;
    letter-spacing: -0.5px;
    background: var(--accent-gradient);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}

.text-slate {
    color: var(--text-secondary);
}

@media (max-width:768px) {
    .register-card {
        flex-direction: column;
        width: 100%;
        max-width: 450px;
    }
    .register-form {
        width: 100%;
        padding: 30px;
    }
    .register-welcome {
        width: 100%;
        padding: 35px;
        border-left: none;
        border-top: 1px solid rgba(255, 255, 255, 0.06);
    }
}
</style>