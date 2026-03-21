<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getProductById, editProduct } from '@/services/productService'
import type { Product } from '@/types/product'

// router
const route = useRoute()
const router = useRouter()
const id = route.params.id as string

// state
const product = ref<Product>({
  id: '',
  name: '',
  price: 0,
  stock: 0
})

// loading
const loading = ref(false)

// ✅ โหลดข้อมูล
const fetchProduct = async () => {
  try {
    loading.value = true
    const data = await getProductById(id)
    product.value = data
  } catch (error) {
    console.error(error)
    alert('โหลดข้อมูลไม่สำเร็จ')
  } finally {
    loading.value = false
  }
}

// ✅ กด Save
const handleSubmit = async () => {
  try {
    const { id, ...payload } = product.value

    await editProduct(id, payload)

    alert('อัพเดตสำเร็จ')
    router.push('/products')
  } catch (error) {
    console.error(error)
    alert('อัพเดตไม่สำเร็จ')
  }
}

// load ตอนเข้า
onMounted(() => {
  fetchProduct()
})
</script>

<template>
  <div class="container mt-5">
    <div class="card shadow p-4">

      <h3 class="mb-4">Edit Product</h3>

      <!-- loading -->
      <div v-if="loading" class="text-center">
        Loading...
      </div>

      <div v-else>
        <!-- Name -->
        <div class="mb-3">
          <label class="form-label">Product Name</label>
          <input
            v-model="product.name"
            type="text"
            class="form-control"
          />
        </div>

        <!-- Price -->
        <div class="mb-3">
          <label class="form-label">Price</label>
          <input
            v-model="product.price"
            type="number"
            class="form-control"
          />
        </div>

        <!-- Stock -->
        <div class="mb-3">
          <label class="form-label">Stock</label>
          <input
            v-model="product.stock"
            type="number"
            class="form-control"
          />
        </div>

        <!-- Buttons -->
        <div class="d-flex gap-2">
          <button
            class="btn btn-success"
            @click="handleSubmit"
          >
            Save
          </button>

          <router-link
            to="/products"
            class="btn btn-secondary"
          >
            Cancel
          </router-link>
        </div>
      </div>

    </div>
  </div>
</template>