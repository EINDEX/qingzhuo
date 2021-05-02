<template>
  <article class="py-12 divide-y">
    <header class="py-6">
      <h1 class="text-3xl leading-8 font-bold tracking-tight text-gray-900">{{ blog.title }}</h1>
    </header>
    <main class="py-6 h-auto prose lg:prose-xl">
      {{ blog.content }}
    </main>
  </article>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

type Post = {
  slug: String;
  title: String;
  content: String;
};

export default defineComponent({
  setup() {},
  props: {
    slug: String,
  },
  data() {
    return {
      blog: {} as Post,
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
