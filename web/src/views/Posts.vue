<template>
<article>
    <h1>{{blog.title}}</h1>
    <section>
        {{blog.content}}
    </section>
</article>
    
</template>

<script lang="ts">
import { defineComponent } from 'vue'

type Post = {
    slug: String,
    title: String,
    content: String,
}

export default defineComponent({
    setup() {

    },
    props: {
        slug: String
    },
    data() {
        return { 
            blog: {} as Post
        }
    },
    async mounted() {
        try {
            const postResp = await this.axios.get(`/api/posts/${this.slug}`)
            if (postResp.status == 200) {
                this.blog = postResp.data;
            }
        } catch (error) {
            
        }
    }
})
</script>
