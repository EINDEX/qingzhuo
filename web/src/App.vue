<template>
  <div class="h-screen">
    <nav-bar class="h-auto" />
    <header class="shadow h-auto">
      <!-- <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
      <h1 class="text-3xl font-bold text-gray-900">Dashboard</h1>
    </div> -->
    </header>
    <main class="max-w-3xl mx-auto px-4 h-auto">
      <router-view />
    </main>

    <blog-footer class="h-auto" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import BlogFooter from '/@/components/BlogFooter.vue';
import NavBar from '/@/components/NavBar.vue';

export default defineComponent({
  name: 'App',
  components: {
    NavBar,
    BlogFooter,
  },
});
</script>

<style>
#app {
  font-family: Inter var, system-ui, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto,
    Helvetica Neue, Arial, Noto Sans, sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol,
    Noto Color Emoji;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100%;
}
main {
  min-height: 80%;
}
</style>
