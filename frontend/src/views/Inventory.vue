<template>
  <div class="page">
    <div class="header">
      <h2>库存查询</h2>
    </div>

    <div class="filters">
      <select v-model="filters.warehouse_id" @change="loadData">
        <option value="">全部仓库</option>
        <option v-for="w in warehouses" :key="w.id" :value="w.id">{{ w.name }}</option>
      </select>
      <select v-model="filters.product_id" @change="loadData">
        <option value="">全部商品</option>
        <option v-for="p in products" :key="p.id" :value="p.id">{{ p.name }}</option>
      </select>
    </div>

    <table class="table">
      <thead>
        <tr>
          <th>仓库</th>
          <th>库位</th>
          <th>商品</th>
          <th>数量</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in list" :key="item.id">
          <td>{{ item.warehouse?.name }}</td>
          <td>{{ item.storage_location?.code }}</td>
          <td>{{ item.product?.name }}</td>
          <td>{{ item.quantity }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '../api'

const list = ref<any[]>([])
const warehouses = ref<any[]>([])
const products = ref<any[]>([])
const filters = ref({
  warehouse_id: '',
  product_id: ''
})

const loadData = async () => {
  const params: any = {}
  if (filters.value.warehouse_id) params.warehouse_id = filters.value.warehouse_id
  if (filters.value.product_id) params.product_id = filters.value.product_id

  const [invRes, whRes, prodRes]: any = await Promise.all([
    api.inventory.list(params),
    api.warehouses.list(),
    api.products.list()
  ])
  list.value = invRes.data
  warehouses.value = whRes.data
  products.value = prodRes.data
}

onMounted(loadData)
</script>

<style scoped>
.page { max-width: 1200px; }
.header { margin-bottom: 20px; }
.header h2 { color: #2c3e50; }
.filters { display: flex; gap: 10px; margin-bottom: 20px; }
.filters select { padding: 8px 12px; border: 1px solid #ddd; border-radius: 4px; }
.table { width: 100%; background: white; border-collapse: collapse; border-radius: 8px; overflow: hidden; box-shadow: 0 2px 8px rgba(0,0,0,0.1); }
.table th, .table td { padding: 12px 16px; text-align: left; border-bottom: 1px solid #eee; }
.table th { background: #f8f9fa; color: #7f8c8d; font-weight: 500; }
</style>
