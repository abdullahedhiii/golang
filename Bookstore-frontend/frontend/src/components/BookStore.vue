<template>
  <div class="bookstore-container">
    <header class="header">
      <h1>📚 Book Store</h1>
      <button @click="logout" class="logout-btn">Logout</button>
    </header>
    
    <div class="content">
      <div v-if="loading" class="loading">Loading books...</div>
      
      <div v-else-if="error" class="error-message">{{ error }}</div>
      
      <div v-else class="books-grid">
        <div v-for="book in books" :key="book.id" class="book-card">
          <div class="book-info">
            <h3>{{ book.title }}</h3>
          </div>
          <div class="book-actions">
            <button class="btn-secondary" @click="fetchBook(book.id)">View Details</button>
          </div>
        </div>
      </div>
      
      <div v-if="!loading && books.length === 0" class="no-books">
        <p>No books available at the moment.</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BookStore',
  data() {
    return {
      books: [],
      loading: false,
      error: ''
    }
  },
  mounted() {
    this.checkAuth()
    this.fetchBooks()
  },
  methods: {
    checkAuth() {
      const token = localStorage.getItem('authToken')
      if (!token) {
        this.$router.push('/login')
      }
    },
    async fetchBooks() {
      this.loading = true
      this.error = ''
      
      try {
        const response = await fetch('http://localhost:8080/books', {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('authToken')}`
          }
        })
        
        if (response.ok) {
        
          this.books = await response.json()
        } else {
          this.error = 'Failed to fetch books'
        }
      } catch (error) {
        this.error = 'Network error. Please try again.'
      } finally {
        this.loading = false
      }
    },
    logout() {
      localStorage.removeItem('authToken')
      this.$router.push('/login')
    },

    fetchBook(id) {
      this.$router.push(`/books/${id}`)
    }
  }
}
</script>

<style scoped>
.bookstore-container {
  min-height: 100vh;
  background: #f8f9fa;
}

.header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 1rem 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.header h1 {
  margin: 0;
  font-size: 1.5rem;
}

.logout-btn {
  background: rgba(255, 255, 255, 0.2);
  color: white;
  border: 1px solid rgba(255, 255, 255, 0.3);
  padding: 0.5rem 1rem;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s;
}

.logout-btn:hover {
  background: rgba(255, 255, 255, 0.3);
}

.content {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.loading {
  text-align: center;
  font-size: 1.2rem;
  color: #666;
  margin-top: 2rem;
}

.error-message {
  color: #e53e3e;
  text-align: center;
  padding: 1rem;
  background: #fed7d7;
  border-radius: 5px;
  margin-top: 1rem;
}

.books-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1.5rem;
  margin-top: 1rem;
}

.book-card {
  background: white;
  border-radius: 10px;
  padding: 1.5rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s, box-shadow 0.3s;
}

.book-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.book-info h3 {
  margin: 0 0 0.5rem 0;
  color: #333;
  font-size: 1.2rem;
}

.author {
  color: #666;
  font-style: italic;
  margin: 0.5rem 0;
}

.book-id {
  color: #999;
  font-size: 0.9rem;
  margin: 0.5rem 0 1rem 0;
}

.book-actions {
  display: flex;
  gap: 0.5rem;
  margin-top: 1rem;
}

.btn-primary {
  background: #667eea;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s;
  flex: 1;
}

.btn-primary:hover {
  background: #5a67d8;
}

.btn-secondary {
  background: #e2e8f0;
  color: #4a5568;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s;
  flex: 1;
}

.btn-secondary:hover {
  background: #cbd5e0;
}

.no-books {
  text-align: center;
  margin-top: 2rem;
  color: #666;
}
</style>