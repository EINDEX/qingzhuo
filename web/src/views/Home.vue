<template>
  <post-list></post-list>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import PostList from '/@/components/PostList.vue';

export default defineComponent({
  name: 'Home',
  components: {
    PostList,
  },
  setup() {
    PostList;
  },
});
</script>
