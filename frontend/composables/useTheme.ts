export function useTheme() {
  const isDark = useState<boolean>('isDark', () => false)

  const toggleTheme = () => {
    isDark.value = !isDark.value
    updateTheme()
  }

  const setTheme = (dark: boolean) => {
    isDark.value = dark
    updateTheme()
  }

  const updateTheme = () => {
    if (process.client) {
      if (isDark.value) {
        document.documentElement.classList.add('dark')
        localStorage.setItem('theme', 'dark')
      } else {
        document.documentElement.classList.remove('dark')
        localStorage.setItem('theme', 'light')
      }
    }
  }

  const loadTheme = () => {
    if (process.client) {
      const savedTheme = localStorage.getItem('theme')
      if (savedTheme === 'dark') {
        isDark.value = true
      } else if (savedTheme === 'light') {
        isDark.value = false
      } else {
        // Check system preference
        isDark.value = window.matchMedia('(prefers-color-scheme: dark)').matches
      }
      updateTheme()
    }
  }

  // Initialize theme on client side
  if (process.client) {
    loadTheme()
  }

  return {
    isDark,
    toggleTheme,
    setTheme,
  }
}
