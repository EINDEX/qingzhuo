<template>
  <div class="py-12 space-y-3">
    <div v-for="year in archiveYears()" :key="year">
      <div>
        <span>{{ year }}</span>
      </div>
      <div class="divide-y pt-2 px-4">
        <ul v-for="archiveItem in archivesByYears[year]" :key="archiveItem.slug">
          <li class="py-3">
            <article class="space-y-3">
              <div class="text-xl text-gray-700 hover:text-black">
                <a :href="postURL(archiveItem.slug)">{{ archiveItem.title }}</a>
              </div>
            </article>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import {ArchiveList} from '../types/API';

export default defineComponent({
    name: 'ArchiveList',
    setup() {},
    data() {
      return {
        archives: [] as ArchiveList,
      };
    },
    async mounted() {
      try {
        const postsResp = await this.axios.get('api/archives');
        if (postsResp.status == 200) {
          this.archives = postsResp.data.data;
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
      archiveYears(): string[] {
        for (let archive in this.archives) {
        }
        Object.keys(this.archivesByYears).sort().reverse();
      },
    },
  });
</script>
