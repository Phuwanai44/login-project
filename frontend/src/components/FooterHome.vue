<template>
  <footer class="glass-footer py-4 mt-auto">
    <div class="container text-center">
      <p class="mb-1 footer-title fw-bold">AetherAuth</p>
      <p class="mb-2 text-muted small">
        Fullstack authentication system built with Vue 3, Go, and MongoDB.
      </p>
      <div class="footer-divider mx-auto my-2"></div>
      <small class="text-muted">© 2026 Phuwanai. Crafted with modern web standards.</small>
    </div>
  </footer>
</template>

<style scoped>
.glass-footer {
  background: rgba(13, 10, 30, 0.4) !important;
  backdrop-filter: blur(8px);
  border-top: 1px solid rgba(255, 255, 255, 0.05);
  color: var(--text-secondary);
}

.footer-title {
  background: var(--accent-gradient);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  letter-spacing: 1px;
}

.footer-divider {
  width: 50px;
  height: 2px;
  background: var(--accent-gradient);
  border-radius: 2px;
  opacity: 0.3;
}
</style>