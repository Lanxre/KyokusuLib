<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { ModalWindow, Label } from "@kyokusu-ui/vue";

import DashboardUserImageEdit from "./DashboardUserImageEdit.vue";
import DashboardUserBasicEdit from "./DashboardUserBasicEdit.vue";
import DashboardUserRoleEdit from "./DashboardUserRoleEdit.vue";
import DashboardUserVisibilityEdit from "./DashboardUserVisibilityEdit.vue";
import DashboardUserLevelEdit from "./DashboardUserLevelEdit.vue";
import DashboardUserTagsEdit from "./DashboardUserTagsEdit.vue";

import type { DashboardUser } from "@/types/frontend/dashboard/users";

import { useUserEdit } from "@/composables/api/dashboard/useUserEdit";
import { useRolePermissions } from "@/composables/api/role/useRolePermissions";
import { KyokusuAppRole } from "~/types/enums/role-enum";

const props = defineProps<{
	modelValue: boolean;
	user: DashboardUser | null;
}>();

const emit = defineEmits<{
	"update:modelValue": [value: boolean];
	saved: [];
}>();

const visible = computed({
	get: () => props.modelValue,
	set: (val) => emit("update:modelValue", val),
});

const saving = ref(false);

const { form, isDirty, loadUser, reset, save } = useUserEdit();

const { hasPermission } = useRolePermissions();

const userLevel = ref(1);
const userExperience = ref(0);
const userTags = ref<DashboardUser["tags"]>([]);

watch(
	() => props.user,
	(u) => {
		if (u) {
			loadUser(u);
			userLevel.value = u.user_level?.level ?? 1;
			userExperience.value = u.user_level?.experience ?? 0;
			userTags.value = [...(u.tags ?? [])];
		}
	},
	{ immediate: true },
);

async function handleSave() {
	if (!props.user) return;
	saving.value = true;
	const ok = await save(props.user.id);
	saving.value = false;
	if (ok) {
		emit("saved");
		visible.value = false;
	}
}

function handleCancel() {
	reset();
	visible.value = false;
}
</script>

<template>
	<ModalWindow
		v-model="visible"
		:title="'Редактирование: ' + user?.name"
		center-title
		width="w-full max-w-250"
	>
		<div v-if="user" class="flex flex-col gap-6">
			<!-- image -->
			<div class="flex flex-col gap-2">
				<Label label="Настройка изображения"/>
				<div class="flex justify-center p-4 rounded-lg border bg-white dark:bg-zinc-900/50 dark:border-zinc-700/50">
					<DashboardUserImageEdit
						v-model:avatar="form.picture"
						v-model:banner="form.banner"
					/>
				</div>
			</div>

			<!-- basic info -->
			<div class="flex flex-col gap-2">
				<Label label="Основные данные"/>
				<div class="flex p-4 rounded-lg border bg-white dark:bg-zinc-900/50 dark:border-zinc-700/50">
					<DashboardUserBasicEdit :form="form" />
				</div>
			</div>

			<!-- role & status -->
			<div class="flex flex-col gap-2">
				<Label label="Роль и статус"/>
				<div class="flex p-4 rounded-lg border bg-white dark:bg-zinc-900/50 dark:border-zinc-700/50">
					<DashboardUserRoleEdit :form="form" />
				</div>
			</div>
			
			<div v-if="hasPermission(KyokusuAppRole.ADMIN)" class="flex flex-col gap-2">
			    <Label label="Уровень и опыт"/>
				<div class="flex p-4 rounded-lg border bg-white dark:bg-zinc-900/50 dark:border-zinc-700/50">
					<DashboardUserLevelEdit
							:key="user?.id"
							v-model:level="userLevel"
							v-model:experience="userExperience"
						/>
				</div>
			</div>

			<div class="flex flex-col gap-2">
			    <Label label="Теги"/>
				<div class="flex p-4 rounded-lg border bg-white dark:bg-zinc-900/50 dark:border-zinc-700/50">
					<DashboardUserTagsEdit
						:key="user?.id"
						v-model:tags="userTags"
					/>
				</div>
			</div>

			<div v-if="hasPermission(KyokusuAppRole.ADMIN)" class="flex flex-col gap-2">
			    <Label label="Настройки видимости"/>
				<div class="flex p-4 rounded-lg border bg-white dark:bg-zinc-900/50 dark:border-zinc-700/50">
					<DashboardUserVisibilityEdit :form="form" />
				</div>
			</div>
		</div>
	</ModalWindow>
</template>