<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"
import { loginUser } from "@/services/authService"

const router = useRouter()
const email = ref("")
const password = ref("")
const loading = ref(false)
const errorMessage = ref("")

const login = async () => {
    loading.value = true
    errorMessage.value = ""
    try {
        const response = await loginUser(email.value, password.value)
        const token = response.token
        const role = response.user.role

        localStorage.setItem("token", token)
        localStorage.setItem("role", role)

        if (role === "admin") {
            router.push("/dashboard")
        } else {
            router.push("/dashboard") // Defer to dashboard as main console
        }
    } catch (error) {
        console.error(error)
        errorMessage.value = "Invalid email or password. Please try again."
    } finally {
        loading.value = false
    }
}
</script>

<template>
    <div class="login-page">
        <!-- Glowing background orbs -->
        <div class="glow-orb-container">
          <div class="glow-orb orb-purple"></div>
          <div class="glow-orb orb-blue"></div>
        </div>

        <div class="login-card shadow-lg">
            <!-- LEFT - FORM -->
            <div class="login-form">
                <div class="mb-4 text-center text-md-start">
                    <span class="badge glass-badge glass-badge-success mb-2">Secure Entry</span>
                    <h2 class="fw-bold mb-1 text-gradient">Welcome back</h2>
                    <p class="text-muted small">Access the administrative system console.</p>
                </div>

                <div v-if="errorMessage" class="alert alert-danger border-0 bg-danger bg-opacity-10 text-danger small py-2 px-3 rounded-3 mb-3">
                    <i class="bi bi-exclamation-triangle-fill me-2"></i>{{ errorMessage }}
                </div>

                <form @submit.prevent="login">
                    <div class="mb-3 position-relative">
                        <label class="form-label text-muted small fw-bold">EMAIL ADDRESS</label>
                        <div class="input-icon-wrapper">
                            <i class="bi bi-envelope-fill input-icon"></i>
                            <input type="email" class="form-control glass-input" placeholder="name@company.com" v-model="email" required />
                        </div>
                    </div>

                    <div class="mb-3 position-relative">
                        <label class="form-label text-muted small fw-bold">PASSWORD</label>
                        <div class="input-icon-wrapper">
                            <i class="bi bi-shield-lock-fill input-icon"></i>
                            <input type="password" class="form-control glass-input" placeholder="••••••••" v-model="password" required />
                        </div>
                    </div>

                    <div class="d-flex justify-content-between align-items-center mb-4">
                        <div class="form-check">
                            <input type="checkbox" class="form-check-input glass-checkbox" id="rememberMe">
                            <label class="form-check-label text-muted small" for="rememberMe">Remember me</label>
                        </div>
                        <a href="#" class="small-link text-decoration-none">Forgot password?</a>
                    </div>

                    <button class="btn glass-btn w-100 py-3" type="submit" :disabled="loading">
                        <span v-if="loading" class="spinner-border spinner-border-sm me-2" role="status" aria-hidden="true"></span>
                        {{ loading ? 'Authenticating...' : 'Sign In' }}
                    </button>
                </form>
            </div>

            <!-- RIGHT - HERO DECORATION -->
            <div class="login-welcome text-white text-center">
                <div class="hero-glow"></div>
                <div class="welcome-content z-2">
                    <div class="badge-icon mb-4">
                        <i class="bi bi-cpu text-white fs-2"></i>
                    </div>
                    <h1 class="fw-extrabold mb-3 title-main">AetherAuth</h1>
                    <p class="mb-4 text-slate small">
                        Manage users, view system performance, and customize client roles inside our high fidelity control dashboard.
                    </p>
                    <p class="small text-muted mb-2">New here?</p>
                    <router-link to="/register" class="btn glass-btn-secondary w-75 py-2">
                        Create an Account
                    </router-link>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.login-page {
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

.login-card {
    display: flex;
    width: 900px;
    min-height: 520px;
    background: rgba(18, 14, 38, 0.45);
    backdrop-filter: blur(25px);
    -webkit-backdrop-filter: blur(25px);
    border: 1px solid rgba(255, 255, 255, 0.08);
    border-radius: 20px;
    overflow: hidden;
}

.login-form {
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

.small-link {
    font-size: 0.85rem;
    color: var(--accent);
    transition: color 0.2s ease;
}

.small-link:hover {
    color: #a78bfa;
}

.login-welcome {
    width: 50%;
    background: radial-gradient(circle at 0% 0%, rgba(139, 92, 246, 0.25) 0%, rgba(13, 10, 30, 0.95) 100%);
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

.glass-checkbox {
    background-color: rgba(255, 255, 255, 0.05);
    border: 1px solid var(--glass-border);
    transition: all 0.2s ease;
}

.glass-checkbox:checked {
    background-color: var(--accent);
    border-color: var(--accent);
}

@media (max-width:768px) {
    .login-card {
        flex-direction: column;
        width: 100%;
        max-width: 450px;
    }
    .login-form {
        width: 100%;
        padding: 30px;
    }
    .login-welcome {
        width: 100%;
        padding: 35px;
        border-left: none;
        border-top: 1px solid rgba(255, 255, 255, 0.06);
    }
}
</style>