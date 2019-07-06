<template>
  <div>
    <Header title="New session" subtitle="" class="h-screen rounded-none">
      <template v-slot:buttons-left>
        <button class="my-3 mr-4">
          <fa-icon size="2x" :icon="['far', 'arrow-left']" />
        </button>
      </template>
      <template v-slot:content>
        <div class="mx-4 my-8" title="Major graph">
          <h2 class="font-medium">
            Loss
          </h2>
        </div>
      </template>
    </Header>
    <input v-model="sessionName", placeholder="Enter session name">
    <input ref="fileloader" type="file" webkitdirectory mozdirectory/>
  </div>
</template>

<script>
import Header from '~/components/common/Header.vue'
import trend from 'vuetrend'

export default {
  layout: 'client',
  components: {
    Header,
    trend
  },
  methods: {
    onSubmit: async function (files, submitType) {
      // let files = this.$refs.fileloader.files

      var url = (function(type) {
        switch(type) {
          case 'dataset':
            return "localhost:10200/data";
          case 'labels':
            return "localhost:10200/labels";
          case 'scripts':
            return "localhost:10200/data_parse";
          default:
            return "";
        }
      })(submitType);

      for (var i = 0; i < files.length; i++) {
        await fetch(url, {
          method: "PUT",
          redirect: "follow",
          body: files[i]
        })

        console.log(`Uploaded ${files[i].name}`)
      }
    },
  },
}
</script>

<style>
</style>
