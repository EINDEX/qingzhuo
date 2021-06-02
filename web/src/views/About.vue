<template>
  <post-page :title="title">
    <p>我是 EINDEX</p>
    <p>一个咸鱼开发者，摄影、咖啡爱好者</p>

    <p>：）</p>
  </post-page>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import PostPage from '/@/components/PostPage.vue';

export default defineComponent({
    components: { PostPage },
    setup() {},
    data() {
      return {
        title: 'About Me',
      };
    },
  });
</script>
