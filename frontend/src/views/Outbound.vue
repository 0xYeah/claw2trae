<template>
  <div class="page">
    <div class="header">
      <h2>出库管理</h2>
      <button @click="showModal = true" class="btn-primary">新增出库</button>
    </div>

    <table class="table">
      <thead>
        <tr>
          <th>出库单号</th>
          <th>商品</th>
          <th>仓库</th>
          <th>库位</th>
          <th>数量</th>
          <th>目的地</th>
          <th>操作人</th>
          <th>时间</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in list" :key="item.id">
          <td>{{ item.outbound_no }}</td>
          <td>{{ item.product?.name }}</td>
          <td>{{ item.warehouse?.name }}</td>
          <td>{{ item.storage_location?.code }}</td>
          <td>{{ item.quantity }}</td>
          <td>{{ item.destination_info }}</td>
          <td>{{ item.operator }}</td>
          <td>{{ new Date(item.created_at).toLocaleString() }}</td>
        </tr>
      </tbody>
    </table>

    <div v-if="showModal" class="modal">
      <div class="modal-content">
        <h3>新增出库</h3>
        <form @submit.prevent="submitForm">
          <div class="form-group">
            <label>仓库</label>
            <select v-model="form.warehouse_id" required @change="loadLocations">
              <option value="">请选择</option>
              <option v-for="w in warehouses" :key="w.id" :value="w.id">{{ w.name }}</option>
            </select>
          </div>
          <div class="form-group">
            <label>库位</label>
            <select v-model="form.storage_location_id" required>
              <option value="">请选择</option>
              <option v-for="l in locations" :key="l.id" :value="l.id">{{ l.code }}</option>
            </select>
          </div>
          <div class="form-group">
            <label>商品</label>
            <select v-model="form.product_id" required>
              <option value="">请选择</option>
              <option v-for="p in products" :key="p.id" :value="p.id">{{ p.name }}</option>
            </select>
          </div>
          <div class="form-group">
            <label>数量</label>
            <input v-model.number="form.quantity" type="number" step="0.01" required />
          </div>
          <div class="form-group">
            <label>目的地</label>
            <input v-model="form.destination_info" />
          </div>
          <div class="form-group">
            <label>操作人</label>
            <input v-model="form.operator" />
          </div>
          <div class="form-actions">
            <button type="button" @click="closeModal" class="btn-secondary">取消</button>
            <button type="submit" class="btn-primary">提交</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '../api'

const list = ref<any[]>([])
const warehouses = ref<any[]>([])
const products = ref<any[]>([])
const locations = ref<any[]>([])
const showModal = ref(false)
const form = ref({
  warehouse_id: '' as any,
  storage_location_id: '' as any,
  product_id: '' as any,
  quantity: 0,
  destination_info: '',
  operator: ''
})

const loadData = async () => {
  const res: any = await api.outbounds.list()
  list.value = res.data
}

const loadRelated = async () => {
  const [whRes, prodRes]: any = await Promise.all([
    api.warehouses.list(),
    api.products.list()
  ])
  warehouses.value = whRes.data
  products.value = prodRes.data
}

const loadLocations = async () => {
  if (form.value.warehouse_id) {
    const res: any = await api.locations.list({ warehouse_id: form.value.warehouse_id })
    locations.value = res.data
  } else {
    locations.value = []
  }
}

const closeModal = () => {
  showModal.value = false
  form.value = { warehouse_id: '', storage_location_id: '', product_id: '', quantity: 0, destination_info: '', operator: '' }
  locations.value = []
}

const submitForm = async () => {
  await api.outbounds.create(form.value)
  closeModal()
  loadData()
}

onMounted(() => {
  loadData()
  loadRelated()
})
</script>

<style scoped>
.page { max-width: 1400px; }
.header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px; }
.header h2 { color: #2c3e50; }
.table { width: 100%; background: white; border-collapse: collapse; border-radius: 8px; overflow: hidden; box-shadow: 0 2px 8px rgba(0,0,0,0.1); }
.table th, .table td { padding: 12px 16px; text-align: left; border-bottom: 1px solid #eee; font-size: 14px; }
.table th { background: #f8f9fa; color: #7f8c8d; font-weight: 500; }
.btn-primary { background: #42b983; color: white; border: none; padding: 10px 20px; border-radius: 4px; cursor: pointer; }
.btn-secondary { background: #95a5a6; color: white; border: none; padding: 10px 20px; border-radius: 4px; cursor: pointer; }
.modal { position: fixed; top: 0; left: 0; right: 0; bottom: 0; background: rgba(0,0,0,0.5); display: flex; align-items: center; justify-content: center; }
.modal-content { background: white; padding: 30px; border-radius: 8px; width: 400px; }
.modal-content h3 { margin-bottom: 20px; color: #2c3e50; }
.form-group { margin-bottom: 15px; }
.form-group label { display: block; margin-bottom: 5px; color: #7f8c8d; }
.form-group input, .form-group select { width: 100%; padding: 8px 12px; border: 1px solid #ddd; border-radius: 4px; }
.form-actions { display: flex; gap: 10px; justify-content: flex-end; margin-top: 20px; }
</style>
