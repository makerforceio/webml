<template>
  <div
    class="flex-grow flex flex-col bg-white rounded-card shadow-md m-4 py-4 px-6"
  >
    <h2 v-if="title" class="font-medium">{{ title }}</h2>
    <slot></slot>
  </div>
</template>

<script>
export default {
  props: {
    title: String
  }
}
</script>
