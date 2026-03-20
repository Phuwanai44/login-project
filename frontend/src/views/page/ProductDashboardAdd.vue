<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"
import type { Product } from "@/types/product"
import { createProduct } from "@/services/productService"

const router = useRouter()


// state form
const product = ref<Product>({
  id: "",
  name: "",
  price: 0,
  stock: 0
})



const handleSubmit = async () => {
  try {
    const res = await createProduct(product.value)

    console.log(res)

    // redirect กลับ list
    router.push("/products")

  } catch (err) {
    console.error(err)
  }
}
</script>

<template>
  <div class="container mt-5">
    <div class="card shadow p-4">
      <h3 class="mb-4">Add Product</h3>

      <!-- Name -->
      <div class="mb-3">
        <label class="form-label">Product Name</label>
        <input
          v-model="product.name"
          type="text"
          class="form-control"
          placeholder="Enter product name"
        />
      </div>

      <!-- Price -->
      <div class="mb-3">
        <label class="form-label">Price</label>
        <input
          v-model="product.price"
          type="number"
          class="form-control"
          placeholder="Enter price"
        />
      </div>

      <!-- Stock -->
      <div class="mb-3">
        <label class="form-label">Stock</label>
        <input
          v-model="product.stock"
          type="number"
          class="form-control"
          placeholder="Enter stock"
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
</template>