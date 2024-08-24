/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./public/**/*.{html,js}",
    "./public/*.html",  // この行を追加
    "./*.html"  // この行も追加（index.htmlがプロジェクトのルートにある場合）
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}