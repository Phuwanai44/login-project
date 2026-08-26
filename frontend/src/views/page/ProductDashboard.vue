<script setup lang="ts">
import { useRouter } from 'vue-router'
import { ref, onMounted, computed } from "vue"
import { getProducts, deleteProduct } from "@/services/productService"
import type { Product, ProductResponse } from "@/types/product"

const products = ref<Product[]>([])
const loading = ref(true)
const router = useRouter()

const currentPage = ref(1)
const totalPages = ref(1)
const search = ref("")

const loadProducts = async () => {
    loading.value = true
    try {
        const res: ProductResponse = await getProducts(
            currentPage.value,
            search.value
        )
        products.value = res.data
        totalPages.value = res.totalPages
    } catch (err) {
        console.error(err)
    } finally {
        loading.value = false
    }
}

const handleSearch = () => {
    currentPage.value = 1
    loadProducts()
}

const handleDelete = async (id: string) => {
    const confirmDelete = confirm("Are you sure you want to delete this product?")
    if (!confirmDelete) return

    try {
        await deleteProduct(id)
        products.value = products.value.filter(p => p.id !== id)
    } catch (err) {
        console.error(err)
    }
}

const handleEdit = (id: string) => {
    router.push(`/products/edit/${id}`)
}

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

const firstPage = computed(() => pagesToShow.value[0] ?? 0)
const lastPage = computed(() => pagesToShow.value[pagesToShow.value.length - 1] ?? 0)

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
    <div class="products-content">
        <div class="d-flex justify-content-between align-items-center mb-4">
            <div>
                <h2 class="fw-bold mb-1 text-gradient">Products Inventory</h2>
                <p class="text-muted small">Monitor storage stock, pricing, and distribution entries.</p>
            </div>
        </div>

        <!-- Search & Control -->
        <div class="d-flex justify-content-between align-items-center gap-3 mb-4">
            <div class="position-relative search-box-wrapper">
                <i class="bi bi-search search-icon"></i>
                <input type="text" class="form-control glass-input search-input" placeholder="Search products..." v-model="search" @input="handleSearch" />
            </div>

            <router-link to="/add-products" class="btn glass-btn d-flex align-items-center gap-2">
                <i class="bi bi-plus-lg"></i>
                <span>Add Product</span>
            </router-link>
        </div>

        <!-- Table -->
        <div class="glass-card mb-4">
            <div class="card-header-glass px-4 py-3">
                <h5 class="fw-bold mb-0">Product Listings</h5>
            </div>

            <div class="p-3">
                <div v-if="loading" class="text-center py-5">
                    <div class="spinner-border text-primary spinner-border-sm me-2" role="status"></div>
                    <span class="text-muted small">Loading product catalogue...</span>
                </div>

                <div v-else class="table-responsive glass-table-container">
                    <table class="glass-table">
                        <thead>
                            <tr>
                                <th>#</th>
                                <th>Name</th>
                                <th>Price</th>
                                <th>Stock</th>
                                <th class="text-end">Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-if="products.length === 0">
                                <td colspan="5" class="text-center text-muted small py-4">
                                    No products found in inventory
                                </td>
                            </tr>
                            <tr v-for="(product, index) in products" :key="product.id">
                                <td>{{ index + 1 }}</td>
                                <td class="fw-semibold">{{ product.name }}</td>
                                <td>
                                    <span class="price-tag">${{ product.price }}</span>
                                </td>
                                <td>
                                    <span class="badge glass-badge" :class="product.stock > 0 ? 'glass-badge-success' : 'glass-badge-admin'">
                                        {{ product.stock > 0 ? `${product.stock} In Stock` : 'Out of Stock' }}
                                    </span>
                                </td>
                                <td class="text-end">
                                    <button class="btn glass-btn-warning btn-sm me-2 py-1 px-3" @click="handleEdit(product.id)">
                                        <i class="bi bi-pencil-fill me-1"></i>Edit
                                    </button>
                                    <button class="btn glass-btn-danger btn-sm py-1 px-3" @click="handleDelete(product.id)">
                                        <i class="bi bi-trash-fill me-1"></i>Delete
                                    </button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>

        <!-- Pagination -->
        <nav v-if="totalPages > 1" class="d-flex justify-content-center">
            <ul class="pagination glass-pagination">
                <!-- Prev -->
                <li class="page-item" :class="{ disabled: currentPage === 1 }">
                    <button class="page-link" @click="prevPage">
                        <i class="bi bi-chevron-left"></i>
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
                <li v-for="page in pagesToShow" :key="page" class="page-item" :class="{ active: page === currentPage }">
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
                <li class="page-item" :class="{ disabled: currentPage === totalPages }">
                    <button class="page-link" @click="nextPage">
                        <i class="bi bi-chevron-right"></i>
                    </button>
                </li>
            </ul>
        </nav>
    </div>
</template>

<style scoped>
.products-content {
    color: var(--text-primary);
}

.search-box-wrapper {
    width: 100%;
    max-width: 320px;
}

.search-icon {
    position: absolute;
    left: 14px;
    top: 50%;
    transform: translateY(-50%);
    color: var(--text-muted);
    z-index: 5;
}

.search-input {
    padding-left: 40px !important;
    width: 100%;
}

.card-header-glass {
    border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.price-tag {
    font-family: 'Outfit', sans-serif;
    font-weight: 700;
    color: var(--text-primary);
}
</style>