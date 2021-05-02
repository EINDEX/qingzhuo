<template>
  <post-page :title="blog.title"><div v-html="blog.html"></div></post-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import PostPage from '/@/components/PostPage.vue';
import PostType from '/@/types/API';

export default defineComponent({
  components: { PostPage },
  setup() {},
  props: {
    slug: String,
  },
  data() {
    return {
      blog: {} as PostType,
    };
  },
  async mounted() {
    try {
      const postResp = await this.axios.get(`/api/posts/${this.slug}`);
      if (postResp.status == 200) {
        this.blog = postResp.data;
      }
    } catch (error) {}
  },
});
</script>
