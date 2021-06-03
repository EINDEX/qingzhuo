<template>
  <article class="py-12 divide-y h-full">
    <header class="py-6">
      <h1 class="text-3xl leading-8 font-bold tracking-tight text-gray-900">
        {{ title }}
      </h1>
    </header>
    <content class="py-6 h-auto prose">
      <slot />
    </content>
  </article>
</template>

<script lang="ts">
import {defineComponent} from 'vue';

export default defineComponent({
    props: {
      title: String,
    },
    setup() {},
  });
</script>
