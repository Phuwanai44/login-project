<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import axios from 'axios'
import NavbarHome from '@/components/NavbarHome.vue'
import FooterHome from '@/components/FooterHome.vue'
import ProductCard from '@/components/ProductCard.vue'

interface ProductVariant {
  color: string
  size: string
  stock: number
}

interface Product {
  id: string
  name: string
  price: number
  category?: string
  image?: string
  variants?: ProductVariant[]
}

const products = ref<Product[]>([])
const loading = ref(true)
const searchQuery = ref('')
const selectedCategory = ref('')

const categories = ref(['T-Shirt', 'Pants', 'Jacket', 'Accessories'])

const fetchProducts = async () => {
  try {
    loading.value = true
    const response = await axios.get('http://localhost:8080/products')
    // Backend may not have new fields yet, so we map and add fallbacks if needed
    products.value = response.data.data.map((p: any) => ({
      id: p.id,
      name: p.name,
      price: p.price,
      category: p.category || categories.value[Math.floor(Math.random() * categories.value.length)],
      image: p.image || `https://via.placeholder.com/300x400?text=${encodeURIComponent(p.name)}`,
      variants: p.variants || [
        { color: 'Black', size: 'M', stock: 10 },
        { color: 'White', size: 'L', stock: 5 }
      ]
    }))
  } catch (error) {
    console.error('Failed to fetch products:', error)
    // Fallback dummy data if backend fails
    products.value = [
      { id: '1', name: 'Classic White T-Shirt', price: 590, category: 'T-Shirt', image: 'https://via.placeholder.com/300x400?text=White+T-Shirt', variants: [] },
      { id: '2', name: 'Black Denim Jeans', price: 1290, category: 'Pants', image: 'https://via.placeholder.com/300x400?text=Denim+Jeans', variants: [] },
    ]
  } finally {
    loading.value = false
  }
}

const filteredProducts = computed(() => {
  return products.value.filter(p => {
    const matchSearch = p.name.toLowerCase().includes(searchQuery.value.toLowerCase())
    const matchCategory = selectedCategory.value === '' || p.category === selectedCategory.value
    return matchSearch && matchCategory
  })
})

onMounted(() => {
  fetchProducts()
})
</script>

<template>
  <div class="page-wrapper">
    <div class="glow-orb-container">
      <div class="glow-orb orb-purple"></div>
      <div class="glow-orb orb-blue"></div>
      <div class="glow-orb orb-cyan"></div>
    </div>

    <NavbarHome />

    <main class="container mt-5 flex-grow-1">
      <div class="text-center mb-5">
        <h1 class="display-4 fw-extrabold text-gradient">Shop Collection</h1>
        <p class="lead text-muted">Discover our latest apparel</p>
      </div>

      <!-- Filters & Search -->
      <div class="glass-card p-4 mb-5">
        <div class="row g-3">
          <div class="col-md-6">
            <div class="input-group">
              <span class="input-group-text bg-transparent border-secondary text-white">
                <i class="bi bi-search"></i>
              </span>
              <input 
                type="text" 
                class="form-control bg-dark text-white border-secondary" 
                placeholder="Search products..." 
                v-model="searchQuery"
              >
            </div>
          </div>
          <div class="col-md-6">
            <select class="form-select bg-dark text-white border-secondary" v-model="selectedCategory">
              <option value="">All Categories</option>
              <option v-for="cat in categories" :key="cat" :value="cat">{{ cat }}</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Product Grid -->
      <div v-if="loading" class="text-center py-5">
        <div class="spinner-border text-primary" role="status">
          <span class="visually-hidden">Loading...</span>
        </div>
      </div>
      
      <div v-else-if="filteredProducts.length === 0" class="text-center py-5">
        <i class="bi bi-inbox fs-1 text-muted"></i>
        <h4 class="mt-3 text-white">No products found</h4>
      </div>

      <div v-else class="row g-4 mb-5">
        <div class="col-12 col-sm-6 col-md-4 col-lg-3" v-for="product in filteredProducts" :key="product.id">
          <ProductCard :product="product" />
        </div>
      </div>
    </main>

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

.form-control:focus, .form-select:focus {
  border-color: #8b5cf6;
  box-shadow: 0 0 0 0.25rem rgba(139, 92, 246, 0.25);
}
</style>
