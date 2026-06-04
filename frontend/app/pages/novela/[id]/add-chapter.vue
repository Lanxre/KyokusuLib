<script setup lang="ts">
import { useNovela } from "~/composables/api/novela/useNovela";

const route = useRoute();
const router = useRouter();
const novelaId = route.params.id as string;

const { novela, fetchNovela, addChapter, addChapterImage } = useNovela();

await useAsyncData(`novela-${novelaId}`, () => fetchNovela(novelaId));

const volumes = computed(() => novela.value?.volumes ?? []);

const selectedVolumeId = ref(volumes.value[0]?.id ?? "");
const chapterNumber = ref(1);
const chapterTitle = ref("");
const chapterContent = ref("");
const isSubmitting = ref(false);

const images = ref<{ imageUrl: string; caption: string; position: number }[]>(
	[],
);

function addImageField() {
	images.value.push({ imageUrl: "", caption: "", position: images.value.length });
}

function removeImageField(index: number) {
	images.value.splice(index, 1);
}

async function onSubmit() {
	if (!selectedVolumeId.value || !chapterContent.value) return;
	if (!chapterTitle.value) {
		chapterTitle.value = `Глава ${chapterNumber.value}`;
	}
	isSubmitting.value = true;
	try {
		const res = await addChapter(
			Number(novelaId),
			selectedVolumeId.value,
			chapterNumber.value,
			chapterTitle.value,
			chapterContent.value,
		);
		if (res?.id) {
			for (const img of images.value) {
				if (img.imageUrl) {
					await addChapterImage(res.id, img.imageUrl, img.caption, img.position);
				}
			}
		}
		router.push(`/novela/${novelaId}`);
	} finally {
		isSubmitting.value = false;
	}
}
</script>

<template>
	<div class="min-h-screen bg-zinc-50 dark:bg-[#0f0f0f] text-zinc-900 dark:text-zinc-200">
		<div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
			<div class="flex items-center gap-4 mb-8">
				<button
					class="p-2 rounded-xl hover:bg-zinc-200 dark:hover:bg-zinc-800 transition-colors cursor-pointer"
					@click="router.push(`/novela/${novelaId}`)"
				>
					<Icon name="ph:arrow-left-bold" size="20" />
				</button>
				<div>
					<h1 class="text-2xl font-black tracking-tight">Добавить главу</h1>
					<p v-if="novela" class="text-sm text-zinc-500 mt-1">{{ novela.title }}</p>
				</div>
			</div>

			<form
				class="space-y-6 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-2xl p-6"
				@submit.prevent="onSubmit"
			>
				<div>
					<label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
						Том
					</label>
					<select
						v-model="selectedVolumeId"
						class="w-full px-3 py-2.5 rounded-xl border border-zinc-200 dark:border-zinc-700 bg-white dark:bg-zinc-800 text-zinc-900 dark:text-zinc-200 text-sm focus:ring-2 focus:ring-zinc-500 outline-none transition-all"
					>
						<option v-for="vol in volumes" :key="vol.id" :value="vol.id">
							Том {{ vol.number }}{{ vol.title ? ` — ${vol.title}` : "" }}
						</option>
					</select>
				</div>

				<div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
					<div>
						<label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
							Номер главы
						</label>
						<input
							v-model.number="chapterNumber"
							type="number"
							min="1"
							step="0.5"
							class="w-full px-3 py-2.5 rounded-xl border border-zinc-200 dark:border-zinc-700 bg-white dark:bg-zinc-800 text-zinc-900 dark:text-zinc-200 text-sm focus:ring-2 focus:ring-zinc-500 outline-none transition-all"
						/>
					</div>
					<div>
						<label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
							Название
						</label>
						<input
							v-model="chapterTitle"
							type="text"
							placeholder="Название главы"
							class="w-full px-3 py-2.5 rounded-xl border border-zinc-200 dark:border-zinc-700 bg-white dark:bg-zinc-800 text-zinc-900 dark:text-zinc-200 text-sm focus:ring-2 focus:ring-zinc-500 outline-none transition-all"
						/>
					</div>
				</div>

				<div>
					<label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
						Содержание
					</label>
					<textarea
						v-model="chapterContent"
						rows="12"
						placeholder="Текст главы..."
						class="w-full px-3 py-2.5 rounded-xl border border-zinc-200 dark:border-zinc-700 bg-white dark:bg-zinc-800 text-zinc-900 dark:text-zinc-200 text-sm focus:ring-2 focus:ring-zinc-500 outline-none transition-all resize-y font-mono leading-relaxed"
					></textarea>
				</div>

				<div class="space-y-3">
					<div class="flex items-center justify-between">
						<label class="text-sm font-medium text-zinc-700 dark:text-zinc-300">
							Изображения
						</label>
						<button
							type="button"
							class="text-sm text-blue-600 dark:text-blue-400 hover:underline cursor-pointer"
							@click="addImageField"
						>
							+ Добавить изображение
						</button>
					</div>

					<div
						v-for="(img, index) in images"
						:key="index"
						class="flex flex-col sm:flex-row gap-2 p-3 rounded-xl bg-zinc-50 dark:bg-zinc-800/50 border border-zinc-200 dark:border-zinc-700"
					>
						<div class="flex-1">
							<input
								v-model="img.imageUrl"
								type="text"
								placeholder="URL изображения"
								class="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-700 bg-white dark:bg-zinc-800 text-zinc-900 dark:text-zinc-200 text-sm focus:ring-2 focus:ring-zinc-500 outline-none transition-all"
							/>
						</div>
						<div class="w-full sm:w-40">
							<input
								v-model="img.caption"
								type="text"
								placeholder="Подпись"
								class="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-700 bg-white dark:bg-zinc-800 text-zinc-900 dark:text-zinc-200 text-sm focus:ring-2 focus:ring-zinc-500 outline-none transition-all"
							/>
						</div>
						<div class="w-full sm:w-24">
							<input
								v-model.number="img.position"
								type="number"
								min="0"
								placeholder="Поз."
								class="w-full px-3 py-2 rounded-lg border border-zinc-200 dark:border-zinc-700 bg-white dark:bg-zinc-800 text-zinc-900 dark:text-zinc-200 text-sm focus:ring-2 focus:ring-zinc-500 outline-none transition-all"
							/>
						</div>
						<button
							type="button"
							class="p-2 text-red-500 hover:text-red-700 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors cursor-pointer shrink-0"
							@click="removeImageField(index)"
						>
							<Icon name="ph:trash-bold" size="16" />
						</button>
					</div>

					<p v-if="!images.length" class="text-sm text-zinc-400 italic">
						Изображения не добавлены
					</p>
				</div>

				<div class="flex items-center justify-end gap-3 pt-2">
					<button
						type="button"
						class="px-4 py-2 rounded-xl text-sm font-medium border border-zinc-200 dark:border-zinc-700 text-zinc-600 dark:text-zinc-400 hover:bg-zinc-100 dark:hover:bg-zinc-800 transition-colors cursor-pointer"
						@click="router.push(`/novela/${novelaId}`)"
					>
						Отмена
					</button>
					<button
						type="submit"
						:disabled="isSubmitting"
						class="px-6 py-2 rounded-xl text-sm font-bold bg-zinc-900 text-white dark:bg-zinc-100 dark:text-zinc-900 hover:bg-zinc-800 dark:hover:bg-zinc-200 disabled:opacity-50 transition-all cursor-pointer"
					>
						{{ isSubmitting ? "Добавление..." : "Добавить главу" }}
					</button>
				</div>
			</form>
		</div>
	</div>
</template>
