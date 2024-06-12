/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./web/**/*.{html,templ,css}",
    "./node_modules/flowbite/**/*.js",
    './internal/web/**/*.templ',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
  ],
}

