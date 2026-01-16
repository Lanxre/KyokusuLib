<script setup lang="ts">
import { onClickOutside } from "@vueuse/core";
import { storeToRefs } from "pinia";
import { ref } from "vue";
import { staticImage } from "@/utils/str";

import BookMarkIcon from "@/assets/images/header/bookmark-white.png";
import LogoIcon from "@/assets/images/Kyokusu/kyokusulib_logo_2.png";
import NotificationIcon from "~/assets/images/header/ringing.png";
import SearchIcon from "@/assets/images/header/search-interface-symbol.png";
import CloseIcon from "@/assets/images/special/close.png";
import LogOutIcon from "@/assets/images/special/log-out.png";
import MenuIcon from "@/assets/images/special/menu.png";
import ReferIcon from "@/assets/images/special/refer.png";
import SettingsIcon from "@/assets/images/special/setting.png";
import PersonIcon from "@/assets/images/special/user.png";
import ColorPaletteIcon from "@/assets/images/special/color-palette.png";
import BookIcon from "@/assets/images/special/open-book.png";
import WritterIcon from "@/assets/images/special/writer.png";

import { useAuthStore } from "@/stores/auth";
import { useRolePermissions } from "~/composables/api/role/useRolePermissions";
import { KyokusuAppRole } from "@/types/enums/role-enum";

const router = useRouter();
const authStore = useAuthStore();
const { user, isAuthenticated } = storeToRefs(authStore);
const { hasPermission } = useRolePermissions();

const isMobileMenuOpen = ref(false);
const isUserDropdownOpen = ref(false);
const isContentSubmenuOpen = ref(false);
const userDropdownRef = ref(null);

const toggleMobileMenu = () =>
	(isMobileMenuOpen.value = !isMobileMenuOpen.value);
const closeMobileMenu = () => (isMobileMenuOpen.value = false);

onClickOutside(userDropdownRef, () => {
	isUserDropdownOpen.value = false;
	isContentSubmenuOpen.value = false;
});

const handleLogout = async () => {
	await authStore.logout();
	isUserDropdownOpen.value = false;
	closeMobileMenu();
};

const goToLogin = () => {
	closeMobileMenu();
	router.push("/login");
};
</script>

