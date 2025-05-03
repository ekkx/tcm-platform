import { heroui } from "@heroui/react";

import type { Config } from "tailwindcss";

export default {
  content: [
    // ...
    "./app/**/{**,.client,.server}/**/*.{js,jsx,ts,tsx}",
    "./node_modules/@heroui/theme/dist/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    fontFamily: {
      primary: ['"Zen Maru Gothic"', "sans-serif"],
      secondary: ['"Kosugi Maru"', "sans-serif"],
    },
    extend: {
      colors: {
        border:
          "hsl(var(--nextui-divider) / var(--nextui-divider-opacity, var(--tw-bg-opacity)))",
      },
    },
  },
  darkMode: "class",
  plugins: [heroui()],
} satisfies Config;
