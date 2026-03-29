<template>
  <div class="page">
    <div class="header">
      <h2>仓库管理</h2>
      <button @click="showModal = true" class="btn-primary">新增仓库</button>
    </div>

    <table class="table">
      <thead>
        <tr>
          <th>编号</th>
          <th>名称</th>
          <th>地址</th>
          <th>面积(m²)</th>
          <th>负责人</th>
          <th>状态</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in list" :key="item.id">
          <td>{{ item.code }}</td>
          <td>{{ item.name }}</td>
          <td>{{ item.address }}</td>
          <td>{{ item.area }}</td>
          <td>{{ item.manager }}</td>
          <td>{{ item.status === 1 ? '启用' : '禁用' }}</td>
          <td>
            <button @click="editItem(item)" class="btn-text">编辑</button>
            <button @click="deleteItem(item.id)" class="btn-danger">删除</button>
          </td>
        </tr>
      </tbody>
    </table>

    <div v-if="showModal" class="modal">
      <div class="modal-content">
        <h3>{{ editingId ? '编辑仓库' : '新增仓库' }}</h3>
        <form @submit.prevent="submitForm">
          <div class="form-group">
            <label>编号</label>
            <input v-model="form.code" required />
          </div>
          <div class="form-group">
            <label>名称</label>
            <input v-model="form.name" required />
          </div>
          <div class="form-group">
            <label>地址</label>
            <input v-model="form.address" />
          </div>
          <div class="form-group">
            <label>面积</label>
            <input v-model.number="form.area" type="number" step="0.01" />
          </div>
          <div class="form-group">
            <label>负责人</label>
            <input v-model="form.manager" />
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
const showModal = ref(false)
const editingId = ref<number | null>(null)
const form = ref({
  code: '',
  name: '',
  address: '',
  area: 0,
  manager: '',
  status: 1
})

const loadData = async () => {
  const res: any = await api.warehouses.list()
  list.value = res.data
}

const editItem = (item: any) => {
  editingId.value = item.id
  form.value = { ...item }
  showModal.value = true
}

const deleteItem = async (id: number) => {
  if (confirm('确定删除?')) {
    await api.warehouses.delete(id)
    loadData()
  }
}

const closeModal = () => {
  showModal.value = false
  editingId.value = null
  form.value = { code: '', name: '', address: '', area: 0, manager: '', status: 1 }
}

const submitForm = async () => {
  if (editingId.value) {
    await api.warehouses.update(editingId.value, form.value)
  } else {
    await api.warehouses.create(form.value)
  }
  closeModal()
  loadData()
}

onMounted(loadData)
</script>

<style scoped>
.page { max-width: 1200px; }
.header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px; }
.header h2 { color: #2c3e50; }
.table { width: 100%; background: white; border-collapse: collapse; border-radius: 8px; overflow: hidden; box-shadow: 0 2px 8px rgba(0,0,0,0.1); }
.table th, .table td { padding: 12px 16px; text-align: left; border-bottom: 1px solid #eee; }
.table th { background: #f8f9fa; color: #7f8c8d; font-weight: 500; }
.btn-primary { background: #42b983; color: white; border: none; padding: 10px 20px; border-radius: 4px; cursor: pointer; }
.btn-secondary { background: #95a5a6; color: white; border: none; padding: 10px 20px; border-radius: 4px; cursor: pointer; }
.btn-text { background: none; border: none; color: #3498db; cursor: pointer; margin-right: 10px; }
.btn-danger { background: none; border: none; color: #e74c3c; cursor: pointer; }
.modal { position: fixed; top: 0; left: 0; right: 0; bottom: 0; background: rgba(0,0,0,0.5); display: flex; align-items: center; justify-content: center; }
.modal-content { background: white; padding: 30px; border-radius: 8px; width: 400px; }
.modal-content h3 { margin-bottom: 20px; color: #2c3e50; }
.form-group { margin-bottom: 15px; }
.form-group label { display: block; margin-bottom: 5px; color: #7f8c8d; }
.form-group input { width: 100%; padding: 8px 12px; border: 1px solid #ddd; border-radius: 4px; }
.form-actions { display: flex; gap: 10px; justify-content: flex-end; margin-top: 20px; }
</style>
