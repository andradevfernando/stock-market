
import { defineStore } from 'pinia';

interface CounterState {
  count: number;
}

export const useCounterStore = defineStore('counter', {
  state: (): CounterState => ({
    count: 0,
  }),
  actions: {
    increment() {
      this.count++;
    },
  },
  getters: {
    doubleCount: (state) => state.count * 2,
  },
});