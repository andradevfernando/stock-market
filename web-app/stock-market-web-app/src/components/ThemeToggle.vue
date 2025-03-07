
<script setup lang="ts">   
    import {useThemeStore} from '../stores/themeStore';
    import { onMounted } from 'vue';

const themeStore = useThemeStore();

onMounted(() => {
  themeStore.initializeTheme();
});
</script>

<template>
    <button
      @click="themeStore.toggleTheme()"
      class="bg-transparent flex justify-center items-center cursor-pointer hover:bg-white/30 px-2 py-1 rounded-xl transition-colors"
    >
    <div class="text-4xl">
      {{ themeStore.themeIcon }}
    </div>
    </button>
</template>

