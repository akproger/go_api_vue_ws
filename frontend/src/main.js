import './assets/main.css'
import { createApp } from 'vue'
import axios from 'axios'
import App from './App.vue'

// Настройка axios для использования прокси
axios.defaults.baseURL = '/api'

// Создаем приложение Vue
const app = createApp(App)

// Добавляем axios в глобальные переменные, чтобы можно было обращаться из компонентов
app.config.globalProperties.$axios = axios

// Монтируем приложение
app.mount('#app')
