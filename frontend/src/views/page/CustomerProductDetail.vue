<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'
import { useCartStore } from '@/stores/cart'
import NavbarHome from '@/components/NavbarHome.vue'
import FooterHome from '@/components/FooterHome.vue'

const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()

interface ProductVariant {
  color: string
  size: string
  stock: number
}

interface Product {
  id: string
  name: string
  price: number
  description?: string
  category?: string
  image?: string
  variants?: ProductVariant[]
}

const product = ref<Product | null>(null)
const loading = ref(true)

const selectedColor = ref<string>('')
const selectedSize = ref<string>('')
const quantity = ref(1)

const fetchProduct = async () => {
  try {
    loading.value = true
    const response = await axios.get(`http://localhost:8080/products/${route.params.id}`)
    const p = response.data
    
    // Add dummy data for fields missing from backend
    product.value = {
      id: p.id,
      name: p.name,
      price: p.price,
      description: p.description || 'Premium quality apparel made for everyday comfort and style.',
      category: p.category || 'Apparel',
      image: p.image || `https://via.placeholder.com/600x800?text=${encodeURIComponent(p.name)}`,
      variants: p.variants || [
        { color: 'Black', size: 'S', stock: 5 },
        { color: 'Black', size: 'M', stock: 10 },
        { color: 'White', size: 'M', stock: 2 },
        { color: 'White', size: 'L', stock: 0 },
      ]
    }
  } catch (error) {
    console.error('Failed to fetch product:', error)
    // Fallback if not found
    if (!product.value) {
      router.push('/store')
    }
  } finally {
    loading.value = false
  }
}

const availableColors = computed(() => {
  if (!product.value?.variants) return []
  const colors = new Set(product.value.variants.map(v => v.color))
  return Array.from(colors)
})

const availableSizes = computed(() => {
  if (!product.value?.variants || !selectedColor.value) return []
  return product.value.variants
    .filter(v => v.color === selectedColor.value)
    .map(v => v.size)
})

const currentVariant = computed(() => {
  if (!product.value?.variants || !selectedColor.value || !selectedSize.value) return null
  return product.value.variants.find(
    v => v.color === selectedColor.value && v.size === selectedSize.value
  )
})

const handleColorSelect = (color: string) => {
  selectedColor.value = color
  selectedSize.value = '' // Reset size when color changes
  quantity.value = 1
}

const handleSizeSelect = (size: string) => {
  selectedSize.value = size
  quantity.value = 1
}

const addToCart = () => {
  if (!currentVariant.value) return
  if (currentVariant.value.stock < quantity.value) return
  
  cartStore.addToCart(product.value, selectedColor.value, selectedSize.value, quantity.value)
  alert(`Added ${quantity.value}x ${product.value?.name} (${selectedColor.value} - ${selectedSize.value}) to cart!`)
}

onMounted(() => {
  fetchProduct()
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
      <div v-if="loading" class="text-center py-5">
        <div class="spinner-border text-primary" role="status">
          <span class="visually-hidden">Loading...</span>
        </div>
      </div>

      <div v-else-if="product" class="row g-5 mb-5 align-items-center">
        <!-- Product Image -->
        <div class="col-md-6">
          <div class="glass-card p-2 text-center h-100 d-flex align-items-center justify-content-center overflow-hidden" style="min-height: 500px;">
            <img :src="product.image" class="img-fluid rounded product-main-img" :alt="product.name" />
          </div>
        </div>

        <!-- Product Details -->
        <div class="col-md-6 text-white">
          <nav aria-label="breadcrumb">
            <ol class="breadcrumb">
              <li class="breadcrumb-item"><router-link to="/store" class="text-muted">Store</router-link></li>
              <li class="breadcrumb-item text-muted" aria-current="page">{{ product.category }}</li>
            </ol>
          </nav>
          
          <h1 class="display-5 fw-extrabold mb-2">{{ product.name }}</h1>
          <h3 class="text-success fw-bold mb-4">฿{{ product.price.toLocaleString() }}</h3>
          
          <p class="text-muted mb-4 lh-lg">{{ product.description }}</p>

          <hr class="border-secondary mb-4" />

          <!-- Color Selection -->
          <div class="mb-4">
            <h6 class="fw-bold mb-3">Color: <span class="text-primary">{{ selectedColor || 'Select a color' }}</span></h6>
            <div class="d-flex gap-2">
              <button 
                v-for="color in availableColors" 
                :key="color"
                class="btn btn-outline-light rounded-pill px-4"
                :class="{ 'active bg-primary text-white border-primary': selectedColor === color }"
                @click="handleColorSelect(color)"
              >
                {{ color }}
              </button>
            </div>
          </div>

          <!-- Size Selection -->
          <div class="mb-4" v-if="selectedColor">
            <h6 class="fw-bold mb-3">Size: <span class="text-primary">{{ selectedSize || 'Select a size' }}</span></h6>
            <div class="d-flex gap-2 flex-wrap">
              <button 
                v-for="size in availableSizes" 
                :key="size"
                class="btn rounded-circle size-btn d-flex align-items-center justify-content-center"
                :class="{ 
                  'bg-primary text-white border-primary': selectedSize === size,
                  'btn-outline-secondary text-muted': selectedSize !== size && product.variants?.find(v => v.color === selectedColor && v.size === size)?.stock === 0,
                  'btn-outline-light': selectedSize !== size && (product.variants?.find(v => v.color === selectedColor && v.size === size)?.stock || 0) > 0
                }"
                :disabled="product.variants?.find(v => v.color === selectedColor && v.size === size)?.stock === 0"
                @click="handleSizeSelect(size)"
              >
                {{ size }}
              </button>
            </div>
          </div>

          <!-- Stock & Add to Cart -->
          <div v-if="currentVariant" class="mt-4">
            <div class="d-flex align-items-center mb-3">
              <span class="badge" :class="currentVariant.stock > 0 ? 'bg-success' : 'bg-danger'">
                {{ currentVariant.stock > 0 ? 'In Stock' : 'Out of Stock' }}
              </span>
              <span class="ms-2 text-muted small" v-if="currentVariant.stock > 0">
                {{ currentVariant.stock }} items available
              </span>
            </div>

            <div class="d-flex gap-3 align-items-center" v-if="currentVariant.stock > 0">
              <div class="input-group" style="width: 120px;">
                <button class="btn btn-outline-secondary text-white" type="button" @click="quantity > 1 && quantity--">-</button>
                <input type="text" class="form-control text-center bg-dark text-white border-secondary" v-model="quantity" readonly>
                <button class="btn btn-outline-secondary text-white" type="button" @click="quantity < currentVariant.stock && quantity++">+</button>
              </div>
              
              <button class="btn btn-primary flex-grow-1 py-2 fw-bold" @click="addToCart">
                <i class="bi bi-cart-plus me-2"></i> Add to Cart
              </button>
            </div>
          </div>

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

.product-main-img {
  max-height: 500px;
  object-fit: contain;
  transition: transform 0.3s ease;
}

.product-main-img:hover {
  transform: scale(1.05);
}

.size-btn {
  width: 45px;
  height: 45px;
  font-weight: bold;
}
</style>
