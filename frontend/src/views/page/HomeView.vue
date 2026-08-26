<script setup lang="ts">
import FooterHome from '@/components/FooterHome.vue'
import NavbarHome from '@/components/NavbarHome.vue'
</script>

<template>
  <div class="page-wrapper">
    <!-- Glowing background orbs -->
    <div class="glow-orb-container">
      <div class="glow-orb orb-purple"></div>
      <div class="glow-orb orb-blue"></div>
      <div class="glow-orb orb-cyan"></div>
    </div>

    <!-- Navbar -->
    <NavbarHome />

    <!-- Main Content -->
    <main class="container mt-5 flex-grow-1 d-flex flex-column justify-content-center">
      <div class="text-center hero-section mb-5 py-4">
        <span class="badge glass-badge glass-badge-success mb-3">Enterprise Ready</span>
        <h1 class="display-3 fw-extrabold text-gradient mb-3">
          AetherSecure Gateway
        </h1>
        <p class="lead text-muted mx-auto max-width-600">
          A full-stack secure authentication and administrative dashboard system built using Vue 3, Go (Gin), and MongoDB.
        </p>

        <div class="d-flex justify-content-center gap-3 mt-4">
          <router-link to="/login" class="btn glass-btn px-4 py-3">
            Go to Console
          </router-link>
          <router-link to="/register" class="btn glass-btn-secondary px-4 py-3">
            Register Account
          </router-link>
        </div>
      </div>

      <div class="row g-4 text-center mt-4 mb-5">
        <div class="col-md-3 col-sm-6">
          <div class="glass-card p-4 h-100 d-flex flex-column align-items-center">
            <div class="icon-wrapper mb-3 purple-glow">
              <i class="bi bi-key-fill text-white fs-3"></i>
            </div>
            <h5 class="fw-bold mb-2">JWT Authentication</h5>
            <p class="text-muted small mb-0">
              State-of-the-art secure token verification, cookie rotation, and route protection.
            </p>
          </div>
        </div>

        <div class="col-md-3 col-sm-6">
          <div class="glass-card p-4 h-100 d-flex flex-column align-items-center">
            <div class="icon-wrapper mb-3 cyan-glow">
              <i class="bi bi-people-fill text-white fs-3"></i>
            </div>
            <h5 class="fw-bold mb-2">User Administration</h5>
            <p class="text-muted small mb-0">
              Perform CRUD operations and manage profile settings inside an interactive panel.
            </p>
          </div>
        </div>

        <div class="col-md-3 col-sm-6">
          <div class="glass-card p-4 h-100 d-flex flex-column align-items-center">
            <div class="icon-wrapper mb-3 blue-glow">
              <i class="bi bi-cpu-fill text-white fs-3"></i>
            </div>
            <h5 class="fw-bold mb-2">Go Gin API</h5>
            <p class="text-muted small mb-0">
              Lightning fast backend built using Go Gin framework with structured middleware.
            </p>
          </div>
        </div>

        <div class="col-md-3 col-sm-6">
          <div class="glass-card p-4 h-100 d-flex flex-column align-items-center">
            <div class="icon-wrapper mb-3 green-glow">
              <i class="bi bi-layers-fill text-white fs-3"></i>
            </div>
            <h5 class="fw-bold mb-2">Vue 3 Vapor</h5>
            <p class="text-muted small mb-0">
              Lightweight reactive frontend with high-performance styling and component tree.
            </p>
          </div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <FooterHome />
  </div>
</template>

<style scoped>
.page-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  position: relative;
}

.max-width-600 {
  max-width: 600px;
}

.hero-section h1 {
  font-family: 'Outfit', sans-serif;
  font-weight: 800;
  letter-spacing: -1.5px;
}

.icon-wrapper {
  width: 60px;
  height: 60px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 20px;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.purple-glow {
  background: rgba(139, 92, 246, 0.2);
  box-shadow: 0 0 15px rgba(139, 92, 246, 0.3);
}

.cyan-glow {
  background: rgba(6, 182, 212, 0.2);
  box-shadow: 0 0 15px rgba(6, 182, 212, 0.3);
}

.blue-glow {
  background: rgba(37, 99, 235, 0.2);
  box-shadow: 0 0 15px rgba(37, 99, 235, 0.3);
}

.green-glow {
  background: rgba(16, 185, 129, 0.2);
  box-shadow: 0 0 15px rgba(16, 185, 129, 0.3);
}
</style>
