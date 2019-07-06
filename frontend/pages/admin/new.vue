<template>
  <div>
    <input v-model="sessionName", placeholder="Enter session name">
    <input ref="fileloader" type="file" webkitdirectory mozdirectory/>
  </div>
</template>

<script>
export default {
  layout: 'default',
  methods: {
    onSubmit: async function (submitType) {
      let files = this.$refs.fileloader.files

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
