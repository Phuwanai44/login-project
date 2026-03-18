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

    <div class="container mt-5">

        <h2 class="mb-4">My Profile</h2>

        <div class="row g-4">

            <!-- Profile Card -->

            <div class="col-md-4">

                <div class="card shadow text-center p-4">

                    <img src="https://cdn-icons-png.flaticon.com/512/149/149071.png" class="rounded-circle mx-auto mb-3"
                        width="120" />

                    <h4>{{ user.name }}</h4>

                    <p class="text-muted">{{ user.email }}</p>

                    <span class="btn btn-outline text-white  w-100 mb-3" :class="user.role === 'admin' ? 'bg-danger' : 'bg-primary'">

                        {{ user.role }}

                    </span>

                    <!-- <button class="btn btn-outline-primary w-100">
                        Edit Profile
                    </button> -->

                </div>

            </div>


            <!-- Account Details -->

            <div class="col-md-8">

                <div class="card shadow">

                    <div class="card-header">
                        Account Information
                    </div>

                    <div class="card-body">

                        <div class="row mb-3">

                            <div class="col-sm-4 text-muted">
                                Name
                            </div>

                            <div class="col-sm-8">
                                {{ user.name }}
                            </div>

                        </div>

                        <hr>

                        <div class="row mb-3">

                            <div class="col-sm-4 text-muted">
                                Email
                            </div>

                            <div class="col-sm-8">
                                {{ user.email }}
                            </div>

                        </div>

                        <hr>

                        <div class="row mb-3">

                            <div class="col-sm-4 text-muted">
                                Role
                            </div>

                            <div class="col-sm-8">

                                <span class="badge" :class="user.role === 'admin' ? 'bg-danger' : 'bg-primary'">
                                    {{ user.role }}
                                </span>

                            </div>

                        </div>

                    </div>

                </div>


                <!-- System Info -->

                <div class="card shadow mt-4">

                    <div class="card-header">
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

    </div>

</template>