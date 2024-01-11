import type { Config } from 'tailwindcss'

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      colors: {
        limelight: '#9af501',
      },
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-conic': 'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
      },
      fontFamily: {
        'ropa-sans': ['Ropa Sans', 'sans-serif'],
        'bree-serif': ['Bree Serif', 'serif'],
        'suez-one': ['Suez One', 'serif'],
        poppins: ['Poppins', 'sans-serif'],
        merienda: ['Merienda', 'cursive'],
      },
    },
  },
  plugins: [],
}
export default config
