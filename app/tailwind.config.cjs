const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/**/*.{html,js,svelte,ts}',
  ],
  theme: {
    colors: {
      transparent: 'transparent',
      current: 'currentColor',
      black: colors.black,
      white: colors.white,
      gray: colors.gray,
      emerald: colors.emerald,
      indigo: colors.indigo,
      yellow: colors.yellow,
      slate: colors.slate,
      mastodon: '#595AFF',
      'mastodon-disabled': '#a2a2fa',
      'mastodon-hover': '#3839c9',
      dark1: '#191C22',
      dark2: '#282C37',
      dark3: '#313643',
      bglight: '#EFF3F5',
    },
  },
  plugins: [
    require("tailwindcss-debug-screens")
  ],
}
