<script setup lang="ts">
import { useReader } from "~/composables/api/novela/useReader";
import { useReaderSettings } from "~/composables/ui/useReaderSettings";
import { staticImage } from "@/utils/str";
import ReaderSettings from "~/components/app/novela/ReaderSettings.vue";
import ReaderTOC from "~/components/app/novela/ReaderTOC.vue";
import { useInfiniteScroll, useEventListener } from "@vueuse/core";
import { Tooltip } from "@kyokusu-ui/vue";

const route = useRoute();
const { chapters, isLoading, isAppending, fetchChapter } = useReader();
const { fontSize, fontFamily, lineWeight, isAutoScrollEnabled } = useReaderSettings();

const novelaId = computed(() => route.params.novelaId as string);
const chapterId = computed(() => route.params.chapterId as string);

await useAsyncData(`reader-${chapterId.value}`, () => fetchChapter(chapterId.value));

// Reset chapters when navigating via URL directly (not via auto-scroll)
watch(chapterId, (newId) => {
	const isCurrentlyLoaded = chapters.value.some(c => c.id === newId);
	if (!isCurrentlyLoaded) {
		fetchChapter(newId);
	}
});

const firstChapter = computed(() => chapters.value[0]);
const lastChapter = computed(() => chapters.value[chapters.value.length - 1]);

// Track which chapter is currently in view
const currentChapterId = ref(chapterId.value);
useEventListener(window, 'scroll', () => {
	const articles = document.querySelectorAll<HTMLElement>('[data-chapter-id]');
	const triggerLine = window.scrollY + window.innerHeight * 0.3;

	let lastId: string | null = null;
	for (const el of articles) {
		const top = el.getBoundingClientRect().top + window.scrollY;
		if (triggerLine >= top) {
			lastId = el.dataset.chapterId || null;
		}
	}

	if (lastId && lastId !== currentChapterId.value) {
		currentChapterId.value = lastId;
	}
}, { passive: true });

const currentChapter = computed(() =>
	chapters.value.find(c => c.id === currentChapterId.value) || chapters.value[0]
);

// Auto-scroll: load next chapter when nearing the bottom
useInfiniteScroll(
	window,
	() => {
		const last = lastChapter.value;
		if (last?.next_chapter_id && isAutoScrollEnabled.value) {
			fetchChapter(last.next_chapter_id, true);
		}
	},
	{ distance: 400 }
);

const lightboxImage = ref<string | null>(null);

function openLightbox(url: string) {
	lightboxImage.value = url;
}

function closeLightbox() {
	lightboxImage.value = null;
}

useEventListener(window, 'keydown', (e: KeyboardEvent) => {
	if (e.key === 'Escape' && lightboxImage.value) {
		closeLightbox();
	}
});

function onContentClick(e: MouseEvent) {
	const target = e.target as HTMLElement;
	if (target.tagName === 'IMG' && target.closest('.prose')) {
		const src = target.getAttribute('src');
		if (src) openLightbox(src);
	}
}

const isSettingsOpen = ref(false);
const isTocOpen = ref(false);
const backgroundColor = ref("bg-zinc-50 dark:bg-[#0f0f0f]");
const textColor = ref("text-zinc-900 dark:text-zinc-200");
</script>

