<script setup lang="ts">
import FooterHome from '@/components/FooterHome.vue'
import NavbarHome from '@/components/NavbarHome.vue'
</script>

<template>

    <div class="page-wrapper">

        <!-- Navbar -->
        <NavbarHome />

        <!-- Main Content -->
        <main class="container mt-5 flex-grow-1">

            <div class="text-center mb-5">
                <h1>Login System</h1>
                <p class="text-muted">
                    Fullstack authentication system built with Vue + Go + MongoDB
                </p>

               
            </div>

            <div class="row text-center">

                <div class="col-md-3">
                    <h5>JWT Auth</h5>
                    <p>Secure authentication using token</p>
                </div>

                <div class="col-md-3">
                    <h5>User CRUD</h5>
                    <p>Manage users in dashboard</p>
                </div>

                <div class="col-md-3">
                    <h5>Go Backend</h5>
                    <p>Fast API using Gin</p>
                </div>

                <div class="col-md-3">
                    <h5>Vue Frontend</h5>
                    <p>Modern UI with Vue 3</p>
                </div>

            </div>

        </main>

        <!-- Footer -->
        <FooterHome />

    </div>

</template>

<style>
.page-wrapper {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
}

.flex-grow-1 {
    flex: 1;
}
</style>
