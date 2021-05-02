<template>
  <post-page :title="title">
    <p>我是 EINDEX Li</p>
    <p>在 Thoughtworks 养老中</p>

    <p>这是我不知道第多少个 Blog 系统了</p>
    <p>这次自己开发，不管怎样，希望这个系统能和我一起成长</p>
    <p>把自己想用在 BS 架构上的技术都使用这个平台实践一下</p>
    <p>：）</p>
  </post-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
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
