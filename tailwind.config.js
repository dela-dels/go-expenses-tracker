module.exports = {
    purge: [],
    darkMode: false, // or 'media' or 'class'
    theme: {
        fontFamily: {
            sans: ['Work Sans', 'sans-serif']
        },
    },
    variants: {
        extend: {},
    },
    plugins: [
        require("@tailwindcss/forms")
    ],
}
