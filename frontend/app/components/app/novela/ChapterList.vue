<script setup lang="ts">
import { ref, toRef } from "vue";
import type { NovelaVolume } from "~/types/backend/novela";
import { useChapterList } from "~/composables/ui/useChapterList";
import VolumeTabBar from "./VolumeTabBar.vue";
import ChapterSearchHeader from "./ChapterSearchHeader.vue";
import ChapterListContent from "./ChapterListContent.vue";
import NovelaAddVolumeModal from "./NovelaAddVolumeModal.vue";

const props = defineProps<{
	volumes: NovelaVolume[];
	canManage?: boolean;
	novelaId: number;
}>();

const router = useRouter();

const {
	activeVolumeId,
	searchQuery,
	sortedVolumes,
	filteredChapters,
	setActiveVolume,
} = useChapterList(toRef(props, "volumes"));

const showAddVolumeModal = ref(false);

function goToAddChapter() {
	router.push(`/novela/${props.novelaId}/add-chapter`);
}
</script>

<template>
	<div class="space-y-4">
		<div class="flex flex-col sm:flex-row gap-4 justify-between items-start sm:items-center">
			<VolumeTabBar
				:volumes="sortedVolumes"
				:active-volume-id="activeVolumeId"
				@select="setActiveVolume"
			/>

			<ChapterSearchHeader
				v-model="searchQuery"
				:can-manage="canManage"
				@add-volume="showAddVolumeModal = true"
				@add-chapter="goToAddChapter"
			/>
		</div>

		<ChapterListContent :chapters="filteredChapters" />

		<NovelaAddVolumeModal
			v-if="canManage"
			v-model="showAddVolumeModal"
			:novela-id="novelaId"
			:volumes="sortedVolumes"
		/>
	</div>
</template>

