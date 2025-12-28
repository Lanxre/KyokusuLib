<script setup lang="ts">
import DiscordIcon from "@/assets/images/social/discord.png";
import TelegramIcon from "@/assets/images/social/telegram.png";

</script>

<template>
  <footer
    class="w-full border-t-2 border-zinc-200 dark:border-zinc-700 p-6 min-h-[150px] mt-auto bg-zinc-50 dark:bg-[#0f0f0f] text-zinc-900 dark:text-white transition-colors duration-300"
  >
    <div class="mx-auto max-w-7xl">
      <div class="grid grid-cols-1 gap-8 sm:grid-cols-2 md:grid-cols-4">
        <div class="flex flex-col gap-3 items-center text-center md:items-start md:text-left">
          <RouterLink to="/">
            <h1 class="text-xl font-bold sm:text-2xl hover:text-zinc-600 dark:hover:text-zinc-300 transition-colors">KyokusuLib</h1>
          </RouterLink>
          <p class="text-sm leading-5 opacity-80 break-words w-3/4 mx-auto md:mx-0 text-zinc-600 dark:text-zinc-400">
            Если у Вас есть вопросы, обратитесь в техподдержку, напишите нам на почту
            <a href="mailto:burtsevvitalij@gmail.com" class="text-blue-600 dark:text-blue-400 hover:text-blue-500 dark:hover:text-blue-300 transition-colors">
              burtsevvitalij@gmail.com
            </a>
          </p>
        </div>

        <div class="flex flex-col gap-3 items-center text-center md:items-start md:text-left">
          <h1 class="text-xl font-bold sm:text-2xl">Информация</h1>
          <p class="text-sm leading-5 opacity-80 w-3/4 mx-auto md:mx-0 text-zinc-600 dark:text-zinc-400">
            Политика конфиденциальности и условия использования.
          </p>
        </div>

        <div class="flex flex-col gap-3 items-center text-center md:items-start md:text-left">
          <h1 class="text-xl font-bold sm:text-2xl">Разделы</h1>
          <p class="text-sm leading-5 opacity-80 w-3/4 mx-auto md:mx-0 text-zinc-600 dark:text-zinc-400">
            Популярное, Новинки, Каталог манги.
          </p>
        </div>

        <div class="flex flex-col gap-3 items-center text-center md:items-start md:text-left">
          <h1 class="text-xl font-bold sm:text-2xl">Социальные сети</h1>
          <div class="flex flex-row gap-4 w-3/4 justify-center items-center md:justify-center mx-auto md:mx-0">
              <a href="#" class="bg-zinc-300 hover:bg-zinc-400 dark:bg-transparent dark:hover:bg-zinc-800 rounded-xl transition-colors">
                <img :src="TelegramIcon" alt="Telegram" class="w-10 p-2 brightness-0 dark:invert"/>
              </a>
              <a href="#" class="bg-zinc-300 hover:bg-zinc-400 dark:bg-transparent dark:hover:bg-zinc-800 rounded-xl transition-colors">
                <img :src="DiscordIcon" alt="Discord" class="w-10 p-1 brightness-0 dark:invert"/>
              </a>
          </div>
        </div>

        
      </div>
        <div class="flex justify-between p-8 text-zinc-500 dark:text-stone-400 mt-12 border-t border-zinc-200 dark:border-stone-600 gap-4 max-[720px]:flex-col max-[720px]:items-center">
          <span>© 2025 KyokusuLib. Все права защищены.</span>
        </div>
     </div>
  </footer>
</template>