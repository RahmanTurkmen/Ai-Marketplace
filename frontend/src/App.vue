<script setup>
import { ref, computed, onMounted } from 'vue'
import { useModelsStore } from './stores/models'

const store = useModelsStore()

const selectedSection = ref('models')
const selectedCategory = ref('')
const name = ref('')
const description = ref('')
const category = ref('')
const accuracy = ref(90)
const author = ref('')

const categories = computed(() => {
  return [...new Set(store.models.map(m => m.category).filter(Boolean))].sort()
})

const filteredModels = computed(() => {
  const base = store.filteredModels

  if (selectedSection.value === 'trending') {
    return [...base].sort((a, b) => (b.downloads || 0) - (a.downloads || 0)).slice(0, 6)
  }

  if (selectedSection.value === 'categories' && selectedCategory.value) {
    return base.filter(m => m.category === selectedCategory.value)
  }

  return base
})

const sectionTitle = computed(() => {
  switch (selectedSection.value) {
    case 'trending':
      return 'Trending Models'
    case 'categories':
      return selectedCategory.value ? `Category: ${selectedCategory.value}` : 'Categories'
    case 'publish':
      return 'Publish a Model'
    default:
      return 'All Models'
  }
})

onMounted(() => {
  store.fetchModels()
})

function selectSection(section) {
  selectedSection.value = section
  if (section !== 'categories') {
    selectedCategory.value = ''
  }
}

function selectCategory(categoryName) {
  selectedSection.value = 'categories'
  selectedCategory.value = categoryName
}

function publish() {
  if (!name.value) return

  store.publishModel({
    name: name.value,
    description: description.value,
    category: category.value,
    accuracy: Number(accuracy.value),
    author: author.value
  })

  name.value = ''
  description.value = ''
  category.value = ''
  accuracy.value = 90
  author.value = ''
}
</script>

<template>
  <div class="app">
    <div class="sidebar">
      <h2>AI Marketplace</h2>

      <button
        class="sidebar-item"
        :class="{ active: selectedSection === 'models' }"
        @click="selectSection('models')"
      >
        Models
      </button>

      <button
        class="sidebar-item"
        :class="{ active: selectedSection === 'trending' }"
        @click="selectSection('trending')"
      >
        Trending
      </button>

      <button
        class="sidebar-item"
        :class="{ active: selectedSection === 'categories' }"
        @click="selectSection('categories')"
      >
        Categories
      </button>

      <button
        class="sidebar-item"
        :class="{ active: selectedSection === 'publish' }"
        @click="selectSection('publish')"
      >
        Publish
      </button>
    </div>

    <div class="main">
      <div class="topbar">
        <h1>{{ sectionTitle }}</h1>

        <input
          v-if="selectedSection !== 'publish'"
          v-model="store.search"
          placeholder="Search models..."
          class="search"
        />
      </div>

      <div v-if="selectedSection === 'publish'" class="publish-card">
        <h3>Publish Model</h3>

        <div class="grid">
          <input v-model="name" placeholder="Model name" />
          <input v-model="category" placeholder="Category" />
          <input v-model="accuracy" type="number" placeholder="Accuracy" />
          <input v-model="author" placeholder="Author" />
        </div>

        <textarea v-model="description" placeholder="Description"></textarea>

        <button class="btn-primary" @click="publish">
          Publish Model
        </button>
      </div>

      <div v-else>
        <div v-if="selectedSection === 'categories'" class="category-bar">
          <button
            v-for="cat in categories"
            :key="cat"
            class="category-pill"
            :class="{ selected: selectedCategory === cat }"
            @click="selectCategory(cat)"
          >
            {{ cat }}
          </button>
        </div>

        <div class="cards">
          <template v-if="filteredModels.length">
            <div
              class="model-card"
              v-for="m in filteredModels"
              :key="m.id"
            >
              <div class="card-header">
                <h3>{{ m.name }}</h3>
                <span class="category">{{ m.category }}</span>
              </div>

              <p class="desc">{{ m.description }}</p>

              <div class="stats">
                <span>🎯 {{ m.accuracy }}%</span>
                <span>⬇ {{ m.downloads }}</span>
                <span>⭐ {{ m.rating.toFixed(1) }}</span>
              </div>

              <div class="author">by {{ m.author }}</div>

              <div class="actions">
                <button class="btn-primary" @click="store.downloadModel(m.id)">
                  Download
                </button>
                <button class="btn-secondary" @click="store.rateModel(m.id)">
                  ⭐ Rate
                </button>
                <button class="btn-danger" @click="store.deleteModel(m.id)">
                  Delete
                </button>
              </div>
            </div>
          </template>

          <div v-else class="empty-state">
            <p>Aucun modèle trouvé pour cette section.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.sidebar {
  display: flex;
  flex-direction: column;
}

.sidebar-item {
  background: transparent;
  border: 1px solid transparent;
  color: inherit;
  text-align: left;
  padding: 1rem;
  cursor: pointer;
}


.sidebar-item.active {
  background: color-mix(in srgb, var(--accent) 18%, transparent);
  border: 1px solid color-mix(in srgb, var(--accent) 35%, transparent);
  color: var(--text);
  font-weight: 700;
}

.category-bar {
  margin-bottom: 1rem;
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.category-pill {
  border: 1px solid color-mix(in srgb, var(--border) 70%, transparent);
  background: transparent;
  color: inherit;
  padding: 0.5rem 0.8rem;
  border-radius: 999px;
  cursor: pointer;
}

.category-pill.selected {
  background: color-mix(in srgb, var(--accent) 18%, transparent);
  border-color: color-mix(in srgb, var(--accent) 35%, transparent);
}

.empty-state {
  padding: 2rem;
  text-align: center;
  color: color-mix(in srgb, var(--muted) 70%, transparent);
}

</style>
