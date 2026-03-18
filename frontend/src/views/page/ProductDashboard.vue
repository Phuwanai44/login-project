<script setup lang="ts">
import { ref, onMounted, computed } from "vue"
import { getProducts } from "@/services/productService"
import type { Product, ProductResponse } from "@/types/product"

const products = ref<Product[]>([]) // ✅ แก้ตรงนี้
const loading = ref(true)

const currentPage = ref(1)
const totalPages = ref(1)

const search = ref("")

const loadProducts = async () => {
    loading.value = true

    const res: ProductResponse = await getProducts(
        currentPage.value,
        search.value
    )

    products.value = res.data // ✅ ใช้ได้แล้ว
    totalPages.value = res.totalPages // ✅ ใช้ได้แล้ว

    loading.value = false
}

// 👉 search
const handleSearch = () => {
    currentPage.value = 1
    loadProducts()
}

// 👉 pagination logic (แสดง 5 หน้า)
const pagesToShow = computed(() => {
    const pages: number[] = []
    const maxVisible = 5

    let start = Math.max(1, currentPage.value - 2)
    const end = Math.min(totalPages.value, start + maxVisible - 1)

    if (end - start < maxVisible - 1) {
        start = Math.max(1, end - maxVisible + 1)
    }

    for (let i = start; i <= end; i++) {
        pages.push(i)
    }

    return pages
})

// ✅ แก้ error undefined
const firstPage = computed(() => pagesToShow.value[0] ?? 0)
const lastPage = computed(() => pagesToShow.value[pagesToShow.value.length - 1] ?? 0)

// 👉 control
const goToPage = (page: number) => {
    currentPage.value = page
    loadProducts()
}

const nextPage = () => {
    if (currentPage.value < totalPages.value) {
        currentPage.value++
        loadProducts()
    }
}

const prevPage = () => {
    if (currentPage.value > 1) {
        currentPage.value--
        loadProducts()
    }
}

onMounted(() => {
    loadProducts()
})
</script>

<template>
<div class="container mt-4">

    <h2 class="mb-4">Products</h2>

    <!-- Search -->
    <div class="d-flex justify-content-between mb-3">
        <div class="input-group" style="max-width:300px">
            <input 
                type="text" 
                class="form-control" 
                placeholder="Search product..." 
                v-model="search"
                @input="handleSearch"
            />
        </div>

        <button class="btn btn-primary">
            Add Product
        </button>
    </div>

    <!-- Table -->
    <div class="card shadow-sm">
        <div class="card-body">
            <table class="table table-hover">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Name</th>
                        <th>Price</th>
                        <th>Stock</th>
                        <th>Action</th>
                    </tr>
                </thead>

                <tbody>
                    <tr v-if="loading">
                        <td colspan="5" class="text-center">Loading...</td>
                    </tr>

                    <tr v-for="(product, index) in products" :key="product.id">
                        <td>{{ index + 1 }}</td>
                        <td>{{ product.name }}</td>
                        <td>{{ product.price }}</td>
                        <td>{{ product.stock }}</td>
                        <td>
                            <button class="btn btn-sm btn-warning me-2">Edit</button>
                            <button class="btn btn-sm btn-danger">Delete</button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>

    <!-- Pagination -->
    <nav class="mt-3">
        <ul class="pagination">

            <!-- Prev -->
            <li class="page-item">
                <button class="page-link" @click="prevPage" :disabled="currentPage === 1">
                    Previous
                </button>
            </li>

            <!-- First -->
            <li v-if="firstPage > 1" class="page-item">
                <button class="page-link" @click="goToPage(1)">1</button>
            </li>

            <!-- ... -->
            <li v-if="firstPage > 2" class="page-item disabled">
                <span class="page-link">...</span>
            </li>

            <!-- Pages -->
            <li v-for="page in pagesToShow" :key="page"
                class="page-item"
                :class="{ active: page === currentPage }">
                <button class="page-link" @click="goToPage(page)">
                    {{ page }}
                </button>
            </li>

            <!-- ... -->
            <li v-if="lastPage < totalPages - 1" class="page-item disabled">
                <span class="page-link">...</span>
            </li>

            <!-- Last -->
            <li v-if="lastPage < totalPages" class="page-item">
                <button class="page-link" @click="goToPage(totalPages)">
                    {{ totalPages }}
                </button>
            </li>

            <!-- Next -->
            <li class="page-item">
                <button class="page-link" @click="nextPage" :disabled="currentPage === totalPages">
                    Next
                </button>
            </li>

        </ul>
    </nav>

</div>
</template>