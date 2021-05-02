<template>
  <div class="divide-y divide-gray-200">
    <ul class="divide-y divide-gray-200" :key="post.slug" v-for="post in postItems">
      <li class="py-12">
        <article class="space-y-3">
          <div>
            <h2
              class="text-2xl font-bold tracking-tight text-gray-900"
              @click="jumpToPost(post.slug)"
            >
              <a :href="postURL(post.slug)">{{ post.title }}</a>
            </h2>
          </div>
          <div>
            <content class="prose max-w-none text-gray-600" v-html="post.html"></content>
          </div>
          <div class="text-base leading-4 font-medium">
            <a
              class="text-teal-500 hover:text-teal-600 hover:bg-gray-700 hover:text-white -mx-2 px-2 py-2 rounded-md"
              :href="postURL(post.slug)"
              >Read more →</a
            >
          </div>
        </article>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

type PostItem = {
  slug: string;
  title: string;
  content: string;
  html: string;
};

export default defineComponent({
  name: 'PostList',
  data() {
    return {
      postItems: [] as PostItem[],
    };
  },
  setup() {},
  async mounted() {
    try {
      const postsResp = await this.axios.get('api/posts');
      if (postsResp.status == 200) {
        this.postItems = postsResp.data;
      }
    } catch (error) {}
  },
  methods: {
    jumpToPost(slug: string) {
      this.$router.push(this.postURL(slug));
    },
    postURL(slug: string): string {
      return `/posts/${slug}`;
    },
  },
});
</script>
