<template>
  <div>
    <input v-model="sessionName", placeholder="Enter session name">
    <input ref="fileloader" type="file" webkitdirectory mozdirectory/>
</template>

<script>
export default {
  layout: 'default',
  methods: {
    onSubmit: async function (submitType) {
      let files = this.$refs.fileloader.files

      switch(submitType) {
        case "dataset":
          let url = "localhost:10200/data"
          break
        case "labels":
          let url = "localhost:10200/labels"
          break
        case "scripts":
          let url = "localhost:10200/data_parser"
          break
        default:
          let url = ""
      }

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
