<script lang="ts" setup>
import { computed, nextTick, onMounted, onUnmounted, ref } from "vue";
import { parseDateToLocale } from "@/api/utils/date";
import { useProfileSettings } from "@/composables/api/settings/useProfileSettings";

const props = defineProps<{
	modelValue: string;
	label?: string;
	id: string;
}>();

const emit = defineEmits(["update:modelValue"]);
const { isBirthDateDisable } = useProfileSettings();
const isOpen = ref(false);
const viewMode = ref<"days" | "years">("days");
const containerRef = ref<HTMLElement | null>(null);
const yearListRef = ref<HTMLElement | null>(null);

const viewDate = ref(new Date());

const months = [
	"Январь",
	"Февраль",
	"Март",
	"Апрель",
	"Май",
	"Июнь",
	"Июль",
	"Август",
	"Сентябрь",
	"Октябрь",
	"Ноябрь",
	"Декабрь",
];

const weekDays = ["Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"];

const parseDate = (value: string): Date | null => {
	if (!value) return null;
	if (value.includes(".")) {
		const [d, m, y] = value.split(".");
		return new Date(Number(y), Number(m) - 1, Number(d));
	}
	const date = new Date(value);
	return isNaN(date.getTime()) ? null : date;
};

onMounted(() => {
	if (props.modelValue) {
		const date = parseDate(props.modelValue);
		if (date) {
			viewDate.value = date;
		}
	}
	document.addEventListener("click", handleClickOutside);
});

onUnmounted(() => {
	document.removeEventListener("click", handleClickOutside);
});

const handleClickOutside = (event: MouseEvent) => {
	if (
		containerRef.value &&
		!containerRef.value.contains(event.target as Node)
	) {
		isOpen.value = false;
		viewMode.value = "days";
	}
};

const currentMonthName = computed(() => months[viewDate.value.getMonth()]);
const currentYear = computed(() => viewDate.value.getFullYear());

const calendarDays = computed(() => {
	const year = viewDate.value.getFullYear();
	const month = viewDate.value.getMonth();

	const firstDayOfMonth = new Date(year, month, 1);
	const daysInMonth = new Date(year, month + 1, 0).getDate();

	let startDay = firstDayOfMonth.getDay();
	startDay = startDay === 0 ? 6 : startDay - 1;

	const days = [];

	for (let i = 0; i < startDay; i++) {
		days.push({ day: null, fullDate: null });
	}

	for (let i = 1; i <= daysInMonth; i++) {
		const dateStr = `${year}-${String(month + 1).padStart(2, "0")}-${String(i).padStart(2, "0")}`;
		days.push({
			day: i,
			fullDate: dateStr,
		});
	}

	return days;
});

const availableYears = computed(() => {
	const current = new Date().getFullYear();
	const start = 1900;
	const years = [];
	for (let i = current; i >= start; i--) {
		years.push(i);
	}
	return years;
});

const toggleCalendar = () => {
	isOpen.value = !isOpen.value;
	viewMode.value = "days";
	if (isOpen.value && props.modelValue) {
		const date = parseDate(props.modelValue);
		if (date) {
			viewDate.value = date;
		}
	}
};

const prevMonth = () => {
	viewDate.value = new Date(
		viewDate.value.getFullYear(),
		viewDate.value.getMonth() - 1,
		1,
	);
};

const nextMonth = () => {
	viewDate.value = new Date(
		viewDate.value.getFullYear(),
		viewDate.value.getMonth() + 1,
		1,
	);
};

const switchToYearMode = async () => {
	viewMode.value = "years";
	await nextTick();
	if (yearListRef.value) {
		const selectedYearEl = yearListRef.value.querySelector(".selected-year");
		if (selectedYearEl) {
			selectedYearEl.scrollIntoView({ block: "center" });
		}
	}
};

const selectYear = (year: number) => {
	viewDate.value = new Date(year, viewDate.value.getMonth(), 1);
	viewMode.value = "days";
};

const selectDate = (dateStr: string) => {
	emit("update:modelValue", dateStr);
	isOpen.value = false;
};

const formattedValue = computed(() => {
	if (!props.modelValue) return "";

	return parseDateToLocale(props.modelValue);
});

const isSelected = (dateStr: string) => {
	if (!props.modelValue) return false;
	if (props.modelValue === dateStr) return true;

	if (props.modelValue.includes(".")) {
		const [d, m, y] = props.modelValue.split(".");
		return `${y}-${m}-${d}` === dateStr;
	}

	return false;
};

