<script setup lang="ts">
import { ref } from "vue"
import axios from "axios"
import type { User } from "@/types/user"
import router from "@/router"

const user = ref<User>({
    name: "",
    email: "",
    password: ""
})

const register = async () => {

    if (!user.value.name || !user.value.email || !user.value.password) {
        alert("Please fill all fields")
        return
    }

    try {

        const res = await axios.post(
            "http://localhost:3000/register",
            user.value
        )

        console.log("Status Code:", res.status)
        console.log("Response Data:", res.data)

        if (res.status === 201) {

            alert("Register Success")

            router.push("/")

        }

    } catch (error: unknown) {

        if (axios.isAxiosError(error)) {

            const status = error.response?.status

            if (status === 401) {
                alert("Email or Password incorrect")
            }

            if (status === 404) {
                alert("User not found")
            }

        } else {

            alert("Network error")

        }

    }

}
</script>

<template>

    <div class="register-page">

        <div class="register-card shadow-lg">

            <!-- LEFT -->
            <div class="register-form">

                <h3 class="text-success fw-bold mb-4 text-center">
                    Create Account
                </h3>

                <form @submit.prevent="register">

                    <input class="form-control mb-3" placeholder="Name" v-model="user.name" required />

                    <input class="form-control mb-3" type="email" placeholder="Email" v-model="user.email" required />

                    <input class="form-control mb-3" type="password" placeholder="Password" v-model="user.password"
                        required />

                    <button class="btn btn-success w-100" type="submit">
                        Register
                    </button>

                </form>

                <p class="text-center mt-3 small">
                    Already have an account?
                    <router-link to="/" class="text-success">
                        Login
                    </router-link>
                </p>

            </div>

            <!-- RIGHT -->
            <div class="register-welcome text-white text-center">

                <h1 class="fw-bold mb-3">
                    Join Website
                </h1>

                <p class="mb-4">
                    Create your account 
                </p>

                <router-link to="/" class="btn btn-outline-light">
                    Log in
                </router-link>

            </div>

        </div>

    </div>

</template>

<style scoped>
/* PAGE */

.register-page {
    position: fixed;
    inset: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    background: #e8f5e9;
    overflow: hidden;
    padding: 20px;
}

/* CARD */

.register-card {
    display: flex;
    width: 900px;
    max-width: 100%;
    background: white;
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
}

/* FORM */

.register-form {
    width: 50%;
    padding: 40px;
    display: flex;
    flex-direction: column;
    justify-content: center;
}

/* RIGHT PANEL */

.register-welcome {
    width: 50%;
    background: linear-gradient(135deg, #2e8b57, #3cb371);
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 40px;
}

/* TABLET */

@media (max-width:992px) {

    .register-card {
        width: 700px;
    }

}

/* MOBILE */

@media (max-width:768px) {

    .register-card {
        flex-direction: column;
    }

    .register-form {
        width: 100%;
        padding: 30px;
    }

    .register-welcome {
        width: 100%;
        padding: 30px;
    }

}

/* TABLET */

@media (max-width:992px) {

    .register-card {
        width: 700px;
    }

}


/* MOBILE */

@media (max-width:768px) {

    .register-card {
        flex-direction: column;
    }

    .register-form {
        width: 100%;
        padding: 30px;
    }

    .register-welcome {
        width: 100%;
        padding: 30px;
    }

}
</style>