<template>
  <div class="book-details-container">
    <header class="header">
      <h2>📖 Book Details</h2>
      <div class="actions">
        <button @click="goBack" class="btn">⬅ Back</button>
        <button @click="logout" class="btn logout">Logout</button>
      </div>
    </header>

    <div class="content">
      <div v-if="loading" class="loading">Loading...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else-if="book" class="details-card">
        <p><strong>ID:</strong> {{ book.id }}</p>
        <p><strong>Title:</strong> {{ book.title }}</p>
        <p><strong>Author:</strong> {{ book.author }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import api from "../utils/axios";

export default {
  name: 'BookDetails',
  data() {
    return {
      book: null,
      loading: true,
      error: ''
    }
  },
  methods: {
    checkAuth() {
      const token = localStorage.getItem('authToken')
      if (!token) {
        this.$router.push('/login')
      }
    },
    logout() {
      localStorage.removeItem('authToken')
      this.$router.push('/login')
    },
    goBack() {
      this.$router.go(-1)
    }
  },
  async mounted() {
    this.checkAuth()
    const bookId = this.$route.params.id

    try {
      const response = await api.get(`/books/${bookId}`)

      this.book = await response.data
    } catch (err) {
      this.error = err.message
    } finally {
      this.loading = false
    }
  }
}
</script>

<style scoped>
.book-details-container {
  max-width: 700px;
  margin: 2rem auto;
  padding: 2rem;
  background: #f9fafb;
  border-radius: 10px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);
  font-family: 'Segoe UI', sans-serif;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.actions {
  display: flex;
  gap: 0.5rem;
}

.btn {
  background: #667eea;
  color: white;
  border: none;
  padding: 0.4rem 1rem;
  border-radius: 5px;
  cursor: pointer;
  font-weight: 500;
  transition: background 0.3s ease;
}

.btn:hover {
  background: #5a67d8;
}

.logout {
  background: #e53e3e;
}

.logout:hover {
  background: #c53030;
}

.loading,
.error {
  text-align: center;
  font-size: 1.1rem;
  margin-top: 2rem;
  color: #4a5568;
}

.details-card {
  background: white;
  padding: 1.5rem;
  border-radius: 10px;
  box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.02);
  line-height: 1.6;
}
</style>
