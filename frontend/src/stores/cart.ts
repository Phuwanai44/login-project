import { defineStore } from 'pinia'
import { ref, computed, watch } from 'vue'

export interface CartItem {
  id: string
  productId: string
  name: string
  price: number
  color: string
  size: string
  quantity: number
  image: string
}

export const useCartStore = defineStore('cart', () => {
  const items = ref<CartItem[]>([])

  // Load from local storage
  const savedCart = localStorage.getItem('cart')
  if (savedCart) {
    items.value = JSON.parse(savedCart)
  }

  // Save to local storage on change
  watch(items, (newItems) => {
    localStorage.setItem('cart', JSON.stringify(newItems))
  }, { deep: true })

  const totalItems = computed(() => {
    return items.value.reduce((total, item) => total + item.quantity, 0)
  })

  const totalPrice = computed(() => {
    return items.value.reduce((total, item) => total + (item.price * item.quantity), 0)
  })

  const addToCart = (product: any, color: string, size: string, quantity: number) => {
    const existingItem = items.value.find(
      (item) => item.productId === product.id && item.color === color && item.size === size
    )

    if (existingItem) {
      existingItem.quantity += quantity
    } else {
      items.value.push({
        id: `${product.id}-${color}-${size}`,
        productId: product.id,
        name: product.name,
        price: product.price,
        color,
        size,
        quantity,
        image: product.image
      })
    }
  }

  const updateQuantity = (itemId: string, quantity: number) => {
    const item = items.value.find(i => i.id === itemId)
    if (item && quantity > 0) {
      item.quantity = quantity
    }
  }

  const removeItem = (itemId: string) => {
    items.value = items.value.filter(i => i.id !== itemId)
  }

  const clearCart = () => {
    items.value = []
  }

  return {
    items,
    totalItems,
    totalPrice,
    addToCart,
    updateQuantity,
    removeItem,
    clearCart
  }
})
