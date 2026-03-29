<template>
  <div class="app-container">
    <aside class="sidebar">
      <div class="logo">WMS</div>
      <nav class="nav">
        <router-link to="/">Dashboard</router-link>
        <router-link to="/warehouses">仓库管理</router-link>
        <router-link to="/locations">库位管理</router-link>
        <router-link to="/products">商品管理</router-link>
        <router-link to="/inventory">库存查询</router-link>
        <router-link to="/inbounds">入库管理</router-link>
        <router-link to="/outbounds">出库管理</router-link>
      </nav>
    </aside>
    <main class="main-content">
      <header class="header">
        <h1>仓储管理系统</h1>
      </header>
      <div class="content">
        <router-view />
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.app-container {
  display: flex;
  min-height: 100vh;
}

.sidebar {
  width: 200px;
  background-color: #2c3e50;
  color: white;
  padding: 20px 0;
}

.logo {
  font-size: 24px;
  font-weight: bold;
  text-align: center;
  padding: 10px 0 30px;
  color: #42b983;
}

.nav {
  display: flex;
  flex-direction: column;
}

.nav a {
  color: #ecf0f1;
  text-decoration: none;
  padding: 12px 20px;
  transition: background-color 0.3s;
}

.nav a:hover,
.nav a.router-link-active {
  background-color: #34495e;
  color: #42b983;
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.header {
  background: white;
  padding: 20px 30px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.header h1 {
  font-size: 20px;
  color: #2c3e50;
}

.content {
  padding: 20px 30px;
  flex: 1;
}
</style>
