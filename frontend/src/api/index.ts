import axios from 'axios'

const api = axios.create({
  baseURL: '/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

api.interceptors.response.use(
  response => response.data,
  error => {
    console.error('API Error:', error)
    return Promise.reject(error)
  }
)

export default {
  warehouses: {
    list: () => api.get('/warehouses'),
    get: (id: number) => api.get(`/warehouses/${id}`),
    create: (data: any) => api.post('/warehouses', data),
    update: (id: number, data: any) => api.put(`/warehouses/${id}`, data),
    delete: (id: number) => api.delete(`/warehouses/${id}`)
  },
  locations: {
    list: (params?: any) => api.get('/locations', { params }),
    get: (id: number) => api.get(`/locations/${id}`),
    create: (data: any) => api.post('/locations', data),
    update: (id: number, data: any) => api.put(`/locations/${id}`, data),
    delete: (id: number) => api.delete(`/locations/${id}`)
  },
  products: {
    list: (params?: any) => api.get('/products', { params }),
    get: (id: number) => api.get(`/products/${id}`),
    create: (data: any) => api.post('/products', data),
    update: (id: number, data: any) => api.put(`/products/${id}`, data),
    delete: (id: number) => api.delete(`/products/${id}`)
  },
  inventory: {
    list: (params?: any) => api.get('/inventory', { params }),
    summary: () => api.get('/inventory/summary')
  },
  inbounds: {
    list: (params?: any) => api.get('/inbounds', { params }),
    create: (data: any) => api.post('/inbounds', data)
  },
  outbounds: {
    list: (params?: any) => api.get('/outbounds', { params }),
    create: (data: any) => api.post('/outbounds', data)
  }
}
