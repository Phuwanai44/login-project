<script setup lang="ts">

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

const props = defineProps<{
  product: Product
}>()
</script>

<template>
  <div class="glass-card h-100 d-flex flex-column product-card">
    <div class="product-image-wrapper">
      <img 
        :src="product.image || 'https://via.placeholder.com/300x400?text=No+Image'" 
        class="card-img-top product-image" 
        :alt="product.name" 
      />
      <div v-if="product.category" class="category-badge">
        <span class="badge glass-badge">{{ product.category }}</span>
      </div>
    </div>
    
    <div class="p-4 d-flex flex-column flex-grow-1">
      <h5 class="fw-bold mb-2 text-white">{{ product.name }}</h5>
      <p class="text-success fs-5 fw-bold mb-3">฿{{ product.price.toLocaleString() }}</p>
      
      <div class="mt-auto">
        <router-link :to="`/store/${product.id}`" class="btn glass-btn w-100">
          View Details
        </router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
.product-card {
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  overflow: hidden;
}

.product-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
}

.product-image-wrapper {
  position: relative;
  overflow: hidden;
  aspect-ratio: 3/4;
}

.product-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s ease;
}

.product-card:hover .product-image {
  transform: scale(1.05);
}

.category-badge {
  position: absolute;
  top: 10px;
  right: 10px;
  z-index: 2;
}
</style>
