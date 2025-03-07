<template>
    <router-link :to="{name: 'analysis', params: {ticker: currentTicker}}" 
    class="inline-flex items-center justify-center
    bg-blue-600 dark:bg-blue-500
    text-white dark:text-white
    hover:bg-blue-700 dark:hover:bg-blue-600
    focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-500
    border border-transparent
    rounded-full
    px-4 py-2
    transition-all duration-200
    transform hover:-translate-y-0.5 hover:scale-105
    shadow-md hover:shadow-lg
    active:scale-100 active:shadow-sm"
    >
      Analyse stock with AI (beta)
    </router-link>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

defineProps({
  currentTicker: {
    type: String,
    required: true,
  },
});
</script>