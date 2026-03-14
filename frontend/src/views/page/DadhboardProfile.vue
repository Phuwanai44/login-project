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

    <div class="container-fluid d-flex flex-column align-items-center justify-content-center" style="min-height:80vh;">

        <h2 class="mb-4 text-center">My Profile</h2>

        <div class="card shadow-lg p-4" style="max-width:500px; width:100%;">

            <div class="text-center">

                <img src="https://cdn-icons-png.flaticon.com/512/149/149071.png" class="rounded-circle mb-3"
                    width="120" />

                <h4 class="mb-4">{{ user.name }}</h4>

            </div>

            <div>

                <p class="mb-2">
                    <b>Email :</b>
                    <span class="text-muted">{{ user.email }}</span>
                </p>

                <p class="mb-4">
                    <b>Role :</b>

                    <span class="badge" :class="user.role === 'admin' ? 'bg-danger' : 'bg-primary'">
                        {{ user.role }}
                    </span>

                </p>

            </div>

            <div class="text-center">

                <button class="btn btn-outline-primary">
                    Edit Profile
                </button>

            </div>

        </div>

    </div>

</template>