<template>
  <div class="dashboard">
    <h2>数据概览</h2>
    <div class="cards">
      <div class="card">
        <div class="card-title">仓库数</div>
        <div class="card-value">{{ summary?.total_warehouses || 0 }}</div>
      </div>
      <div class="card">
        <div class="card-title">商品数</div>
        <div class="card-value">{{ summary?.total_products || 0 }}</div>
      </div>
      <div class="card">
        <div class="card-title">库存总量</div>
        <div class="card-value">{{ summary?.total_quantity || 0 }}</div>
      </div>
      <div class="card">
        <div class="card-title">入库记录</div>
        <div class="card-value">{{ summary?.total_inbounds || 0 }}</div>
      </div>
      <div class="card">
        <div class="card-title">出库记录</div>
        <div class="card-value">{{ summary?.total_outbounds || 0 }}</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useAppStore } from '../stores/app'

const store = useAppStore()

onMounted(() => {
  store.fetchSummary()
})
</script>

<style scoped>
.dashboard h2 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.card-title {
  color: #7f8c8d;
  font-size: 14px;
  margin-bottom: 10px;
}

.card-value {
  color: #2c3e50;
  font-size: 32px;
  font-weight: bold;
}
</style>
