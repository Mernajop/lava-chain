module.exports = {
    parser: '@typescript-eslint/parser',
    plugins: ['@typescript-eslint'],
    extends: [
        'plugin:@typescript-eslint/recommended',
        'plugin:prettier/recommended',
    ],
    rules: {
        "@typescript-eslint/no-explicit-any": "off",
    },
    ignorePatterns: ["*.d.ts"]
};