/** @type {import('tailwindcss').Config} */
export const content = [
    "./internal/templates/**/*.templ",
    "./internal/templates/*.templ",
];
export const theme = {
    container: {
        center: true,
        padding: {
            DEFAULT: "1rem",
            mobile: "2rem",
            tablet: "4rem",
            desktop: "5rem",
        },
    },
    extend: {},
};
export const colors = {
      'blue': '#1fb6ff',
      'purple': '#7e5bef',
      'pink': '#ff49db',
      'orange': '#ff7849',
      'green': '#13ce66',
      'yellow': '#ffc82c',
      'gray-dark': '#273444',
      'gray': '#8492a6',
      'gray-light': '#d3dce6',
};
export const plugins = [require("@tailwindcss/forms"), require("@tailwindcss/typography")];