<template>
	<div class="min-h-screen transition-colors duration-300" :class="backgroundColor">
		<header class="sticky z-30 top-20.5 md:top-22.5 bg-white/80 dark:bg-zinc-900/80 backdrop-blur-md border-b border-zinc-200 dark:border-zinc-800">
			<div class="max-w-5xl mx-auto px-4 h-16 grid grid-cols-[auto_1fr_auto] items-center gap-4">
				<div class="flex items-center justify-start mt-1">
					<Tooltip text="Вернуться к описанию" position="right">
						<NuxtLink :to="`/novela/${novelaId}`" class="rounded-xl transition-colors shrink-0">
							<Icon name="ph:arrow-left-bold" size="20" />
						</NuxtLink>
					</Tooltip>
				</div>

				<button
					@click="isTocOpen = true"
					class="flex flex-col items-center group overflow-hidden cursor-pointer min-w-0"
				>
					<span class="text-[10px] text-zinc-500 font-bold uppercase tracking-widest truncate max-w-full group-hover:text-yellow-500 transition-colors">
						{{ currentChapter?.novela_title }}
					</span>
					<div class="flex items-center gap-1.5 overflow-hidden max-w-full">
						<h1 class="text-sm font-black truncate group-hover:text-zinc-600 dark:group-hover:text-zinc-300 transition-colors">
							Том {{ currentChapter?.volume_number }} • Глава {{ currentChapter?.number }}: {{ currentChapter?.title }}
						</h1>
						<Icon name="ph:caret-down-bold" size="14" class="text-zinc-400 group-hover:text-yellow-500 transition-all group-hover:translate-y-0.5 shrink-0" />
					</div>
				</button>

				<div class="flex items-center justify-end">
				    <Tooltip text="Настройки" position="right">
						<button @click="isSettingsOpen = true" class="cursor-pointer rounded-xl transition-colors">
							<Icon name="ph:list-bold" size="20" />
						</button>
					</Tooltip>
				</div>
			</div>
		</header>

		<main class="max-w-3xl mx-auto px-4 py-12 min-h-[calc(100vh-128px)] flex flex-col gap-16">
			<div v-if="isLoading && !chapters.length" class="flex-1 flex items-center justify-center">
				<div class="w-12 h-12 border-4 border-zinc-300 border-t-yellow-500 rounded-full animate-spin"></div>
			</div>
			
			<template v-else-if="chapters.length">
				<div v-for="(ch, index) in chapters" :key="ch.id" :data-chapter-id="ch.id" class="flex flex-col gap-12">
					<!-- Chapter Separator -->
					<div v-if="index > 0" class="flex items-center gap-4 py-8">
						<div class="h-px grow bg-zinc-200 dark:bg-zinc-800"></div>
						<div class="px-4 py-2 bg-zinc-100 dark:bg-zinc-800 rounded-full border border-zinc-200 dark:border-zinc-700">
							<span class="text-xs font-black uppercase tracking-widest text-zinc-400">
								Глава {{ ch.number }}: {{ ch.title }}
							</span>
						</div>
						<div class="h-px grow bg-zinc-200 dark:bg-zinc-800"></div>
					</div>

					<article class="flex-1">
						<div 
							class="prose dark:prose-invert max-w-none leading-relaxed whitespace-pre-line cursor-pointer"
							:class="[textColor, fontFamily]"
							:style="{ fontSize: `${fontSize}px`, lineHeight: lineWeight }"
							v-html="ch.content"
							@click="onContentClick"
						></div>

						<div v-if="ch.images?.length" class="mt-12 space-y-8">
							<div v-for="img in ch.images" :key="img.id" class="flex flex-col items-center gap-2">
								<img :src="staticImage(img.image_url)" :alt="img.caption" class="w-full max-w-md h-full object-cover rounded-xl shadow-lg cursor-pointer" @click="openLightbox(staticImage(img.image_url))" />
								<p v-if="img.caption" class="text-sm text-zinc-500 italic">{{ img.caption }}</p>
							</div>
						</div>
					</article>
				</div>

				<div v-if="isAppending" class="py-12 flex justify-center">
					<div class="w-8 h-8 border-3 border-zinc-300 border-t-yellow-500 rounded-full animate-spin"></div>
				</div>
			</template>

			<div v-else class="flex-1 flex flex-col items-center justify-center text-center py-20">
				<Icon name="ph:warning-circle-bold" size="48" class="text-zinc-400 mb-4" />
				<h2 class="text-xl font-bold mb-2">Глава не найдена</h2>
				<p class="text-zinc-500 mb-6">Возможно, она была удалена или еще не одобрена.</p>
				<NuxtLink to="/" class="px-6 py-2 bg-yellow-500 text-white font-bold rounded-xl hover:bg-yellow-600 transition-colors">
					На главную
				</NuxtLink>
			</div>
		</main>

		<ReaderTOC 
			v-if="firstChapter"
			v-model="isTocOpen" 
			:novela-id="firstChapter.novela_id" 
			:current-chapter-id="currentChapterId"
			:novela-title="firstChapter.novela_title"
		/>
		<ReaderSettings v-model="isSettingsOpen" />

		<footer v-if="!isAutoScrollEnabled || !lastChapter?.next_chapter_id" class="bg-white dark:bg-zinc-900 border-t border-zinc-200 dark:border-zinc-800 py-8">
			<div class="max-w-5xl mx-auto px-4 flex flex-col sm:flex-row items-center justify-between gap-6">
				<div class="flex items-center gap-4 w-full sm:w-auto">
					<NuxtLink 
						v-if="firstChapter?.prev_chapter_id" 
						:to="`/novela/reader/${novelaId}/${firstChapter.prev_chapter_id}`"
						class="flex-1 sm:flex-none flex items-center justify-center gap-2 px-6 py-3 bg-zinc-100 dark:bg-zinc-800 hover:bg-zinc-200 dark:hover:bg-zinc-700 rounded-xl font-bold transition-all active:scale-95"
					>
						<Icon name="ph:caret-left-bold" size="20" />
						<span>Назад</span>
					</NuxtLink>
					<div v-else class="flex-1 sm:flex-none px-6 py-3 opacity-0 pointer-events-none">
						<span>Назад</span>
					</div>
				</div>

				<NuxtLink :to="`/novela/${novelaId}?tab=chapters`" class="text-zinc-500 hover:text-zinc-900 dark:hover:text-white transition-colors flex items-center gap-2 font-medium">
					<Icon name="ph:list-bullets-bold" size="20" />
					<span>Оглавление</span>
				</NuxtLink>

				<div class="flex items-center gap-4 w-full sm:w-auto">
					<NuxtLink 
						v-if="lastChapter?.next_chapter_id" 
						:to="`/novela/reader/${novelaId}/${lastChapter.next_chapter_id}`"
						class="flex-1 sm:flex-none flex items-center justify-center gap-2 px-6 py-3 bg-zinc-900 text-white dark:bg-zinc-100 dark:text-zinc-900 hover:bg-zinc-800 dark:hover:bg-zinc-200 rounded-xl font-bold transition-all active:scale-95 shadow-lg shadow-zinc-900/10 dark:shadow-none"
					>
						<span>Далее</span>
						<Icon name="ph:caret-right-bold" size="20" />
					</NuxtLink>
					<NuxtLink 
						v-else 
						:to="`/novela/${novelaId}`"
						class="flex-1 sm:flex-none flex items-center justify-center gap-2 px-6 py-3 bg-yellow-500 text-white rounded-xl font-bold transition-all active:scale-95"
					>
						<span>К описанию</span>
						<Icon name="ph:house-bold" size="20" />
					</NuxtLink>
				</div>
			</div>
		</footer>

		<!-- Lightbox -->
		<Transition name="fade">
			<div
				v-if="lightboxImage"
				class="fixed inset-0 z-200 bg-black/90 flex items-center justify-center p-4"
				@click="closeLightbox"
			>
				<button
					class="absolute top-4 right-4 p-2 text-white/70 hover:text-white transition-colors cursor-pointer"
					@click.stop="closeLightbox"
				>
					<Icon name="ph:x-bold" size="28" />
				</button>
				<img
					:src="lightboxImage"
					class="max-w-full max-h-full object-contain rounded-lg"
					@click.stop
					alt="NO IMAGE"
				/>
			</div>
		</Transition>
	</div>
</template>

<style scoped>
.prose {
	max-width: none;
}

.prose img {
	width: 100%;
	max-width: 28rem;
	height: 384px;
	object-fit: cover;
	border-radius: 0.75rem;
	margin-left: auto;
	margin-right: auto;
	display: block;
}

/* Custom transitions for smoother reading experience */
.fade-enter-active, .fade-leave-active {
	transition: opacity 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
	opacity: 0;
}
</style>
