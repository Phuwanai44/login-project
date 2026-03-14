<script setup lang="ts">

import { ref } from "vue"
import { useRouter } from "vue-router"
import { loginUser } from "@/services/authService"


const router = useRouter()
const email = ref("")
const password = ref("")

const login = async () => {

    try {

        const response = await loginUser(email.value, password.value)

        console.log("login response:", response)

        const token = response.token
        const role = response.user.role

        localStorage.setItem("token", token)
        localStorage.setItem("role", role)

        if (role === "admin") {
            router.push("/dashboard")
        } else {
            router.push("/home")
        }

    } catch (error) {

        console.error(error)

    }
}

</script>

<template>

    <div class="login-page">

        <div class="login-card shadow-lg">

            <!-- LEFT -->

            <div class="login-form">

                <h3 class="text-success fw-bold mb-4 text-center">
                    Log in
                </h3>

                <form @submit.prevent="login">

                    <input type="email" class="form-control mb-3" placeholder="Email" v-model="email" />

                    <input type="password" class="form-control mb-3" placeholder="Password" v-model="password" />

                    <div class="d-flex justify-content-between align-items-center mb-3">

                        <div>
                            <input type="checkbox" class="form-check-input me-1">
                            Remember me
                        </div>

                    </div>

                    <button class="btn btn-success w-100" type="submit">
                        Log in
                    </button>

                </form>

            </div>

            <!-- RIGHT -->

            <div class="login-welcome text-white text-center">

                <h1 class="fw-bold mb-3">
                    Welcome Back!
                </h1>

                <p class="mb-4">
                    Please enter your details
                </p>

                <p class="small mb-3">
                    Don't have an account?
                </p>

                <router-link to="/register" class="btn btn-outline-light">
                    Sign Up
                </router-link>

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
    background: #e8f5e9;
}

.login-card {
    display: flex;
    width: 900px;
    background: white;
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
}

.login-form {
    width: 50%;
    padding: 40px;
    display: flex;
    flex-direction: column;
    justify-content: center;
}

.login-welcome {
    width: 50%;
    background: linear-gradient(135deg, #2e8b57, #3cb371);
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 40px;
}

@media (max-width:768px) {

    .login-card {
        flex-direction: column;
    }

    .login-form {
        width: 100%;
    }

    .login-welcome {
        width: 100%;
    }

}
</style>