<template>
  <header class="w-full border-b-2 border-zinc-200 dark:border-zinc-700 bg-white dark:bg-[#0f0f0f] z-50 sticky top-0 text-zinc-900 dark:text-white transition-colors duration-300">
    <div class="mx-auto flex items-center justify-between px-4 py-4 md:px-8 lg:px-12 max-w-screen-2xl">
      
      <!-- LOGO & NAV -->
      <div class="flex items-center gap-4 md:gap-8 flex-1">
        <NuxtLink to="/" class="shrink-0" @click="closeMobileMenu">
          <img :src="LogoIcon" alt="KLib Logo" class="h-12 w-12 md:h-14 md:w-14 dark:invert transition-all" />
        </NuxtLink>

        <nav class="hidden md:flex items-center gap-2 lg:gap-4">
          <NuxtLink to="/catalog" class="px-6 py-3 rounded-2xl bg-zinc-300 hover:bg-zinc-200 dark:bg-zinc-800/80 dark:hover:bg-zinc-700 transition-colors font-medium text-zinc-700 hover:text-zinc-900 dark:text-zinc-200 dark:hover:text-white">Каталог</NuxtLink>
          <NuxtLink to="/top" class="px-6 py-3 rounded-2xl bg-zinc-300 hover:bg-zinc-200 dark:bg-zinc-800/80 dark:hover:bg-zinc-700 transition-colors font-medium text-zinc-700 hover:text-zinc-900 dark:text-zinc-200 dark:hover:text-white">Топы</NuxtLink>
          <NuxtLink to="/forum" class="px-6 py-3 rounded-2xl bg-zinc-300 hover:bg-zinc-200 dark:bg-zinc-800/80 dark:hover:bg-zinc-700 transition-colors font-medium text-zinc-700 hover:text-zinc-900 dark:text-zinc-200 dark:hover:text-white">Форум</NuxtLink>
        </nav>
      </div>

      <!-- RIGHT SIDE ACTIONS -->
      <div class="flex items-center gap-3 md:gap-4">
        
        <!-- Search Desktop -->
        <div class="hidden w-64 md:flex items-center gap-3 px-4 py-3 rounded-full bg-zinc-300 hover:bg-zinc-200 dark:bg-zinc-800 dark:hover:bg-zinc-700 transition-colors cursor-pointer group">
          <img :src="SearchIcon" alt="Search" class="h-5 w-5 opacity-70 group-hover:opacity-100 transition-opacity dark:invert" />
          <span class="text-sm lg:text-base font-medium text-zinc-500 group-hover:text-zinc-900 dark:text-zinc-400 dark:group-hover:text-white transition-colors">Поиск</span>
        </div>
        
        <!-- Notifications -->
        <button class="p-3 rounded-full bg-zinc-300 hover:bg-zinc-200 dark:bg-zinc-800/80 dark:hover:bg-zinc-700 transition-colors">
          <img :src="NotificationIcon" alt="Notifications" class="h-6 w-6 dark:invert" />
        </button>

        <!-- Bookmarks -->
        <NuxtLink to="/bookmarks" class="hidden sm:flex items-center gap-2 px-4 py-3 rounded-full bg-zinc-300 hover:bg-zinc-200 dark:bg-zinc-800 dark:hover:bg-zinc-700 transition-colors">
          <img :src="BookMarkIcon" alt="Bookmarks" class="h-5 w-5 dark:invert invert-0" />
          <span class="hidden lg:inline font-medium">Закладки</span>
        </NuxtLink>

        <!-- USER DROPDOWN -->
        <div class="relative" ref="userDropdownRef">
            <button 
                v-if="isAuthenticated && user" 
                @click="isUserDropdownOpen = !isUserDropdownOpen"
                class="flex items-center gap-2 focus:outline-none"
            >
                <div class="h-10 w-10 cursor-pointer rounded-full bg-zinc-200 dark:bg-zinc-700 flex items-center justify-center overflow-hidden border-2 border-transparent hover:border-zinc-400 dark:hover:border-zinc-300/50 transition-colors">
                    <img :src="staticImage(user.picture)" class="w-full h-full object-cover" alt="Avatar" />
                </div>
            </button>

            <button 
                v-else 
                @click="goToLogin"
                class="hidden md:flex justify-center gap-2 items-center bg-zinc-300 hover:bg-zinc-200 dark:bg-zinc-800 dark:hover:bg-zinc-700 transition-colors px-6 py-3 rounded-2xl cursor-pointer"
            >
                <img :src="ReferIcon" alt="Login" class="h-5 w-5 dark:brightness-0 dark:invert" />
                <span class="font-medium">Войти</span>
            </button>

            <!-- DROPDOWN MENU -->
            <Transition name="fade">
               <div v-if="isUserDropdownOpen" class="absolute right-0 mt-3 w-56 bg-white dark:bg-[#18181b] border border-zinc-200 dark:border-zinc-700 rounded-xl shadow-xl py-2 px-2 overflow-visible flex flex-col z-60">
                 
                 <div class="px-4 py-2 border-b border-zinc-200 dark:border-zinc-700/50 mb-1">
                   <p class="text-sm text-center font-semibold text-zinc-900 dark:text-white truncate">{{ user?.name }}</p>
                   <p class="text-xs text-zinc-500 truncate">{{ user?.email }}</p>
                 </div>
           
                 <NuxtLink :to="`/profile/${user?.id}`" class="flex justify-between gap-4 items-center rounded-full px-4 py-2 text-sm text-zinc-600 hover:bg-zinc-100 hover:text-zinc-900 dark:text-zinc-300 dark:hover:bg-zinc-700 dark:hover:text-white transition-colors text-left w-full" @click="isUserDropdownOpen = false">
                   <span>Профиль</span>
                   <img :src="PersonIcon" class="h-4 w-4 dark:brightness-0 dark:invert" />
                 </NuxtLink>
           
                 <NuxtLink to="/profile/settings" class="flex justify-between gap-4 items-center rounded-full px-4 py-2 text-sm text-zinc-600 hover:bg-zinc-100 hover:text-zinc-900 dark:text-zinc-300 dark:hover:bg-zinc-700 dark:hover:text-white transition-colors text-left w-full" @click="isUserDropdownOpen = false">
                   <span>Настройки</span>
                   <img :src="SettingsIcon" class="h-4 w-4 dark:brightness-0 dark:invert" />
                 </NuxtLink>
           
                 <!-- SUBMENU (ADD CONTENT) -->
                 <div 
                   class="relative"
                   @mouseenter="isContentSubmenuOpen = true"
                   @mouseleave="isContentSubmenuOpen = false"
                 >
                     <div v-if="hasPermission(KyokusuAppRole.MODERATOR)" class="flex justify-between gap-4 items-center rounded-full cursor-pointer px-4 py-2 text-sm text-zinc-600 hover:bg-zinc-100 hover:text-zinc-900 dark:text-zinc-300 dark:hover:bg-zinc-700 dark:hover:text-white transition-colors text-left w-full">
                       <span>Добавить контент</span>
                       <div class="flex items-center gap-2">
                            <img :src="ColorPaletteIcon" class="h-4 w-4 dark:brightness-0 dark:invert" />
                        </div>
                     </div>
           
                   <Transition name="fade">
                     <div 
                       v-if="isContentSubmenuOpen" 
                       class="absolute top-0 left-full ml-2 w-48 bg-white dark:bg-[#18181b] border border-zinc-200 dark:border-zinc-700 rounded-xl shadow-xl py-2 px-2 z-[70]"
                     > 
                       <NuxtLink to="/novela/add" class="flex justify-between gap-4 items-center rounded-full px-4 py-2 text-sm text-zinc-600 hover:bg-zinc-100 hover:text-zinc-900 dark:text-zinc-300 dark:hover:bg-zinc-700 dark:hover:text-white transition-colors w-full" @click="isUserDropdownOpen = false">
                         <span>Новелла</span>
                         <img :src="BookIcon" class="h-4 w-4 dark:brightness-0 dark:invert" />
                       </NuxtLink>
           
                       <NuxtLink to="/author/add" class="flex justify-between gap-4 items-center rounded-full px-4 py-2 text-sm text-zinc-600 hover:bg-zinc-100 hover:text-zinc-900 dark:text-zinc-300 dark:hover:bg-zinc-700 dark:hover:text-white transition-colors w-full" @click="isUserDropdownOpen = false">
                         <span>Автора</span>
                         <img :src="WritterIcon" class="h-4 w-4 dark:brightness-0 dark:invert" />
                       </NuxtLink>
                     </div>
                   </Transition>
                 </div>
           
                 <div class="h-px bg-zinc-200 dark:bg-zinc-700/50 my-1"></div>
           
                 <button @click="handleLogout" class="flex justify-between gap-4 items-center rounded-full cursor-pointer px-4 py-2 text-sm text-zinc-600 hover:bg-red-50 hover:text-red-500 dark:text-zinc-300 dark:hover:bg-red-500/10 dark:hover:text-red-300 transition-colors text-left w-full">
                   <span>Выйти</span>
                   <img :src="LogOutIcon" class="h-4 w-4 dark:brightness-0 dark:invert" />
                 </button>
               </div>
             </Transition>
        </div>

        <!-- MOBILE BURGER -->
        <button 
          @click="toggleMobileMenu" 
          class="md:hidden p-2 active:scale-95 transition-transform bg-zinc-100 dark:bg-zinc-800 rounded-2xl"
        >
          <img :src="isMobileMenuOpen ? CloseIcon : MenuIcon" alt="Menu" class="h-7 w-7 dark:invert" />
        </button>
      </div>
    </div>

    <!-- MOBILE MENU -->
    <Transition name="slide-fade">
      <div 
        v-if="isMobileMenuOpen" 
        class="absolute left-0 w-full h-[calc(100vh-80px)] overflow-y-auto border-t border-zinc-200 dark:border-zinc-700 bg-white/95 dark:bg-[#0f0f0f]/95 backdrop-blur-md md:hidden shadow-xl"
      >
        <nav class="flex flex-col p-4 gap-2 pb-10">
            <div v-if="isAuthenticated && user" class="flex items-center gap-3 p-3 bg-zinc-100 dark:bg-zinc-800/50 rounded-xl mb-2">
                <NuxtLink :to="`/profile/${user.id}`" class="h-10 w-10 shrink-0 rounded-full flex items-center justify-center overflow-hidden" @click="closeMobileMenu">
                    <div class="h-10 w-10 rounded-full bg-zinc-600 flex items-center justify-center overflow-hidden">
                        <img :src="staticImage(user.picture)" class="w-full h-full object-cover" />
                    </div>
                </NuxtLink>
                <div class="flex-1 min-w-0">
                    <p class="font-medium text-zinc-900 dark:text-white truncate">{{ user.name }}</p>
                    <p class="text-xs text-zinc-500 dark:text-zinc-400 truncate">{{ user.email }}</p>
                </div>
                <button @click="handleLogout" class="p-2 text-zinc-500 hover:text-zinc-900 dark:text-zinc-400 dark:hover:text-white">
                    <img :src="LogOutIcon" class="h-5 w-5 dark:invert" />
                </button>
            </div>
            
            <button v-else @click="goToLogin" class="w-full py-3 bg-zinc-100 dark:bg-zinc-800 rounded-xl font-medium mb-4 text-zinc-900 dark:text-white">
                Войти в аккаунт
            </button>

          <NuxtLink to="/catalog" class="w-full text-center py-3 text-lg font-medium hover:bg-zinc-100 dark:hover:bg-zinc-800 text-zinc-900 dark:text-white rounded-xl transition-colors" @click="closeMobileMenu">Каталог</NuxtLink>
          <NuxtLink to="/top" class="w-full text-center py-3 text-lg font-medium hover:bg-zinc-100 dark:hover:bg-zinc-800 text-zinc-900 dark:text-white rounded-xl transition-colors" @click="closeMobileMenu">Топы</NuxtLink>
          <NuxtLink to="/forum" class="w-full text-center py-3 text-lg font-medium hover:bg-zinc-100 dark:hover:bg-zinc-800 text-zinc-900 dark:text-white rounded-xl transition-colors" @click="closeMobileMenu">Форум</NuxtLink>
          
          <div class="w-full border-t border-zinc-200 dark:border-zinc-700 my-2 opacity-50"></div>
          
          <NuxtLink to="/search" class="flex items-center gap-3 px-6 py-3 rounded-xl bg-zinc-100 dark:bg-zinc-800/50 hover:bg-zinc-200 dark:hover:bg-zinc-800 w-full justify-center transition-colors text-zinc-700 dark:text-zinc-200" @click="closeMobileMenu">
            <img :src="SearchIcon" class="h-5 w-5 dark:invert" />
            <span>Поиск</span>
          </NuxtLink>
          
          <NuxtLink to="/bookmarks" class="flex items-center gap-3 px-6 py-3 rounded-xl bg-zinc-100 dark:bg-zinc-800/50 hover:bg-zinc-200 dark:hover:bg-zinc-800 w-full justify-center transition-colors text-zinc-700 dark:text-zinc-200" @click="closeMobileMenu">
            <img :src="BookMarkIcon" class="h-5 w-5 invert dark:invert-0" />
            <span>Закладки</span>
          </NuxtLink>

          <div v-if="hasPermission(KyokusuAppRole.MODERATOR)" class="w-full border-t border-zinc-200 dark:border-zinc-700 my-2 opacity-50"></div>

          <div v-if="hasPermission(KyokusuAppRole.MODERATOR)" class="flex flex-col items-center gap-2">
            <p class="text-xs text-zinc-500 uppercase font-bold px-2 mb-1">Добавить контент</p>
            <NuxtLink to="/novela/add" class="flex justify-center items-center gap-3 px-6 py-3 rounded-xl bg-zinc-50 dark:bg-zinc-900/30 hover:bg-zinc-200 dark:hover:bg-zinc-800 w-full transition-colors text-zinc-700 dark:text-zinc-200" @click="closeMobileMenu">
                <img :src="BookIcon" class="h-5 w-5 dark:invert opacity-70" />
                <span>Ранобэ</span>
            </NuxtLink>

            <NuxtLink to="/author/add" class="flex justify-center items-center gap-3 px-6 py-3 rounded-xl bg-zinc-50 dark:bg-zinc-900/30 hover:bg-zinc-200 dark:hover:bg-zinc-800 w-full transition-colors text-zinc-700 dark:text-zinc-200" @click="closeMobileMenu">
                <img :src="WritterIcon" class="h-5 w-5 dark:invert opacity-70" />
                <span>Автора</span>
            </NuxtLink>
          </div>

        </nav>
      </div>
    </Transition>
  </header>
</template>

<style scoped>
.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  opacity: 1;
  transform: translateY(0);
}
.slide-fade-enter-from,
.slide-fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-5px);
}
</style>