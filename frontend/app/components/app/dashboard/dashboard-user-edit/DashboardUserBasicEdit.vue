<script setup lang="ts">
import { Input, RichText, DatePicker } from "@kyokusu-ui/vue";
import TabToggleSettings from "../../settings/TabToggleSettings.vue";
import type { UserEditForm } from "@/composables/api/dashboard/useUserEdit";
import { GENDER_OPTIONS } from "@/composables/api/dashboard/useUserEdit";

defineProps<{
    form: UserEditForm;
}>();
</script>

<template>
    <div class="grid grid-cols-1 sm:grid-cols-2 gap-x-4 gap-y-5 w-full">
        <!-- name + gender in one row -->
        <Input
            v-model="form.name"
            label="Имя пользователя"
            placeholder="Введите имя"
        />

        <div class="flex flex-col gap-1.5">
            <label
                class="text-sm font-medium"
                :style="{ color: 'var(--k-editor-label)' }"
            >
                Пол
            </label>
            <select
                v-model="form.gender"
                class="w-full h-10 rounded-lg border px-3 text-sm outline-none transition focus:ring-2"
				:style="{
					backgroundColor: 'var(--k-editor-bg)',
					borderColor: 'var(--k-editor-border)',
					color: 'var(--k-editor-text)',
				}"
            >
                <option
                    v-for="opt in GENDER_OPTIONS"
                    :key="opt.value"
                    :value="opt.value"
                >
                    {{ opt.label }}
                </option>
            </select>
        </div>

        <!-- birthday -->
        <div class="sm:col-span-2">
            <ClientOnly>
                <DatePicker
                    id="birthday"
                    label="Дата рождения"
                    v-model="form.birthday"
                />
                <template #fallback>
                    <Input
                        v-model="form.birthday"
                        label="Дата рождения"
                    />
                </template>
            </ClientOnly>
        </div>

        <!-- is_public — full width -->
        <div class="sm:col-span-2">
            <TabToggleSettings
                v-model="form.isPublic"
                id="show_profile_toggle"
                title="Публичный профиль"
                imgSrc="ph:link-simple"
                description="Переключите для отключения публичного профиля на сайте."
            />
        </div>

        <!-- about — full width -->
        <div class="flex flex-col gap-1.5 sm:col-span-2">
            <RichText
                id="about"
                label="О себе"
                v-model="form.about"
                placeholder="Введите описание"
                :max-length="500"
            />
        </div>
    </div>
</template>
