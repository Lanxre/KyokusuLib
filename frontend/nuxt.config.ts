import tailwindcss from '@tailwindcss/vite'

export default defineNuxtConfig({
	app: {
		head: {
			title: "KyokusuLib",
			charset: "utf-8",
			viewport: "width=device-width, initial-scale=1",
			link: [
				{ rel: 'icon', type: 'image/svg+xml', href: '/images/KLibIcon.svg' }
			]
		},
	},
	devtools: { enabled: true },

	modules: ["@pinia/nuxt", "@vueuse/nuxt", "@nuxtjs/color-mode", "@nuxt/icon"],

	routeRules: {
		"/api/**": { proxy: "http://localhost:8080/api/**" },
	},

	colorMode: {
		preference: 'dark',
		fallback: 'dark',
		classSuffix: '',
		storageKey: 'theme'
	},

	icon: {
		clientBundle: {
            scan: true,
            includeCustomCollections: true,
        },
		serverBundle: 'local',
	},

	vite: {
		plugins: [
			tailwindcss(),
		],
	},

	css: ["~/assets/css/global.css"],

	postcss: {
		plugins: {
			autoprefixer: {},
		},
	},

	sourcemap: {
		client: "hidden",
	},

	runtimeConfig: {
		public: {
			apiBase: process.env.NUXT_PUBLIC_API_BASE || "http://localhost:8080",
		},
	},

	compatibilityDate: "latest",
});