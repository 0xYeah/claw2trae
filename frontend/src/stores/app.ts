import { defineStore } from 'pinia'
import api from '../api'

interface Summary {
  total_warehouses: number
  total_products: number
  total_quantity: number
  total_inbounds: number
  total_outbounds: number
}

export const useAppStore = defineStore('app', {
  state: () => ({
    summary: null as Summary | null,
    loading: false
  }),
  actions: {
    async fetchSummary() {
      this.loading = true
      try {
        const res: any = await api.inventory.summary()
        this.summary = res.data
      } catch (error) {
        console.error('Failed to fetch summary:', error)
      } finally {
        this.loading = false
      }
    }
  }
})