const isToday = (dateStr: string) => {
	const today = new Date();
	const todayStr = `${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2, "0")}-${String(today.getDate()).padStart(2, "0")}`;
	return dateStr === todayStr;
};
</script>

<template>
    <div class="space-y-2 relative" ref="containerRef">
        <label v-if="label" :for="id" class="text-sm font-medium text-dark dark:text-stone-300 ml-[3px]">
            {{ label }}
        </label>
        
        <div class="relative">
            <input 
                :id="id"
                type="text"
                readonly
                :disabled="isBirthDateDisable"
                :value="formattedValue"
                @click="toggleCalendar"
                class="w-full bg-zinc-950/50 border border-zinc-700 rounded-md py-2 px-3 text-sm text-zinc-200 placeholder-zinc-600 focus:outline-none focus:ring-2 focus:ring-zinc-600 focus:border-transparent cursor-pointer transition-all disabled:opacity-50 disabled:cursor-not-allowed"
                placeholder="ДД.ММ.ГГГГ"
            />
            <div class="absolute inset-y-0 right-0 flex items-center px-3 pointer-events-none text-stone-500">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
            </div>
        </div>

        <transition
            enter-active-class="transition duration-100 ease-out"
            enter-from-class="transform scale-95 opacity-0"
            enter-to-class="transform scale-100 opacity-100"
            leave-active-class="transition duration-75 ease-in"
            leave-from-class="transform scale-100 opacity-100"
            leave-to-class="transform scale-95 opacity-0"
        >
            <div v-if="isOpen && !isBirthDateDisable" class="absolute z-50 mt-1 w-64 bg-zinc-900 border border-zinc-700 rounded-lg shadow-xl p-4">
                
                <div class="flex items-center justify-between mb-4">
                    <button 
                        v-if="viewMode === 'days'"
                        @click.prevent="prevMonth"
                        class="p-1 hover:bg-zinc-800 rounded text-zinc-400 hover:text-white transition-colors"
                    >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path></svg>
                    </button>
                    <div v-else class="w-6"></div> 
                    
                    <div class="flex items-center gap-1">
                        <span class="text-sm font-medium text-white">
                            {{ viewMode === 'days' ? currentMonthName : 'Выберите год' }}
                        </span>
                        <button 
                            @click.prevent="viewMode === 'days' ? switchToYearMode() : viewMode = 'days'"
                            class="text-sm font-bold text-white hover:text-zinc-300 hover:bg-zinc-800 px-1 rounded transition-colors"
                        >
                            {{ currentYear }}
                        </button>
                    </div>
                    
                    <button 
                        v-if="viewMode === 'days'"
                        @click.prevent="nextMonth"
                        class="p-1 hover:bg-zinc-800 rounded text-zinc-400 hover:text-white transition-colors"
                    >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path></svg>
                    </button>
                    <div v-else class="w-6"></div>
                </div>

                <div v-if="viewMode === 'days'">
                    <div class="grid grid-cols-7 mb-2">
                        <div v-for="day in weekDays" :key="day" class="text-center text-xs text-stone-500 font-medium">
                            {{ day }}
                        </div>
                    </div>

                    <div class="grid grid-cols-7 gap-1">
                        <template v-for="(dayObj, index) in calendarDays" :key="index">
                            <div v-if="!dayObj.day"></div>
                            
                            <button 
                                v-else
                                @click.prevent="selectDate(dayObj.fullDate!)"
                                class="h-8 w-8 rounded flex items-center justify-center text-xs transition-colors"
                                :class="[
                                    isSelected(dayObj.fullDate!) 
                                        ? 'bg-zinc-100 text-zinc-900 font-bold' 
                                        : 'text-zinc-300 hover:bg-zinc-800',
                                    isToday(dayObj.fullDate!) && !isSelected(dayObj.fullDate!) 
                                        ? 'border border-zinc-700' 
                                        : ''
                                ]"
                            >
                                {{ dayObj.day }}
                            </button>
                        </template>
                    </div>
                </div>

                <div v-else ref="yearListRef" class="h-60 overflow-y-auto grid grid-cols-3 gap-2 pr-1 custom-scrollbar">
                    <button 
                        v-for="year in availableYears" 
                        :key="year"
                        @click.prevent="selectYear(year)"
                        class="py-2 text-sm rounded hover:bg-zinc-800 text-zinc-300 transition-colors"
                        :class="{ 'bg-zinc-800 text-white font-bold selected-year': year === currentYear }"
                    >
                        {{ year }}
                    </button>
                </div>

            </div>
        </transition>
    </div>
</template>