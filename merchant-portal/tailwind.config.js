module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: '#165DFF',
        secondary: '#36CFC9',
        success: '#52C41A',
        warning: '#FAAD14',
        danger: '#F5222D',
        neutral: {
          100: '#F5F7FA',
          200: '#E5E6EB',
          300: '#C9CDD4',
          400: '#86909C',
          500: '#4E5969',
          600: '#272E3B',
          700: '#1D2129',
        }
      },
      fontFamily: {
        inter: ['Inter', 'sans-serif'],
      },
    },
  },
  plugins: [],
}