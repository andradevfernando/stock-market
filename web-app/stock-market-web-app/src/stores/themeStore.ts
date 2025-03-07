import { defineStore } from 'pinia';

export const useThemeStore = defineStore('theme', {
  state: () => ({
    isDark: false,
  }),
  actions: {
    toggleTheme() {
      this.isDark = !this.isDark
      document.documentElement.classList.toggle('dark', this.isDark)
      localStorage.setItem('theme', this.isDark ? 'dark' : 'light')
    },
    initializeTheme() {
      const savedTheme = localStorage.getItem('theme')
      const systemPrefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
      this.isDark = savedTheme ? savedTheme === 'dark' : systemPrefersDark
      document.documentElement.classList.toggle('dark', this.isDark)
    }
  },
  getters: {
    themeIcon: (state) => state.isDark ? '🌙' : '🌞'
  }
})