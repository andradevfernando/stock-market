
<template>
    <div class="flex justify-between items-center p-3 bg-white dark:bg-gray-600 rounded-lg shadow-sm">
      <span class="text-gray-600 dark:text-gray-300">{{ label }}</span>
      <span :class="color" class="font-semibold">
        {{ prefix }}{{ value }}{{ suffix }}
      </span>
    </div>
  </template>
  
  <script setup lang="ts">
  defineProps({
    label: String,
    value: [String, Number],
    prefix: {
      type: String,
      default: ''
    },
    suffix: {
      type: String,
      default: ''
    },
    color: {
      type: String,
      default: 'text-gray-900 dark:text-gray-100'
    }
  })
  </script>