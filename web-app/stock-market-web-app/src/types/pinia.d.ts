import { Pinia } from 'pinia';

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $pinia: Pinia;
  }
}