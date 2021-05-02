<template>
  <nav>
    <div class="bg-gray-200 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <a href="/">
              <img class="h-8 w-8 rounded-lg" src="/@/assets/avatar.png" alt="Workflow" />
            </a>
          </div>
          <div class="hidden md:block">
            <div class="ml-10 flex items-baseline space-x-4" :key="item.link" v-for="item in left">
              <router-link
                class="text-gray-700 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                :to="item.link"
                >{{ item.name }}</router-link
              >
            </div>
          </div>
        </div>
        <div class="hidden md:block">
          <div class="ml-4 flex items-center md:ml-6" :key="item.link" v-for="item in right">
            <router-link
              class="text-gray-700 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
              :to="item.link"
              >{{ item.name }}</router-link
            >
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  setup() {},
  data() {
    return {
      left: [
        //   { name: 'Home', link: '/' }
      ],
      right: [{ name: 'About Me', link: '/about' }],
    };
  },
});
</script>
