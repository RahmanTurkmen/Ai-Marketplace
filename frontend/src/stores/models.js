import { defineStore } from 'pinia'
import api from '../services/api'

export const useModelsStore = defineStore('models', {

  state: () => ({
    models: [],
    search: ''
  }),

  getters: {

    filteredModels(state) {
      return state.models.filter(m =>
        m.name.toLowerCase().includes(state.search.toLowerCase())
      )
    }
  },

  actions: {

    async fetchModels() {
      const res = await api.get('/models')
      this.models = Array.isArray(res.data) ? res.data : []
    },

    async publishModel(model) {
      await api.post('/models', model)
      await this.fetchModels()
    },

    async deleteModel(id) {
      await api.delete(`/models/${id}`)
      await this.fetchModels()
    },

    async downloadModel(id) {
      await api.patch(`/models/${id}/download`)
      await this.fetchModels()
    },

    async rateModel(id) {
      await api.patch(`/models/${id}/rate`)
      await this.fetchModels()
    }
  }
})