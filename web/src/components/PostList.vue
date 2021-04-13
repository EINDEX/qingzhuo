<template>
    <article :key="post.slug" v-for="post in postItems">
        <h2 @click="jumpToPost(post.slug)">{{post.title}}</h2>
        <content v-html="post.content"></content>
    </article>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

type PostItem = {
    slug: string,
    title: String,
    content: String
}

export default defineComponent({
    name: 'PostList',
    data() {
        return {
            postItems: [] as PostItem[]
        }
    },
    setup() {
        
    },
    async mounted() {
        try {
            const postsResp = await this.axios.get('api/posts')
            if (postsResp.status == 200){
                this.postItems = postsResp.data;
            }
        } catch (error) {
            
        }
    },
    methods: {
        jumpToPost(slug: String) {
            this.$router.push(`/posts/${slug}`)
        }
    }
})
</script>
