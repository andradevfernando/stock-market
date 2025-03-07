<template>
    <div class="max-w-3xl mx-auto px-4 py-8">
  <div class="bg-white dark:bg-slate-800 rounded-full shadow-sm p-3 flex items-center space-x-3 border border-gray-200 dark:border-slate-700 transition-all duration-300 hover:shadow-md">
    <input 
      v-model="company" 
      placeholder="Company (ex: Apple)"
      class="flex-1 bg-transparent px-4 py-1 text-sm placeholder-gray-400 dark:placeholder-slate-500 focus:outline-none focus:ring-0 border-none"
    >
    
    <div class="h-6 w-px bg-gray-200 dark:bg-slate-700"></div>
    
    <input 
      type="date" 
      v-model="startDate"
      class="flex-1 bg-transparent px-2 py-1 text-sm text-blue-600 dark:text-white focus:outline-none border-none [&::-webkit-calendar-picker-indicator]:opacity-50"
    >
    
    <div class="h-6 w-px bg-gray-200 dark:bg-slate-700"></div>
        <input 
          type="date" 
          v-model="endDate"
          class="flex-1 bg-transparent px-2 py-1 text-sm text-blue-600 dark:text-white focus:outline-none border-none [&::-webkit-calendar-picker-indicator]:opacity-50"
        >
   <div class="h-6 w-px bg-gray-200 dark:bg-slate-700"></div>

   <div  @click.stop="applyFilters"  class="dark:text-white text-gray-900 hover:text-blue-600 rounded-full cursor-pointer transition-colors flex items-center space-x-2"> 
     <p>Search</p> 
     <p>stocks</p> 
     <svg
     class="w-8 h-8 p-1.5"
     viewBox="0 0 24 24"
     fill="none"
     style="pointer-events: bounding-box;" 
     >
    <path 
      d="M14.9536 14.9458L21 21M17 10C17 13.866 13.866 17 10 17C6.13401 17 3 13.866 3 10C3 6.13401 6.13401 3 10 3C13.866 3 17 6.13401 17 10Z" 
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
      vector-effect="non-scaling-stroke"
    />
  </svg>
  </div>
    </div>
  
</div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  
  const emit = defineEmits(['filter']);
  
  const now = new Date();
  const company = ref('');
  const endDate = ref(now);
  const startDate = ref(new Date(now.getTime() - (48 * 60 * 60 * 1000)));

  const applyFilters = () => {

    emit('filter', {
      company: company.value,
      startDate: new Date(startDate.value).toISOString(),
      endDate: new Date(endDate.value).toISOString(),
    });
  };
  </script>