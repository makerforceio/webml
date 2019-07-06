<template>
  <div class="fader fixed inset-0 p-6 flex items-center justify-center">
    <div class="bg-white rounded-card shadow-md m-6">
      <div class="bg-primary text-white rounded-card flex items-center justify-start">
		  <button class="m-5">
			  <fa-icon size="2x" :icon="['far', 'times']" />
		  </button>
		<h1 class="text-2xl font-bold">New Session</h1>
      </div>
	  <div class="p-2">
		  <div class="m-4">
			  <input 
			   class="text-2xl font-bold"
				v-model="sessionName"
				placeholder="Enter session name" />
		  </div>
		  <div class="m-4">
			  <DropArea @input="sessionModel = $event" placeholder="Drop model here" />
		  </div>
		  <div class="m-4">
			  <DropArea @input="sessionData = $event" placeholder="Drop data here" />
		  </div>
		  <div class="m-4">
			  <DropArea @input="sessionLabels = $event" placeholder="Drop labels here" />
		  </div>
		  <div class="m-4">
			  <DropArea @input="sessionDataParser = $event" placeholder="Drop data parser here" />
		  </div>
	  </div>
    </div>
  </div>
</template>

<style scoped>
.fader {
  background: rgba(0, 0, 0, 0.5);
}
</style>

<script>
import DropArea from '~/components/common/DropArea.vue';
export default {
  components: {
  DropArea
  },
  data() {
    return {
      sessionName: ''
    }
  },
  methods: {
    onSubmit: async function(files, submitType) {
      // let files = this.$refs.fileloader.files

      var url = (function(type) {
        switch (type) {
          case 'dataset':
            return 'localhost:10200/data'
          case 'labels':
            return 'localhost:10200/labels'
          case 'scripts':
            return 'localhost:10200/data_parse'
          default:
            return ''
        }
      })(submitType)

      for (var i = 0; i < files.length; i++) {
        await fetch(url, {
          method: 'PUT',
          redirect: 'follow',
          body: files[i]
        })

        console.log(`Uploaded ${files[i].name}`)
      }
    }
  }
}
</script>

<style></style>
