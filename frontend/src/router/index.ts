import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import Warehouse from '../views/Warehouse.vue'
import Location from '../views/Location.vue'
import Product from '../views/Product.vue'
import Inventory from '../views/Inventory.vue'
import Inbound from '../views/Inbound.vue'
import Outbound from '../views/Outbound.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', name: 'Dashboard', component: Dashboard },
    { path: '/warehouses', name: 'Warehouse', component: Warehouse },
    { path: '/locations', name: 'Location', component: Location },
    { path: '/products', name: 'Product', component: Product },
    { path: '/inventory', name: 'Inventory', component: Inventory },
    { path: '/inbounds', name: 'Inbound', component: Inbound },
    { path: '/outbounds', name: 'Outbound', component: Outbound },
  ]
})

export default router
